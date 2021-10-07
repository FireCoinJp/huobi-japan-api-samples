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
	isSave  bool
	orderId string
}

func (a *SubmitcancelCmd) Name() string {
	return "submitcancel"
}

func (a *SubmitcancelCmd) Synopsis() string {
	return "SubmitcancelCmd"
}

func (a *SubmitcancelCmd) Usage() string {
	return "api-test Submitcancel -save"
}

func (a *SubmitcancelCmd) SetFlags(set *flag.FlagSet) {
	set.BoolVar(&a.isSave, "save", false, "write to json")
	set.StringVar(&a.orderId, "order_id", "376058608563423", "order_id success")
	return
}

func (a *SubmitcancelCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	req, _ := http.NewRequest(http.MethodPost, h.Url(fmt.Sprintf("/v1/order/%s/submitcancel", a.orderId)), nil)
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	if a.isSave {
		err = h.Do(req, api.SaveMsg)
	} else {
		err = h.Do(req, api.PrintMsg)
	}

	if err != nil {
		panic(err)
	}
	return 0
}
