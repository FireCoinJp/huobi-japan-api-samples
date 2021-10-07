package cmds

// 販売所注文履歴

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"
	"net/url"

	"github.com/google/subcommands"
)

type OrderlistCmd struct {
	id             string
	limit          string
	from           string
	direct         string
	base_currency  string
	quote_currency string
	symbol         string
	ordertype      string
	states         string

	isSave bool
}

func (a *OrderlistCmd) Name() string {
	return "orderlist"
}

func (a *OrderlistCmd) Synopsis() string {
	return "OrderlistCmd"
}

func (a *OrderlistCmd) Usage() string {
	return "api-test OrderlistCmd -save"
}

func (a *OrderlistCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.states, "states", "1", "states success")
	set.StringVar(&a.direct, "direct", "1", "direct success")

	set.StringVar(&a.id, "id", "", "id success")
	set.StringVar(&a.limit, "limit", "", "limit success")
	set.StringVar(&a.from, "from", "", "from success")
	set.StringVar(&a.base_currency, "base_currency", "", "base_currency success")
	set.StringVar(&a.quote_currency, "quote_currency", "", "quote_currency success")
	set.StringVar(&a.symbol, "symbol", "btcjpy", "symbol success")
	set.StringVar(&a.ordertype, "ordertype", "", "types success")
	set.BoolVar(&a.isSave, "save", false, "write to json")
	return
}

func (a *OrderlistCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	param := url.Values{}
	param.Add("states", a.states)
	param.Add("direct", a.direct)

	param.Add("id", a.id)
	param.Add("limit", a.limit)
	param.Add("from", a.from)
	param.Add("start_date", a.base_currency)
	param.Add("end_date", a.quote_currency)
	param.Add("symbol", a.symbol)
	param.Add("types", a.ordertype)

	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/retail/order/list")+"?"+param.Encode(), nil)
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
