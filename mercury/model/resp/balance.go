package resp

type Balance struct {
	Unconstrained string `json:"unconstrained"`
	Fleeting      string `json:"fleeting"`
	Locked        string `json:"locked"`
}
