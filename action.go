package main

type action string

const (
	accounts             action = "accounts"
	balance              action = "balance"
	place                action = "place"
	openOrder            action = "openOrders"
	cancel               action = "cancel"
	batchcancel          action = "batchcancel"
	batchcancelopenorder action = "batchcancelopenorder"
	order                action = "order"
	matchresult          action = "matchresult"
	matchresults         action = "matchresults"
	orders               action = "orders"
	orderHistory         action = "order_history"
	fee	action = "fee"
)

func (a action) Path() string {
	switch a {
	case accounts:
		return "/v1/account/accounts"
	case balance:
		return "/v1/account/accounts/%s/balance"
	case place:
		return "/v1/order/orders/place"
	case openOrder:
		return "/v1/order/openOrders"
	case cancel:
		return "/v1/order/orders/%s/submitcancel"
	case batchcancel:
		return "/v1/order/orders/batchcancel"
	case batchcancelopenorder:
		return "/v1/order/orders/batchCancelOpenOrders"
	case order:
		return "/v1/order/orders/%s"
	case matchresult:
		return "/v1/order/orders/%s/matchresults"
	case matchresults:
		return "/v1/order/matchresults?symbol=xrpjpy"
	case orders:
		return "/v1/order/orders?symbol=xrpjpy&states=submitted"
	case orderHistory:
		return "/v1/order/history"
	case fee:
		return "/v2/reference/transact-fee-rate"
	}
	return ""
}

//func (a action) Body() io.Reader {
//	s := ""
//	switch a {
//	case place:
//		s = fmt.Sprintf(`{
//   "account-id": "%s",
//   "amount": "0.1",
//   "price": "105",
//   "source": "api",
//   "symbol": "xrpjpy",
//   "type": "buy-limit"
//}`, accountId)
//	case cancel:
//	case batchcancel:
//	case batchcancelopenorder:
//	case order:
//	case matchresult:
//	case matchresults:
//	case orders:
//	}
//
//	if len(s) == 0 {
//		return nil
//	}
//
//	return strings.NewReader(s)
//}

//func (a action) createRequest(act action) *http.Request {
//	var req *http.Request
//	switch act {
//	case accounts:
//		req, _ = http.NewRequest(http.MethodGet, h.url(action.Path()), nil)
//	case balance:
//		req, _ = http.NewRequest(http.MethodGet, h.url(fmt.Sprintf(action.Path(), accountId)), nil)
//	case place:
//		req, _ = http.NewRequest(http.MethodPost, h.url(action.Path()), action.Body())
//	case openOrder:
//		req, _ = http.NewRequest(http.MethodGet, h.url(action.Path()), nil)
//	case cancel:
//		req, _ = http.NewRequest(http.MethodPost, h.url(fmt.Sprintf(action.Path(), orderId)), nil)
//	case batchcancel:
//		req, _ = http.NewRequest(http.MethodPost, h.url(action.Path()), nil)
//	case batchcancelopenorder:
//		req, _ = http.NewRequest(http.MethodPost, h.url(action.Path()), nil)
//	case order:
//		req, _ = http.NewRequest(http.MethodGet, h.url(fmt.Sprintf(action.Path(), orderId)), nil)
//	case matchresult:
//		req, _ = http.NewRequest(http.MethodGet, h.url(fmt.Sprintf(action.Path(),orderId)), nil)
//		s, _ := httputil.DumpRequest(req, true)
//		fmt.Printf("%s", s)
//	case matchresults:
//		req, _ = http.NewRequest(http.MethodGet, h.url(action.Path()), nil)
//	case orders:
//		req, _ = http.NewRequest(http.MethodGet, h.url(action.Path()), nil)
//	case orderHistory:
//		req, _ = http.NewRequest(http.MethodGet, h.url(action.Path()), nil)
//	case fee:
//		req, _ = http.NewRequest(http.MethodGet, h.url(action.Path()), nil)
//	}
//
//}

//pAction := flag.String("action", "accounts", "")
//flag.Parse()
//action := action(*pAction)

//h := NewH()
//var req *http.Request
//
//err := h.save(action, h.do(req))
//if err != nil {
//	panic(err)
//}
