package cmds

// ローソク足 データ

import (
	"context"
	"flag"
	"fmt"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/ws"
	"huobi-japan-api-samples/data/wsRequest"

	"github.com/google/subcommands"
)

type WsKLineCmd struct {
	symbol string
	period string
}

func (a *WsKLineCmd) Name() string {
	return "wsKline"
}

func (a *WsKLineCmd) Synopsis() string {
	return "ローソク足 データ"
}

func (a *WsKLineCmd) Usage() string {
	return "api-test wsKline \n"
}

func (a *WsKLineCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "ethbtc", "取引ペア - btceth")
	set.StringVar(&a.period, "period", "1min", "チャートタイプ")
}

func (a *WsKLineCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	w := ws.NewBuilder()
	channel := fmt.Sprintf("market.%s.kline.%s", a.symbol, a.period)

	sub := &wsRequest.PublicMarketBody{
		Sub:       channel,
		Id:        "id1",
		IsPrivate: false,
	}
	w.New(sub, config.Cfg)
	return 0
}
