package cmds

// 注文履歴の検索

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"
	"net/url"

	"github.com/google/subcommands"
)

type GetOrderCmd struct {
	symbol    string
	types     string
	startdate string
	enddate   string
	states    string
	from      string
	direct    string
	size      string

	isSave bool
}

func (a *GetOrderCmd) Name() string {
	return "getOrder"
}

func (a *GetOrderCmd) Synopsis() string {
	return "GetOrderCmd"
}

func (a *GetOrderCmd) Usage() string {
	return "api-test GetOrderCmd -save"
}

func (a *GetOrderCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "btcjpy", "symbol success")
	set.StringVar(&a.states, "states", "filled", "states success")

	set.StringVar(&a.types, "types", "", "types success")
	set.StringVar(&a.startdate, "start_date", "", "startdate success")
	set.StringVar(&a.enddate, "end_date", "", "enddate success")
	set.StringVar(&a.from, "from", "", "from success")
	set.StringVar(&a.direct, "direct", "next", "direct success")
	set.StringVar(&a.size, "size", "10", "size success")
	set.BoolVar(&a.isSave, "save", false, "write to json")
	return
}

func (a *GetOrderCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	param := url.Values{}
	param.Add("symbol", a.symbol)
	param.Add("states", a.states)

	param.Add("types", a.types)
	param.Add("start_date", a.startdate)
	param.Add("end_date", a.enddate)
	param.Add("from", a.from)
	param.Add("direct", a.direct)
	param.Add("size", a.size)

	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/order/orders")+"?"+param.Encode(), nil)
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
