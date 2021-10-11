package request

type BatchCancelOpenOrdersBody struct {
	AccountId string `json:"account-id"`
	Side      string `json:"side,omitempty"`
	Size      string `json:"size,omitempty"`
	Types     string `json:"types,omitempty"`
	Symbol    string `json:"symbol,omitempty"`
}
