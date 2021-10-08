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

type CancelCmd struct {
	isSave     bool
	withdrawId string
}

func (a *CancelCmd) Name() string {
	return "cancel"
}

func (a *CancelCmd) Synopsis() string {
	return "暗号資産の出金のキャンセル"
}

func (a *CancelCmd) Usage() string {
	return "api-test CancelCmd -save"
}

func (a *CancelCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.withdrawId, "withdraw-id", "75705660", "出金ID，pathの中に記入")
	set.BoolVar(&a.isSave, "save", false, "write to json")
}

func (a *CancelCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	req, _ := http.NewRequest(http.MethodPost, h.Url(fmt.Sprintf("/v1/dw/withdraw-virtual/%s/cancel", a.withdrawId)), nil)
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	apiDo(req, a.isSave)
	return 0
}
