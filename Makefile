currency:
		curl -X GET "https://api-cloud.huobi.co.jp/v1/common/currencys" | jq > json/currencys.json
symbol:
		curl -X GET "https://api-cloud.huobi.co.jp/v1/common/symbols" | jq > json/symbols.json
time:
		curl -X GET "https://api-cloud.huobi.co.jp/v1/common/timestamp" | jq > json/time.json
kline:
		curl -X GET "https://api-cloud.huobi.co.jp/market/history/kline?period=1day&size=2&symbol=btcjpy" | jq > json/kline.json
error:
		curl -X GET "https://api-cloud.huobi.co.jp/market/history/kline?period=1day&size=2&symbol=btcjpy1" | jq > json/kline-error.json
ticker:
		curl -X GET "https://api-cloud.huobi.co.jp/market/detail/merged?symbol=ethjpy" | jq > json/tiker.json
tickers:
		curl -X GET "https://api-cloud.huobi.co.jp/market/tickers" | jq > json/tickers.json
depth:
		curl -X GET "https://api-cloud.huobi.co.jp/market/depth?symbol=btcjpy&type=step1" | jq > json/depth.json
trade:
		curl -X GET "https://api-cloud.huobi.co.jp/market/trade?symbol=ethjpy" | jq > json/trade.json
trade_history:
		curl -X GET "https://api-cloud.huobi.co.jp/market/history/trade?symbol=ethjpy" | jq > json/trade_history.json
build:
		go build -o api main.go
get-accounts:
		./api -action accounts
get-balance:
		./api -action balance
get-openOrder:
		./api -action openOrder
get-order:
		./api -action order
get-matchresult:
		./api -action matchresult
get-matchresults:
		./api -action matchresults
get-orders:
		./api -action orders
get-history:
		./api -action order_history
place:
		./api -action place
fee:
		./api -action fee
build:
		go build -o api-test *.go
