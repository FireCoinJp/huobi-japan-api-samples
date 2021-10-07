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
	accountId string
	amount    string
	price     string
	source    string
	symbol    string
	steptype  string
	isSave    bool
}

func (a *PlaceCmd) Name() string {
	return "place"
}

func (a *PlaceCmd) Synopsis() string {
	return "PlaceCmd"
}

func (a *PlaceCmd) Usage() string {
	return "api-test PlaceCmd -save"
}

func (a *PlaceCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.accountId, "account_id", config.Cfg.AccountID, "account-id success")
	set.StringVar(&a.amount, "amount", "1", "amount success")
	set.StringVar(&a.price, "price", "9.5", "price success")
	set.StringVar(&a.source, "source", "api", "source success")
	set.StringVar(&a.symbol, "symbol", "trxjpy", "symbol success")
	set.StringVar(&a.steptype, "type", "buy-limit", "type success")
	set.BoolVar(&a.isSave, "save", false, "write to json")
	return
}

func (a *PlaceCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	sendBody := request.PlaceBody{
		AccountId: a.accountId,
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

	if a.isSave {
		err = h.Do(req, api.SaveMsg)
	} else {
		err = h.Do(req, api.PrintMsg)
	}

	if err != nil {
		panic(err)
	}
	return 0
}
