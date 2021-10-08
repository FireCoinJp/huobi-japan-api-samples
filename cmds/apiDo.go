package cmds

import (
	"huobi-japan-api-samples/config"
	"huobi-japan-api-samples/core/api"
	"net/http"
)

func apiDo(req *http.Request, save bool) {
	h := api.New(config.Cfg)
	err := h.Do(req, api.PrintMsg)
	if save {
		err = h.Do(req, api.SaveMsg)
	}

	if err != nil {
		panic(err)
	}
}
