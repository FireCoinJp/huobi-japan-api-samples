package cmds

// 販売所での注文

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"huobi-japan-api-samples/data/request"
	"net/http"

	"github.com/google/subcommands"
)

type RetailPlaceCmd struct {
	id            string
	symbol        string
	types         string
	amount        string
	price         string
	clientOrderId string
	cashAmount    string
}

func (a *RetailPlaceCmd) Name() string {
	return "retailPlace"
}

func (a *RetailPlaceCmd) Synopsis() string {
	return "販売所での注文"
}

func (a *RetailPlaceCmd) Usage() string {
	return "api-test retailPlace \n"
}

func (a *RetailPlaceCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.id, "id", "791025bfdc3b45459b96b5d776ede78f", "websocketから取得されたリアルタイム価格のID, 32桁")
	set.StringVar(&a.symbol, "symbol", "btcjpy", "取引ペア")
	set.StringVar(&a.types, "type", "1", "注文方向, 1:購入, 2:売却")
	set.StringVar(&a.amount, "amount", "10.12", "取引量, decimal(36,18)")
	set.StringVar(&a.price, "price", "9.5", "取引価格, decimal(36,18)")
	set.StringVar(&a.clientOrderId, "client_order_id", "", "クライアントカスタマイズID")
	set.StringVar(&a.cashAmount, "cash_amount", "", "現金金額, decimal(36,18)")
}

func (a *RetailPlaceCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	sendBody := request.RetailPlaceBody{
		Id:               a.id,
		Symbol:           a.symbol,
		Types:            a.types,
		Amount:           a.amount,
		Price:            a.price,
		Source:           "4",
		ClientOrderId:    a.clientOrderId,
		CashAmount:       a.cashAmount,
		OrderInstruction: "1",
	}

	retailPlaceBody, _ := json.Marshal(sendBody)

	req, _ := http.NewRequest(http.MethodPost, h.Url("/v1/retail/order/place"), bytes.NewReader(retailPlaceBody))
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	h.Process(req)
	return 0
}
