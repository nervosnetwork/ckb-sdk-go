package model

type TransferPayload struct {
	UdtHash string          `json:"udt_hash,omitempty"`
	From    *FromAccount    `json:"from"`
	Items   []*TransferItem `json:"items"`
	Change  string          `json:"change,omitempty"`
	Fee     uint            `json:"fee"`
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

type TransferBuilder struct {
	UdtHash string          `json:"udt_hash,omitempty"`
	From    *FromAccount    `json:"from"`
	Items   []*TransferItem `json:"items"`
	Change  string          `json:"change,omitempty"`
	Fee     uint            `json:"fee"`
}

func (builder *TransferBuilder) AddUdtHash(udtHash string) {
	builder.UdtHash = udtHash
}

func (builder *TransferBuilder) AddFrom(idents []string, source string) {
	form := &FromAccount{
		Idents: idents,
		Source: source,
	}
	builder.From = form
}

func (builder *TransferBuilder) AddItem(ident, action string, amount uint) {
	to := &ToAccount{
		Ident:  ident,
		Action: action}
	item := &TransferItem{Amount: amount}
	item.To = to
	builder.Items = append(builder.Items, item)
}

func (builder *TransferBuilder) AddChange(change string) {
	builder.Change = change
}

func (builder *TransferBuilder) AddFee(fee uint) {
	builder.Fee = fee
}

func (builder *TransferBuilder) Build() *TransferPayload {
	return &TransferPayload{
		builder.UdtHash,
		builder.From,
		builder.Items,
		builder.Change,
		builder.Fee,
	}

}
