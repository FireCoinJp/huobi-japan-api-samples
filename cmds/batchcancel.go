package cmds

// 注文の一括キャンセル

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"huobi-japan-api-samples/data/request"
	"net/http"
	"strings"

	"github.com/google/subcommands"
)

type BatchcancelCmd struct {
	orderIds       string
	clientOrderIds string
	isSave         bool
}

func (a *BatchcancelCmd) Name() string {
	return "batchcancel"
}

func (a *BatchcancelCmd) Synopsis() string {
	return "注文の一括キャンセル"
}

func (a *BatchcancelCmd) Usage() string {
	return "api-test BatchcancelCmd -save"
}

func (a *BatchcancelCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&a.orderIds, "order-ids", "", "注文番号リスト")
	set.StringVar(&a.clientOrderIds, "client-order-ids", "", "ユーザ定義された注文番号")
	set.BoolVar(&a.isSave, "save", false, "write to json")
}

func (a *BatchcancelCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	h := api.New(config.Cfg)

	var orderId_arr []string
	var clientOrderIds_arr []string
	if a.orderIds != "" {
		orderId_arr = strings.Split(a.orderIds, " ")
	}
	if a.clientOrderIds != "" {
		clientOrderIds_arr = strings.Split(a.clientOrderIds, " ")
	}

	sendBody := request.BatchcancelBody{
		OrderIds:       orderId_arr,
		ClientOrderIds: clientOrderIds_arr,
	}

	BatchcancelBody, _ := json.Marshal(sendBody)

	req, _ := http.NewRequest(http.MethodPost, h.Url("/v1/order/orders/batchcancel"), bytes.NewReader(BatchcancelBody))
	err := h.Auth(req)
	if err != nil {
		panic(err)
	}
	apiDo(req, a.isSave)
	return 0
}
