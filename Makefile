build:
		go build -o api-test main.go
run-ws-public:
		./api-test wsBbo
		./api-test wsDepth
		./api-test wsKline
		./api-test wsMarketDetail
		./api-test wsTicker
run-ws-private:
		./api-test wsAccounts
		./api-test wsClearing
		./api-test wsOrder
run-all-account:
		./api-test accounts
		./api-test balance
run-all-wallet:
		./api-test cancel
		./api-test create
run-all-system:
		./api-test currencys
		./api-test symbols
		./api-test timestamp
run-all-market:
		./api-test depth
		./api-test historytrade
		./api-test kline
		./api-test merge
		./api-test tickers
		./api-test trade
run-all-shop:
		./api-test maintainTime
		./api-test orderlist
		./api-test retailPlace
run-all-trade:
		@echo "now show examples"
