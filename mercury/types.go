package mercury

//type rpcBalanceResp struct {
//	KeyAddress    string `json:"key_address"`
//	AssetInfo       string `json:"udt_hash"`
//	Unconstrained string `json:"unconstrained"`
//	Fleeting      string `json:"fleeting"`
//	Locked        string `json:"locked"`
//}

//type rpcGetBalanceResponse struct {
//	Balances []*rpcBalanceResp `json:"balances"`
//}

// -----------------------------------------------------------------------------------------------------------------------------------------------

//type rpcTransactionInfoWithStatusResponse struct {
//	Transaction     rpcTransactionInfoResponse `json:"transaction"`
//	Status          types.TransactionStatus    `json:"status"`
//	BlockHash       string                     `json:"block_hash"`
//	BlockNumber     uint64                     `json:"block_number"`
//	ConfirmedNumber uint64                     `json:"confirmed_number"`
//}

//type rpcTransactionInfoResponse struct {
//	TxHash     string               `json:"tx_hash"`
//	Operations []*rpcRecordResponse `json:"operations"`
//}

//type rpcRecordResponse struct {
//	Id            uint          `json:"id"`
//	KeyAddress    string        `json:"key_address"`
//	NormalAddress string        `json:"normal_address"`
//	Amount        rpcAmountResp `json:"amount"`
//}

//type rpcAmountResp struct {
//	Value   string      `json:"value"`
//	AssetInfo string      `json:"udt_hash"`
//	Status  interface{} `json:"status"`
//}

//func toTransactionInfoWithStatusResponse(tx *rpcTransactionInfoWithStatusResponse) (*resp.GetTransactionInfoResponse, error) {
//	transactionInfoResponse, err := toTransactionInfoResponse(tx.Transaction.Operations, tx.Transaction.TxHash)
//	if err != nil {
//		return nil, err
//	}
//	result := &resp.GetTransactionInfoResponse{
//		Status:          tx.Status,
//		BlockHash:       tx.BlockHash,
//		BlockNumber:     tx.BlockNumber,
//		ConfirmedNumber: tx.ConfirmedNumber,
//		Transaction:     transactionInfoResponse,
//	}
//
//	return result, nil
//}
//
//func toTransactionInfoResponse(txs []*rpcRecordResponse, txHash string) (*resp.TransactionInfo, error) {
//	infoResponse := &resp.TransactionInfo{TxHash: txHash}
//	for _, op := range txs {
//
//		var asset *common.AssetInfo
//		if op.Amount.Status == common.CKB {
//			asset = common.NewCkbAsset()
//		} else {
//			asset = common.NewUdtAsset(op.Amount.AssetInfo)
//		}
//
//		var status map[resp.AssetStatus]uint
//		data, err := json.Marshal(op.Amount.Status)
//		if err != nil {
//			return nil, err
//		}
//		json.Unmarshal(data, status)
//
//		var assetStatus resp.AssetStatus
//		var blockNumber uint
//		if _, ok := status[resp.FIXED]; ok {
//			assetStatus = resp.FIXED
//			blockNumber = status[resp.FIXED]
//		} else {
//			assetStatus = resp.CLAIMABLE
//			blockNumber = status[resp.FIXED]
//		}
//
//		infoResponse.Operations = append(infoResponse.Operations, &resp.Record{
//			op.Id,
//			op.KeyAddress,
//			op.Amount.Value,
//			asset,
//			assetStatus,
//			blockNumber,
//		})
//
//	}
//
//	return infoResponse, nil
//}

// -----------------------------------------------------------------------------------------------------------------------------------------------
//type rpcBlockInfoResponse struct {
//	BlockNumber     uint64                        `json:"block_number"`
//	BlockHash       string                        `json:"block_hash"`
//	ParentBlockHash string                        `json:"parent_block_hash"`
//	Timestamp       uint64                        `json:"timestamp"`
//	Transactions    []*rpcTransactionInfoResponse `json:"transactions"`
//}
