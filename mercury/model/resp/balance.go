package resp

type Balance struct {
	Unconstrained string `json:"owned"`
	Fleeting      string `json:"claimable"`
	locked        string `json:"locked"`
}
