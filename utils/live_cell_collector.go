package utils

import (
	"context"
	"github.com/nervosnetwork/ckb-sdk-go/indexer"
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/pkg/errors"
)

type LiveCellCollectResult struct {
	LiveCells []*indexer.LiveCell
	Capacity  uint64
	Options   map[string]interface{}
}

type LiveCellProcessor interface {
	Process(*indexer.LiveCell, *LiveCellCollectResult) (bool, error)
}

type CapacityLiveCellProcessor struct {
	Max        uint64
	EmptyData  bool
	TypeScript *types.Script
}

func NewCapacityLiveCellProcessor(capacity uint64) *CapacityLiveCellProcessor {
	return &CapacityLiveCellProcessor{
		Max: capacity,
	}
}

func (p *CapacityLiveCellProcessor) Process(liveCell *indexer.LiveCell, result *LiveCellCollectResult) (bool, error) {
	if p.TypeScript != nil {
		if !p.TypeScript.Equals(liveCell.Output.Type) {
			return false, nil
		}
	} else {
		if liveCell.Output.Type != nil {
			return false, nil
		}
	}
	if p.EmptyData && len(liveCell.OutputData) > 0 {
		return false, nil
	}
	result.Capacity = result.Capacity + liveCell.Output.Capacity
	result.LiveCells = append(result.LiveCells, liveCell)
	if p.Max > 0 && result.Capacity >= p.Max {
		return true, nil
	}
	return false, nil
}

type LiveCellCollector struct {
	Client      rpc.Client
	SearchKey   *indexer.SearchKey
	SearchOrder indexer.SearchOrder
	Limit       uint64
	AfterCursor string
	Processor   LiveCellProcessor
}

func (c *LiveCellCollector) collectFromCkbIndexer() (*LiveCellCollectResult, error) {
	cursor := c.AfterCursor
	var result LiveCellCollectResult
	var stop bool
	for {
		liveCells, err := c.Client.GetCells(context.Background(), c.SearchKey, c.SearchOrder, c.Limit, cursor)
		if err != nil {
			return nil, err
		}
		for _, cell := range liveCells.Objects {
			s, err := c.Processor.Process(cell, &result)
			if err != nil {
				return nil, err
			}
			if s {
				stop = s
				break
			}
		}
		if stop || len(liveCells.Objects) < int(c.Limit) || liveCells.LastCursor == "" {
			break
		}
		cursor = liveCells.LastCursor
	}
	return &result, nil
}

func NewLiveCellCollector(client rpc.Client, searchKey *indexer.SearchKey, searchOrder indexer.SearchOrder, limit uint64, afterCursor string, processor LiveCellProcessor) *LiveCellCollector {
	return &LiveCellCollector{
		Client:      client,
		SearchKey:   searchKey,
		SearchOrder: searchOrder,
		Limit:       limit,
		AfterCursor: afterCursor,
		Processor:   processor,
	}
}

func (c *LiveCellCollector) Collect() (*LiveCellCollectResult, error) {
	if c.SearchKey == nil {
		return nil, errors.New("missing SearchKey error")
	}
	if c.SearchOrder != indexer.SearchOrderAsc && c.SearchOrder != indexer.SearchOrderDesc {
		return nil, errors.New("missing SearchOrder error")
	}
	return c.collectFromCkbIndexer()
}
