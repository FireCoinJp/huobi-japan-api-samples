package cmds

// 入出金記録

import (
	"context"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"
	"net/url"

	"github.com/google/subcommands"
)

type DepositWithdrawCmd struct {
	currency string
	types    string
	from     string
	size     string
}

func (a *DepositWithdrawCmd) Name() string {
	return "depositWithdraw"
}

func (a *DepositWithdrawCmd) Synopsis() string {
	return "入出金記録"
}

func (a *DepositWithdrawCmd) Usage() string {
	return "api-test depositWithdraw \n"
}

func (a *DepositWithdrawCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.currency, "symbol", "btc", "銘柄")
	set.StringVar(&a.types, "types", "deposit", "'deposit' or 'withdraw'")

	set.StringVar(&a.from, "from", "", "開始照会ID, 注文約定記録ID（最大值）")
	set.StringVar(&a.size, "size", "10", "記録数")
}

func (a *DepositWithdrawCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	param := url.Values{}
	param.Add("currency", a.currency)
	param.Add("type", a.types)
	param.Add("size", a.size)
	if a.from != "" {
		param.Add("from", a.from)
	}

	req, _ := http.NewRequest(http.MethodGet, h.Url("/v1/query/deposit-withdraw")+"?"+param.Encode(), nil)
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	h.Process(req)
	return 0
}
