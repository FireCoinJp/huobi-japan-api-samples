package cmds

// マーケット概要

import (
	"context"
	"flag"
	"fmt"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/ws"
	"huobi-japan-api-samples/data/wsRequest"

	"github.com/google/subcommands"
)

type WsMarketDetialCmd struct {
	symbol string
}

func (a *WsMarketDetialCmd) Name() string {
	return "wsMarketDetail"
}

func (a *WsMarketDetialCmd) Synopsis() string {
	return "マーケット概要"
}

func (a *WsMarketDetialCmd) Usage() string {
	return "api-test wsMarketDetail \n"
}

func (a *WsMarketDetialCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "ethbtc", "Pairs")
}

func (a *WsMarketDetialCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	w := ws.NewBuilder()
	channel := fmt.Sprintf("market.%s.detail", a.symbol)

	req := &wsRequest.PublicMarketOverviewBody{
		Req:       channel,
		Id:        "id12",
		IsPrivate: false,
	}
	w.New(req, config.Cfg)
	return 0
}
