package request

type CreateBody struct {
	Address  string `json:"address"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Fee      string `json:"fee"`
	AddrTag  string `json:"addr-tag,omitempty"`
}
