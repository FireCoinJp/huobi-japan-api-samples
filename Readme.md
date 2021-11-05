## **Huobi Japan API Sample**

Huobi Japan API 使用说明

[toc]

### 使用前的配置要求

​	1.在项目根目录创建config.yaml文件。

​	2.粘贴以下内容到config.yaml文件内。

```bash
setting:
  access_key: 
  secret_key: 
  account_id: 
  host: api-cloud.huobi.co.jp
  save: false
  timeout: 10s
```
​	3.access_key和secret_key取自huobi.co.jp个人账户内，在火币日本登陆个人账号后点击 https://www.huobi.co.jp/ja-jp/user/api/ 链接，作成一个有读取、提现、交易的权限密钥，作成完成时的弹框内的公钥为access_key，密钥为secret_key，请复制公钥和密钥的值到config.yaml内对应的access_key和secret_key处。

​	4.account_id为用户账户ID，在access_key和secret_key填写完成的前提下，取值方法为在Terminal中输入`go run main.go accounts`并运行。在save值为false的情况下，在Terminal可以看到以下数据，id的值就是当前账号的account_id。在save值为ture的情况下，json目录下会生成一个accounts.json文件，打开文件会看到以下数据，id的值就是当前账号的account_id。请复制id的值到config.yaml内对应的account_id处。

```bash
{
  "data": [
    {
      "id": 1234567,
      "state": "working",
      "subtype": "",
      "type": "spot"
    }
  ],
  "status": "ok"
}
```

​	5.save值的作用是判定是否保存程序运行后得到的数据，若填写值为ture，程序运行后会则会在json目录下生成相应的xxx.json文件且在Terminal不输出数据，若填写值为false，程序运行后则会在Terminal输出运行程序后的数据且json文件夹下不生成相应的xxx.json文件。

​	6.timeout值的作用是控制websocket api的运行时间，若填写时间为10s，被执行的websocket程序则会在运行10s后停止。



## 如何创建命令



```bash
# clone code
git clone git@github.com:FireCoinJp/huobi-japan-api-samples.git

$ cd 进入项目目录

# mac

go build -o api-test main.go

# linux 

 GOOS=linux GOARCH=amd64 go build -o api-test main.go

# windows

 GOOS=windows GOARCH=amd64 go build -o api-test main.go
```



### 命令说明


```bash
# 查询命令行

$ ./api-test

Subcommands for アカウント関連:
        accounts         ユーザアカウント
        balance          残高照合

Subcommands for ウォレット関連:
        cancel           暗号資産の出金のキャンセル
        create           暗号資産の出金申請
        depositWithdraw  入出金記録

Subcommands for システム情報関連:
        currencys        対応取引通貨
        symbols          取引ペア情報
        timestamp        システム時間を調べる

Subcommands for マーケット関連:
        depth            板情報
        historytrade     取引履歴の取得
        kline            ローソク足
        merge            ティッカー
        tickers          全取引ペアの相場情報
        trade            最新の取引データ

Subcommands for 取引関連:
        batchCancelOpenOrders  条件付き注文の一括キャンセル
        batchcancel      注文の一括キャンセル
        getMatchresults  約定履歴の検索
        getOrder         注文履歴の検索
        matchresults     注文の約定詳細
        openOrders       未約定注文一覧
        order            注文の照会
        place            注文実行
        submitcancel     注文キャンセル

Subcommands for 販売所関連:
        maintainTime     販売所メンテナンス時間
        orderlist        販売所注文履歴
        retailPlace      販売所での注文

Subcommands for Websocket (Private):
        wsAccounts       資産変動
        wsClearing       注文状態更新
        wsOrder          注文データ

Subcommands for Websocket (Public):
        wsBbo            BBO
        wsDepth          板情報
        wsKline          ローソク足 データ
        wsMarketDetail   マーケット概要
        wsTicker         ティッカー
   
   
# 如需查询子命令参数
$ ./api-test help kline

```

**例**

​	2件btcjpy的取引履歴の取得步骤为

​	在Terminal中输入 ./api-test historytrade -symbol btcjpy -size 2 

