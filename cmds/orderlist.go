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
}

func (a *OrderlistCmd) Name() string {
	return "orderlist"
}

func (a *OrderlistCmd) Synopsis() string {
	return "販売所注文履歴"
}

func (a *OrderlistCmd) Usage() string {
	return "api-test orderlist \n"
}

func (a *OrderlistCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.states, "states", "1", "成約状態, 1: 進行中, 2: 完全約定, 3: 未成約")
	set.StringVar(&a.direct, "direct", "1", "注文方向, 1:next, 2:previous")

	set.StringVar(&a.id, "id", "", "注文番号")
	set.StringVar(&a.limit, "limit", "", "表示件数, default=10, max: 100")
	set.StringVar(&a.from, "from", "", "開始ID, １ページ以後必要")
	set.StringVar(&a.base_currency, "base_currency", "", "基礎通貨")
	set.StringVar(&a.quote_currency, "quote_currency", "", "通貨単位")
	set.StringVar(&a.symbol, "symbol", "btcjpy", "取引ペア")
	set.StringVar(&a.ordertype, "ordertype", "", "取引タイプ, 1:buy, 2:sell")
}

func (a *OrderlistCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	param := url.Values{}
	param.Add("states", a.states)
	param.Add("direct", a.direct)
	param.Add("symbol", a.symbol)
	if a.id != "" {
		param.Add("id", a.id)
	}
	if a.limit != "" {
		param.Add("limit", a.limit)
	}
	if a.from != "" {
		param.Add("from", a.from)
	}
	if a.base_currency != "" {
		param.Add("start_date", a.base_currency)
	}
	if a.quote_currency != "" {
		param.Add("end_date", a.quote_currency)
	}
	if a.ordertype != "" {
		param.Add("types", a.ordertype)
	}

	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/retail/order/list")+"?"+param.Encode(), nil)
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	h.Process(req)
	return 0
}
