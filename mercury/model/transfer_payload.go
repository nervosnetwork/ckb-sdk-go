package model

type TransferPayload struct {
	UdtHash string          `json:"udt_hash,omitempty"`
	From    *FromAccount    `json:"from"`
	Items   []*TransferItem `json:"items"`
	Change  string          `json:"change,omitempty"`
	FeeRate uint            `json:"fee_rate"`
}

type FromAccount struct {
	Idents []string `json:"idents"`
	Source string   `json:"source"`
}

type TransferItem struct {
	To     *ToAccount `json:"to"`
	Amount uint       `json:"amount"`
}

type ToAccount struct {
	Ident  string `json:"ident"`
	Action string `json:"action"`
}

type transferBuilder struct {
	UdtHash string          `json:"udt_hash,omitempty"`
	From    *FromAccount    `json:"from"`
	Items   []*TransferItem `json:"items"`
	Change  string          `json:"change,omitempty"`
	FeeRate uint            `json:"fee_rate"`
}

func NewTransferBuilder() *transferBuilder {
	// default fee rate
	return &transferBuilder{
		FeeRate: 1000,
	}
}

func (builder *transferBuilder) AddUdtHash(udtHash string) {
	builder.UdtHash = udtHash
}

func (builder *transferBuilder) AddFrom(idents []string, source string) {
	form := &FromAccount{
		Idents: idents,
		Source: source,
	}
	builder.From = form
}

func (builder *transferBuilder) AddItem(ident, action string, amount uint) {
	to := &ToAccount{
		Ident:  ident,
		Action: action}
	item := &TransferItem{Amount: amount}
	item.To = to
	builder.Items = append(builder.Items, item)
}

func (builder *transferBuilder) AddChange(change string) {
	builder.Change = change
}

func (builder *transferBuilder) AddFeeRate(feeRate uint) {
	builder.FeeRate = feeRate
}

func (builder *transferBuilder) Build() *TransferPayload {
	return &TransferPayload{
		builder.UdtHash,
		builder.From,
		builder.Items,
		builder.Change,
		builder.FeeRate,
	}

}