​	返回值为
```bash
	{
  "ch": "market.btcjpy.trade.detail",
  "data": [
    {
      "data": [
        {
          "amount": 0.00001,
          "direction": "sell",
          "id": 1.0102043633540256e+26,
          "price": 7116558,
          "trade-id": 100029930874,
          "ts": 1636010851910
        },
        {
          "amount": 0.00001,
          "direction": "sell",
          "id": 1.0102043633540256e+26,
          "price": 7116661,
          "trade-id": 100029930873,
          "ts": 1636010851910
        },
        {
          "amount": 0.00001,
          "direction": "sell",
          "id": 1.0102043633540256e+26,
          "price": 7116713,
          "trade-id": 100029930872,
          "ts": 1636010851910
        },
        {
          "amount": 0.00001,
          "direction": "sell",
          "id": 1.0102043633540256e+26,
          "price": 7116777,
          "trade-id": 100029930871,
          "ts": 1636010851910
        },
        {
          "amount": 0.00002,
          "direction": "sell",
          "id": 1.0102043633540256e+26,
          "price": 7116858,
          "trade-id": 100029930870,
          "ts": 1636010851910
        }
      ],
      "id": 101020436335,
      "ts": 1636010851910
    },
    {
      "data": [
        {
          "amount": 0.00001,
          "direction": "sell",
          "id": 1.0102043613940257e+26,
          "price": 7116547,
          "trade-id": 100029930869,
          "ts": 1636010848886
        },
        {
          "amount": 0.00001,
          "direction": "sell",
          "id": 1.0102043613940257e+26,
          "price": 7116611,
          "trade-id": 100029930868,
          "ts": 1636010848886
        },
        {
          "amount": 0.00001,
          "direction": "sell",
          "id": 1.0102043613940257e+26,
          "price": 7116737,
          "trade-id": 100029930867,
          "ts": 1636010848886
        },
        {
          "amount": 0.00001,
          "direction": "sell",
          "id": 1.0102043613940257e+26,
          "price": 7116745,
          "trade-id": 100029930866,
          "ts": 1636010848886
        },
        {
          "amount": 0.00002,
          "direction": "sell",
          "id": 1.0102043613940257e+26,
          "price": 7116819,
          "trade-id": 100029930865,
          "ts": 1636010848886
        }
      ],
      "id": 101020436139,
      "ts": 1636010848886
    }
  ],
  "status": "ok",
  "ts": 1636010866794
}
```

​	在Terminal中输入 ./api-test help historytrade并执行，则会取得执行取引履歴の取得命令所需的全部参数及备注
```bash
api-test historytrade 
  -size string
        データサイズ, Range: {1, 2000} (default "2")
  -symbol string
        取引ペア, 例えば btcjpy (default "btcjpy")
```

### 验签逻辑

##### 		http端

```go
# path : huobi-japan-api-samples/core/api/http.go 

func (c Client) Auth(req *http.Request) error {

    authParams := url.Values{}

    if req.Method == http.MethodGet {

        authParams, _ = url.ParseQuery(req.URL.RawQuery)

    }

    authParams.Set("AccessKeyId", c.config.AccessKey)

    authParams.Set("SignatureMethod", "HmacSHA256")

    authParams.Set("SignatureVersion", "2")

    authParams.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05"))

    s := fmt.Sprintf("%s\n%s\n%s\n%s", req.Method, req.URL.Host, req.URL.Path, authParams.Encode())

    signature := crypto.Hmac256(s, c.config.SecretKey)

    authParams.Set("Signature", signature)

    req.URL, _ = url.Parse(fmt.Sprintf("%s://%s%s?%s", req.URL.Scheme, req.URL.Host, req.URL.Path, authParams.Encode()))

    req.Header.Set("Content-Type", "application/json")

    return nil

}		
```
##### 		websocket端

```go
# path : huobi-japan-api-samples/core/ws/websocket.go 

func (w *Client) handleAuth() error {

    authParams := url.Values{}

    utc := time.Now().UTC().Format("2006-01-02T15:04:05")

    authParams.Set("accessKey", w.config.accessKey)

    authParams.Set("signatureMethod", "HmacSHA256")

    authParams.Set("signatureVersion", "2.1")

    authParams.Set("timestamp", utc)

    host := "api-cloud.huobi.co.jp"

    path := "/ws/v2"

    s := fmt.Sprintf("GET\n%s\n%s\n%s", host, path, authParams.Encode())

    signature := crypto.Hmac256(s, w.config.secretKey)

    param := wsRequest.Param{

        AuthType:         "api",

        AccessKey:        w.config.accessKey,

        SignatureMethod:  "HmacSHA256",

        SignatureVersion: "2.1",

        Timestamp:        utc,

        Signature:        signature,

    }

    auth := wsRequest.AuthJson{

        Action: "req",

        Ch:     "auth",

        Params: param,

    }



    authBody, _ := json.Marshal(auth)

    return w.WriteMessage(websocket.TextMessage, authBody)

}
```


### 目录结构说明
```bash
.

├── Makefile 			// 编译规则
├── Readme.md			// 说明文档
├── api-test  			// 可执行文件
├── cmds    			// 运行代码
├── config			// 配置
├── config.yaml			// 配置文件
├── core			// Library
├── data			// 结构体存储位置
├── go.mod			// 包依赖
├── go.sum			// 包依赖
├── json			// response生成的.json文件的存储位置
└── main.go					
```
