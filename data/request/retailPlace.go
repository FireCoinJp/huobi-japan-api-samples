package request

type RetailPlaceBody struct {
	Id               string `json:"id"`
	Symbol           string `json:"symbol"`
	Types            string `json:"type"`
	Amount           string `json:"amount,omitempty"`
	Price            string `json:"price"`
	Source           string `json:"source"`
	ClientOrderId    string `json:"client_order_id,omitempty"`
	CashAmount       string `json:"cash_amount,omitempty"`
	OrderInstruction string `json:"order-instruction"`
}
