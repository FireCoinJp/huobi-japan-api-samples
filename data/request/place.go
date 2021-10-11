package request

type PlaceBody struct {
	AccountId string `json:"account-id"`
	Amount    string `json:"amount"`
	Price     string `json:"price,omitempty"`
	Source    string `json:"source,omitempty"`
	Symbol    string `json:"symbol"`
	Steptype  string `json:"type"`
}
