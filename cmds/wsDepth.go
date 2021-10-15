package cmds

// 板情報

import (
	"context"
	"flag"
	"fmt"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/ws"
	"huobi-japan-api-samples/data/wsRequest"

	"github.com/google/subcommands"
)

type WsDepthCmd struct {
	symbol    string
	depthType string
}

func (a *WsDepthCmd) Name() string {
	return "wsDepth"
}

func (a *WsDepthCmd) Synopsis() string {
	return "板情報"
}

func (a *WsDepthCmd) Usage() string {
	return "api-test wsDepth \n"
}

func (a *WsDepthCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "ethbtc", "Pairs")
	set.StringVar(&a.depthType, "type", "step1", "Market depth")
}

func (a *WsDepthCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	w := ws.NewBuilder()
	channel := fmt.Sprintf("market.%s.depth.%s", a.symbol, a.depthType)

	sub := &wsRequest.PublicMarketBody{
		Sub:       channel,
		Id:        "id1",
		IsPrivate: false,
	}
	w.New(sub, config.Cfg)
	return 0
}
