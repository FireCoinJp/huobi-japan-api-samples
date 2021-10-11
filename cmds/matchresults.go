package cmds

// 注文の約定詳細

import (
	"context"
	"flag"
	"fmt"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"

	"github.com/google/subcommands"
)

type MatchresultsCmd struct {
	orderId string
}

func (a *MatchresultsCmd) Name() string {
	return "matchresults"
}

func (a *MatchresultsCmd) Synopsis() string {
	return "注文の約定詳細"
}

func (a *MatchresultsCmd) Usage() string {
	return "api-test matchresults \n"
}

func (a *MatchresultsCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.orderId, "order_id", "375977348044411", "パスに記載された注文ID")
}

func (a *MatchresultsCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	req, _ := http.NewRequest(http.MethodGet, h.Url(fmt.Sprintf("/v1/order/orders/%s/matchresults", a.orderId)), nil)
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	h.Process(req)
	return 0
}
