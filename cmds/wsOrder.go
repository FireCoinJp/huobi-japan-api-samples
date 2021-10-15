package cmds

// 注文データ

import (
	"context"
	"flag"
	"fmt"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/ws"
	"huobi-japan-api-samples/data/wsRequest"

	"github.com/google/subcommands"
)

type WsOrderCmd struct {
	symbol string
}

func (a *WsOrderCmd) Name() string {
	return "wsOrder"
}

func (a *WsOrderCmd) Synopsis() string {
	return "注文データ"
}

func (a *WsOrderCmd) Usage() string {
	return "api-test wsOrder \n"
}

func (a *WsOrderCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "*", "取引ペア（ワイルドカード　*　使用可）")
}

func (a *WsOrderCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	channel := fmt.Sprintf("orders#%s", a.symbol)
	sub := &wsRequest.PrivateRequest{
		Action:    "sub",
		Ch:        channel,
	}
	w := ws.NewBuilder(config.Cfg, sub).Build()
	w.Run(config.Cfg.Timeout)
	return 0
}
