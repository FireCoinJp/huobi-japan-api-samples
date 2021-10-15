package cmds

// BBO

import (
	"context"
	"flag"
	"fmt"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/ws"
	"huobi-japan-api-samples/data/wsRequest"

	"github.com/google/subcommands"
)

type WsBboCmd struct {
	symbol string
}

func (a *WsBboCmd) Name() string {
	return "wsBbo"
}

func (a *WsBboCmd) Synopsis() string {
	return "BBO"
}

func (a *WsBboCmd) Usage() string {
	return "api-test wsBbo \n"
}

func (a *WsBboCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "ethbtc", "取引ペア	")
}

func (a *WsBboCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	w := ws.NewBuilder()
	channel := fmt.Sprintf("market.%s.bbo", a.symbol)

	sub := &wsRequest.PublicMarketBody{
		Sub:       channel,
		Id:        "id1",
		IsPrivate: true,
	}
	w.New(sub, config.Cfg)

	return 0
}
