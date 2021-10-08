package cmds

// 約定履歴の検索

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"
	"net/url"

	"github.com/google/subcommands"
)

type GetMatchresultsCmd struct {
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

func (a *GetMatchresultsCmd) Name() string {
	return "GetMatchresults"
}

func (a *GetMatchresultsCmd) Synopsis() string {
	return "約定履歴の検索"
}

func (a *GetMatchresultsCmd) Usage() string {
	return "api-test GetMatchresultsCmd -save"
}

func (a *GetMatchresultsCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "btcjpy", "取引通貨ペア")
	set.StringVar(&a.states, "states", "filled", "オーダーのタイプの組み合わせ照会，区切り記号は','を使用。[submitted 提出済み, partial-filled 部分約定, partial-canceled 部分約定キャンセル, filled 完全約定, canceled キャンセル済み]")

	set.StringVar(&a.types, "types", "", "オーダータイプの組み合わせ照会，複数可, カンマ区切り")
	set.StringVar(&a.startdate, "start_date", "", "開始日の照会, 日時フォマットyyyy-mm-dd, Range: [-61日, Now]")
	set.StringVar(&a.enddate, "end_date", "", "終了日の照会, 日時フォマットyyyy-mm-dd")
	set.StringVar(&a.from, "from", "", "開始ID")
	set.StringVar(&a.direct, "direct", "next", "照会方向,約定IDの新着順 default: next, Range: {'prev', 'next'}")
	set.StringVar(&a.size, "size", "10", "記録数, Range: [0, 100]")
	set.BoolVar(&a.isSave, "save", false, "write to json")
}

func (a *GetMatchresultsCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	param := url.Values{}
	param.Add("symbol", a.symbol)
	param.Add("states", a.states)
	param.Add("direct", a.direct)
	param.Add("size", a.size)

	if a.types != "" {
		param.Add("types", a.types)
	}
	if a.startdate != "" {
		param.Add("start_date", a.startdate)
	}
	if a.enddate != "" {
		param.Add("end_date", a.enddate)
	}
	if a.from != "" {
		param.Add("from", a.from)
	}

	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/order/matchresults")+"?"+param.Encode(), nil)
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	apiDo(req, a.isSave)
	return 0
}
