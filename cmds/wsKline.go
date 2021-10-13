package cmds

// ローソク足 データ

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/ws"

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
	set.StringVar(&a.symbol, "symbol", "btcusdt", "取引ペア - btceth")
	set.StringVar(&a.period, "period", "1min", "チャートタイプ")
}

func (a *WsKLineCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {

	urlpath := "wss://api.huobi.pro/ws"
	sub := []string{`{"id": "id1", "sub": "market.` + a.symbol + `.kline.` + a.period + `"}`}

	ws.NewBuilder().New(sub, urlpath, *config.Cfg)

	return 0
}
