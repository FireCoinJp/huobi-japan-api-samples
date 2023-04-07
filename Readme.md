# Huobi Japan API Sample

> Huobi Japan API のサンプルコード。このサンプルを使ってAPIの機能を検証することができます。
>
> 関連コードを参照することで、自動取引機能を実現できるプログラムになります。



[API DOC](https://api-doc.bittrade.co.jp/#api)

## API Keyの取得方法

> ログイン後、下図の通りAPIの秘密鍵を取得することができます。

![image-20211106095122808](.asset/image-20211106095122808.png)


## プログラムの作成について

> 以下のコマンドによって、各OSに適用のコマンドを作成できます。

```bash
# clone code
git clone git@github.com:FireCoinJp/huobi-japan-api-samples.git

# mac
$ go build -o api-test main.go

# linux 
$ GOOS=linux GOARCH=amd64 go build -o api-test main.go

# windows
$ GOOS=windows GOARCH=amd64 go build -o api-test main.go
```

## 開発環境の構築について

> コマンドを実行するフォルダーの中で、`config.yaml`ファイルを作成し、下記の通りコマンドを実行してください。

 1. ルートディレクトリでconfig.yamlファイルを作成し、下記内容を入力します。

 ```bash
 $ mv config.yaml.sample config.yaml
 $ vi config.yaml
 
 setting:
 	access_key: xxxxxx          # 公開鍵
 	secret_key: xxxxxx          # 秘密鍵
 	account_id: 12345678        # accountsサブコマンドを使って取得できる
 	host: api-cloud.bittrade.co.jp # 固定値
 	save: false                 # 結果を出力するか
 	timeout: 10s                # WSの実行する時間（デフォルト値は10秒）
 ```

 2. account_idの取得方法

    > 構成ファイルにて秘密鍵を記入する必要があります。

```bash
$ ./api-test accounts

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

## 実行について

> 本プログラムはサブコマンドモードです。使い方は下記の通りになります。


```bash
# コマンドリストを参照

$ ./api-test
Usage: api-test <flags> <subcommand> <subcommand args>

Subcommands:
        commands         list all command names
        flags            describe all known top-level flags
        help             describe subcommands and their syntax

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
        wsDataSubscription  データ購読

   
   
# サブコマンドパラメータを参照
$ ./api-test help kline
api-test kline 
  -period string
        チャートタイプ (default "1day")
  -size string
        サイズ - default=150 max=2000  (default "20")
  -symbol string
        取引ペア - btceth (default "btcjpy")

# コマンドを実行
$ ./api-test kline -symbol btcjpy -size 2
{
  "ch": "market.btcjpy.kline.1day",
  "data": [
    {
      "amount": 2.6079233285649943,
      "close": 6903272,
      "count": 13332,
      "high": 6958855,
      "id": 1636128000,
      "low": 6884737,
      "open": 6893137,
      "vol": 18044270.39024
    },
    {
      "amount": 23.07694295966621,
      "close": 6891578,
      "count": 37439,
      "high": 7189999,
      "id": 1636041600,
      "low": 6891578,
      "open": 6938545,
      "vol": 161380835.70746
    }
  ],
  "status": "ok",
  "ts": 1636157137864
}

```



## 認証ロジック (API & WebSocket)

> ユーザ認証とセキュリティのため、プライベートAPIをアクセスする際には署名が必要となります。下記コマンドにて署名に関する手順を説明します。

+ Rest API

	>  プライベートAPIをアクセスする前に署名を作成し、既存パラメータでAPIをアクセスします

```go
// file: core/api/http.go 
func (c Client) Auth(req *http.Request) error {
    authParams := url.Values{}
    // GET API contains request parameters, POST API only contains URI
    if req.Method == http.MethodGet {
        authParams, _ = url.ParseQuery(req.URL.RawQuery)
    }

    // Add fixed parameters
    authParams.Set("AccessKeyId", c.config.AccessKey)
    authParams.Set("SignatureMethod", "HmacSHA256")
    authParams.Set("SignatureVersion", "2")
    authParams.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05"))
  
    // make a signature
    s := fmt.Sprintf("%s\n%s\n%s\n%s", req.Method, req.URL.Host, req.URL.Path, authParams.Encode())
    signature := crypto.Hmac256(s, c.config.SecretKey)
    authParams.Set("Signature", signature)
    req.URL, _ = url.Parse(fmt.Sprintf("%s://%s%s?%s", req.URL.Scheme, req.URL.Host, req.URL.Path, authParams.Encode()))
    
    req.Header.Set("Content-Type", "application/json")
    return nil
}		
```

+ WebSocket

  > 接続後、署名コマンドを実行し、websocketプライベートコマンドを購読や起動します。署名の手順はRest APIと同様になります。

```go
// file : core/ws/websocket.go 
func (w *Client) handleAuth() error {
    // add fixed parameters
    authParams := url.Values{}
    utc := time.Now().UTC().Format("2006-01-02T15:04:05")
    authParams.Set("accessKey", w.config.accessKey)
    authParams.Set("signatureMethod", "HmacSHA256")
    authParams.Set("signatureVersion", "2.1")
    authParams.Set("timestamp", utc)
    
    // make a signature
    host := "api-cloud.bittrade.co.jp"
    path := "/ws/v2"
    s := fmt.Sprintf("GET\n%s\n%s\n%s", host, path, authParams.Encode())
    signature := crypto.Hmac256(s, w.config.secretKey)

    // build request json
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
    
    // send command
    return w.WriteMessage(websocket.TextMessage, authBody)
}
```


## ディレクトリ構造

```bash
# 本プロジェクトのディレクトリ構造
$ tree -L 1 

├── Makefile            # デフォルトコマンド集
├── Readme.md           # 説明書
├── api-test            # 実行可能ファイル
├── cmds                # コマンド
├── config              # 構成定義
├── config.yaml.sample  # 構成ファイル（ユーザで生成する必要がある）
├── core                # Library
├── data                # データ構造の定義
├── json                # 結果の保存
└── main.go             # main関数
```

