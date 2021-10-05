# Huobi Japan API Sample
----

### 功能需求

+ 支持sub command

```bash
$ hbj-api-test get-kline -symbol btcjpy

{
  "ch": "market.btcjpy.kline.1day",
  "status": "ok",
  "ts": 1632969868010,
  "data": [
    {
      "id": 1632931200,
      "open": 4663135,
      "close": 4815018,
      "low": 4580081,
      "high": 4830000,
      "amount": 6.634525507825414,
      "vol": 31208129.2684,
      "count": 7534
    },
    {
      "id": 1632844800,
      "open": 4621790,
      "close": 4662895,
      "low": 4550403,
      "high": 4763039,
      "amount": 32.57668369170352,
      "vol": 152697278.3816,
      "count": 18230
    }
  ]
}

```

+ 支持websocket订阅数据， 输出10秒钟左右的接受结果

```bash
hbj-api-test ws-kline -symbol btcjpy

// 输出10s左右的结果
```

+ 支持私有API的访问， 并返回结果

```bash
hbj-api-test order-history -f config.yaml //
```

### 模块设计

```bash

├── Makefile     //
├── Readme.md
├── api          //
├── command      //
├── config       //
├── config.yaml
├── core         // 共通函数
├── data         // request, response
├── json         // データの格納場所
└── main.go      // 入り口


```

### 依赖库

+ https://github.com/google/subcommands
+
