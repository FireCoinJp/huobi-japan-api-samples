package cmds

// 条件付き注文の一括キャンセル

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"huobi-japan-api-samples/data/request"
	"net/http"

	"github.com/google/subcommands"
)

type BatchCancelOpenOrdersCmd struct {
	symbol string
	side   string
	size   string
	types  string
}

func (a *BatchCancelOpenOrdersCmd) Name() string {
	return "batchCancelOpenOrders"
}

func (a *BatchCancelOpenOrdersCmd) Synopsis() string {
	return "条件付き注文の一括キャンセル"
}

func (a *BatchCancelOpenOrdersCmd) Usage() string {
	return "api-test batchCancelOpenOrders \n"
}

func (a *BatchCancelOpenOrdersCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.symbol, "symbol", "", "取引通貨ペア")
	set.StringVar(&a.side, "side", "", "取引方向 , Range: {“buy”,“sell”}， デフォルトでは、条件が満たされていない全ての注文が返されます。")
	set.StringVar(&a.size, "size", "", "必要な記録数, default: 100, Range: {0,100}")
	set.StringVar(&a.types, "types", "", "カンマで区切られた注文タイプの組み合わせ")
}

func (a *BatchCancelOpenOrdersCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)
	sendBody := request.BatchCancelOpenOrdersBody{
		AccountId: config.Cfg.AccountID,
		Symbol:    a.symbol,
		Side:      a.side,
		Size:      a.size,
		Types:     a.types,
	}

	batchCancelOpenOrdersBody, _ := json.Marshal(sendBody)

	req, _ := http.NewRequest(http.MethodPost, h.Url("/v1/order/orders/batchCancelOpenOrders"), bytes.NewReader(batchCancelOpenOrdersBody))
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}

	h.Process(req)
	return 0
}
