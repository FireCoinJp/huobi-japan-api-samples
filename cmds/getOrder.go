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
	starttime string
	endtime   string
	states    string
	from      string
	direct    string
	size      string
}

func (a *GetOrderCmd) Name() string {
	return "getOrder"
}

func (a *GetOrderCmd) Synopsis() string {
	return "注文履歴の検索"
}

func (a *GetOrderCmd) Usage() string {
	return "api-test getOrder \n"
}

func (a *GetOrderCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "trxjp", "取引ペア, [btcjpy, bchbtc,...]")
	set.StringVar(&a.states, "states", "filled", "オーダーのタイプの組み合わせ照会，区切り記号は','を使用。[submitted 提出済み, partial-filled 部分約定, partial-canceled 部分約定キャンセル, filled 完全約定, canceled キャンセル済み]")

	set.StringVar(&a.types, "types", "", "オーダータイプの組み合わせ照会，カンマ区切り, [buy-market：成り行き買い, sell-market：成り行き売り, buy-limit：指値買い, sell-limit：指値売り, buy-ioc：IOC買い注文, sell-ioc：IOC売り注文]")
	set.StringVar(&a.starttime, "start_time", "", "クエリの開始時間。時間形式はミリ秒単位のUTC時間です。 注文トランザクション時間によるクエリ,値の範囲は[（（end-time）– 48h）、（end-time）]、最大クエリウィンドウは48時間、ウィンドウシフト範囲は過去120日です。")
	set.StringVar(&a.endtime, "end_time", "", "クエリの終了時間。時間形式はミリ秒単位のUTC時間です。 注文トランザクション時間によるクエリ,値の範囲は[（present-120d）、present]、最大クエリウィンドウは48時間、ウィンドウシフト範囲は過去120日です。")
	set.StringVar(&a.from, "from", "", "開始照会ID, 注文約定記録ID（最大值）")
	set.StringVar(&a.direct, "direct", "next", "照会方向,約定IDの新着順 default: next, Range: {'prev', 'next'}")
	set.StringVar(&a.size, "size", "10", "記録数, default:100, max: 100")
}

func (a *GetOrderCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	param := url.Values{}
	param.Add("symbol", a.symbol)
	param.Add("states", a.states)
	param.Add("direct", a.direct)
	param.Add("size", a.size)
	if a.types != "" {
		param.Add("types", a.types)
	}
	if a.starttime != "" {
		param.Add("start-time", a.starttime)
	}
	if a.endtime != "" {
		param.Add("end-time", a.endtime)
	}
	if a.from != "" {
		param.Add("from", a.from)
	}

	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/order/orders")+"?"+param.Encode(), nil)
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	h.Process(req)
	return 0
}
