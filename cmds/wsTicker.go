package cmds

// ティッカー

import (
	"context"
	"flag"
	"fmt"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/ws"
	"huobi-japan-api-samples/data/wsRequest"

	"github.com/google/subcommands"
)

type WsTickerCmd struct {
	symbol string
}

func (a *WsTickerCmd) Name() string {
	return "wsTicker"
}

func (a *WsTickerCmd) Synopsis() string {
	return "ティッカー"
}

func (a *WsTickerCmd) Usage() string {
	return "api-test wsTicker \n"
}

func (a *WsTickerCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "ethbtc", "Pairs")
}

func (a *WsTickerCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	channel := fmt.Sprintf("market.%s.trade.detail", a.symbol)
	sub := &wsRequest.PublicRequest{
		Sub:       channel,
		Id:        "id1",
	}
	w := ws.NewBuilder(config.Cfg, sub).Build()
	w.Run(config.Cfg.Timeout)
	return 0
}
