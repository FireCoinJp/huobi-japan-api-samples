package cmds

// 注文キャンセル

import (
	"context"
	"flag"
	"fmt"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"

	"github.com/google/subcommands"
)

type SubmitcancelCmd struct {
	orderId string
}

func (a *SubmitcancelCmd) Name() string {
	return "submitcancel"
}

func (a *SubmitcancelCmd) Synopsis() string {
	return "注文キャンセル"
}

func (a *SubmitcancelCmd) Usage() string {
	return "api-test submitcancel \n"
}

func (a *SubmitcancelCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.orderId, "order_id", "376058608563423", "注文ID")
}

func (a *SubmitcancelCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	req, _ := http.NewRequest(http.MethodPost, h.Url(fmt.Sprintf("/v1/order/orders/%s/submitcancel", a.orderId)), nil)
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	h.Process(req)
	return 0
}
