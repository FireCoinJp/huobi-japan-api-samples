package cmds

// 注文の照会

import (
	"context"
	"flag"
	"fmt"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"

	"github.com/google/subcommands"
)

type OrderCmd struct {
	orderId string
}

func (a *OrderCmd) Name() string {
	return "order"
}

func (a *OrderCmd) Synopsis() string {
	return "注文の照会"
}

func (a *OrderCmd) Usage() string {
	return "api-test order \n"
}

func (a *OrderCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.orderId, "order_id", "375977348044411", "注文ID")
}

func (a *OrderCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	req, _ := http.NewRequest(http.MethodGet, h.Url(fmt.Sprintf("/v1/order/orders/%s", a.orderId)), nil)
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	h.Process(req)
	return 0
}
