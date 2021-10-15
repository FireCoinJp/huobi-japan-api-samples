package cmds

// 注文実行

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

type PlaceCmd struct {
	amount   string
	price    string
	source   string
	symbol   string
	steptype string
}

func (a *PlaceCmd) Name() string {
	return "place"
}

func (a *PlaceCmd) Synopsis() string {
	return "注文実行"
}

func (a *PlaceCmd) Usage() string {
	return "api-test place \n"
}

func (a *PlaceCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.amount, "amount", "1", "取引数量")
	set.StringVar(&a.price, "price", "9.5", "指値の注文価格")
	set.StringVar(&a.source, "source", "api", "注文のソース, default: api")
	set.StringVar(&a.symbol, "symbol", "trxjpy", "取引通貨ペア")
	set.StringVar(&a.steptype, "type", "buy_limit", "注文タイプ,buy-market：成行買い,sell-market：成行売り,buy-limit：指値買い,sell-limit：指値売り,buy-ioc：IOC買い注文,sell-ioc：IOC売り注文,buy-limit-maker,sell-limit-maker")
}

func (a *PlaceCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	sendBody := request.PlaceBody{
		AccountId: config.Cfg.AccountID,
		Amount:    a.amount,
		Price:     a.price,
		Source:    "api",
		Symbol:    a.symbol,
		Steptype:  a.steptype,
	}

	placeBody, _ := json.Marshal(sendBody)

	req, _ := http.NewRequest(http.MethodPost, h.Url("/v1/order/orders/place"), bytes.NewReader(placeBody))
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	h.Process(req)
	return 0
}
