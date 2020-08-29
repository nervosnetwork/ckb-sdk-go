package utils

import (
	"context"

	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type CollectResult struct {
	Cells    []*types.Cell
	Capacity uint64
	Options  map[string]interface{}
}

type CellProcessor interface {
	Process(*types.Cell, *CollectResult) (bool, error)
}

type CapacityCellProcessor struct {
	Max uint64
}

func NewCapacityCellProcessor(capacity uint64) *CapacityCellProcessor {
	return &CapacityCellProcessor{
		Max: capacity,
	}
}

func (p *CapacityCellProcessor) Process(cell *types.Cell, result *CollectResult) (bool, error) {
	result.Capacity = result.Capacity + cell.Capacity
	result.Cells = append(result.Cells, cell)
	if p.Max > 0 && result.Capacity >= p.Max {
		return true, nil
	}
	return false, nil
}

type CellCollector struct {
	Client          rpc.Client
	LockScript      *types.Script
	TypeScript      *types.Script
	Processor       CellProcessor
	UseIndex        bool
	EmptyData       bool
	FromBlockNumber uint64
}

func NewCellCollector(client rpc.Client, lockScript *types.Script, processor CellProcessor, fromBlockNumber uint64) *CellCollector {
	return &CellCollector{
		Client:          client,
		LockScript:      lockScript,
		Processor:       processor,
		EmptyData:       true,
		FromBlockNumber: fromBlockNumber,
	}
}

func (c *CellCollector) Collect() (*CollectResult, error) {
	lockHash, err := c.LockScript.Hash()
	if err != nil {
		return nil, err
	}
	return c.collectFromBlocks(lockHash)
}

func (c *CellCollector) collectFromBlocks(lockHash types.Hash) (*CollectResult, error) {
	header, err := c.Client.GetTipHeader(context.Background())
	if err != nil {
		return nil, err
	}
	var result CollectResult
	result.Options = make(map[string]interface{})
	start := c.FromBlockNumber
	var stop bool
	for {
		end := start + 100
		if end > header.Number {
			end = header.Number
			stop = true
		}
		cells, err := c.Client.GetCellsByLockHash(context.Background(), lockHash, start, end)
		if err != nil {
			return nil, err
		}
		for _, cell := range cells {
			if c.TypeScript != nil {
				if !c.TypeScript.Equals(cell.Type) {
					continue
				}
			} else {
				if cell.Type != nil {
					continue
				}
			}
			if c.EmptyData && cell.OutputDataLen > 0 {
				continue
			}
			s, err := c.Processor.Process(cell, &result)
			if err != nil {
				return nil, err
			}
			if s {
				stop = s
				break
			}
		}
		if stop {
			break
		}
		start = end + 1
	}
	return &result, nil
}
