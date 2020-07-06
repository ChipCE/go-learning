

1. Go言語選択理由

- クロスプラットフォーム
- パッケージ管理
- 標準ライブラリが整備されている

Jsonをデコードする
<pre>
import "encoding/json"
//...

// ... 
myJsonString := `{"some":"json"}`
var myStoredVariable
json.Unmarshal([]byte(myJsonString), &myStoredVariable)
//...
</pre>


Jsonをオブジェクト化

<pre>
type TestJsonObj struct {
  Id int `json:id`
  Content interface `json:content`
}
sampleJson := `{"id": "1","content": "{"data1":10,"data2":20}"}`
var testJsonObj TestJsonObj	
json.Unmarshal([]byte(sampleJson), &testJsonObj)
fmt.Printf("Id: %s, Content: %s", testJsonObj.Id, testJsonObj.Content)
</pre>
- 処理速度

## 2. ベンチマーク

### 2.1 ウェブフレームワーク

#### システムスペック

|    |    |
|----|:---|
| Processor | Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz |
| RAM | 15.85 GB |
| OS | Microsoft Windows 10 Pro for Workstations |
| [Bombardier](https://github.com/codesenberg/bombardier) | v1.2.4 |
| [Go](https://golang.org) | go1.14.2 |
| [.Net Core](https://dotnet.microsoft.com/) | 3.1.102 |
| [Node.js](https://nodejs.org/) | v13.10.1 |

> Last updated: Apr 12, 2020 at 2:41am (UTC)

#### Static

リクエスト数　: 1000000
レスポンス: 固定のページ

| Name | Language | Reqs/sec | Latency | Throughput | Time To Complete |
|------|:---------|:---------|:--------|:-----------|:-----------------|
| [Iris](https://github.com/kataras/iris) | Go |205603 |606.83us |35.86MB |4.87s |
| [Kestrel](https://github.com/dotnet/aspnetcore) | C# |194821 |639.91us |33.06MB |5.13s |
| [Chi](https://github.com/pressly/chi) | Go |192223 |648.71us |33.54MB |5.20s |
| [Gin](https://github.com/gin-gonic/gin) | Go |187246 |666.13us |32.67MB |5.34s |
| [Echo](https://github.com/labstack/echo) | Go |185804 |671.59us |32.40MB |5.39s |
| [Martini](https://github.com/go-martini/martini) | Go |160296 |777.36us |27.98MB |6.24s |
| [Koa](https://github.com/koajs/koa) | Javascript |107892 |1.14ms |21.53MB |9.17s |
| [Express](https://github.com/expressjs/express) | Javascript |91181 |1.44ms |22.20MB |11.60s |
| [Buffalo](https://github.com/gobuffalo/buffalo) | Go |39010 |3.20ms |6.81MB |25.63s |

#### Parameterized

リクエスト数　: 250000
レスポンス : 送信のパラメータのデータ表示

| Name | Language | Reqs/sec | Latency | Throughput | Time To Complete |
|------|:---------|:---------|:--------|:-----------|:-----------------|
| [Iris](https://github.com/kataras/iris) | Go |192088 |649.16us |36.78MB |1.30s |
| [Chi](https://github.com/pressly/chi) | Go |187973 |663.14us |36.00MB |1.33s |
| [Echo](https://github.com/labstack/echo) | Go |177471 |703.68us |33.96MB |1.41s |
| [Kestrel](https://github.com/dotnet/aspnetcore) | C# |172411 |728.64us |31.78MB |1.46s |
| [Gin](https://github.com/gin-gonic/gin) | Go |169312 |740.39us |32.27MB |1.48s |
| [Martini](https://github.com/go-martini/martini) | Go |145707 |0.85ms |27.94MB |1.71s |
| [Koa](https://github.com/koajs/koa) | Javascript |74756 |1.67ms |15.99MB |3.35s |
| [Express](https://github.com/expressjs/express) | Javascript |65553 |2.20ms |15.57MB |4.41s |
| [Buffalo](https://github.com/gobuffalo/buffalo) | Go |38577 |3.24ms |7.40MB |6.48s |

#### Sessions

リクエスト数　: 250000
レスポンス : セッション値を設定、その値を表示

| Name | Language | Reqs/sec | Latency | Throughput | Time To Complete |
|------|:---------|:---------|:--------|:-----------|:-----------------|
| [Iris](https://github.com/kataras/iris) | Go |102580 |1.22ms |35.22MB |2.44s |
| [Kestrel](https://github.com/dotnet/aspnetcore) | C# |74056 |1.69ms |34.07MB |3.39s |
| [Echo](https://github.com/labstack/echo) | Go |73521 |1.70ms |38.14MB |3.40s |
| [Chi](https://github.com/pressly/chi) | Go |67068 |1.86ms |34.77MB |3.73s |
| [Martini](https://github.com/go-martini/martini) | Go |66955 |1.86ms |34.81MB |3.73s |
| [Gin](https://github.com/gin-gonic/gin) | Go |60140 |2.10ms |24.95MB |4.20s |
| [Koa](https://github.com/koajs/koa) | Javascript |52796 |2.75ms |20.47MB |5.51s |
| [Express](https://github.com/expressjs/express) | Javascript |31982 |4.27ms |7.82MB |8.56s |
| [Buffalo](https://github.com/gobuffalo/buffalo) | Go |16548 |7.55ms |23.74MB |15.11s |

#### REST

リクエスト数　: 200000
レスポンス : ランダムのJsonを送信、受信したjsonを返す

| Name | Language | Reqs/sec | Latency | Throughput | Time To Complete |
|------|:---------|:---------|:--------|:-----------|:-----------------|
| [Iris](https://github.com/kataras/iris) | Go |147754 |843.88us |40.43MB |1.35s |
| [Chi](https://github.com/pressly/chi) | Go |141918 |0.88ms |38.15MB |1.41s |
| [Kestrel](https://github.com/dotnet/aspnetcore) | C# |136747 |0.92ms |39.68MB |1.47s |
| [Gin](https://github.com/gin-gonic/gin) | Go |136480 |0.91ms |37.32MB |1.47s |
| [Echo](https://github.com/labstack/echo) | Go |134209 |0.93ms |36.84MB |1.49s |
| [Martini](https://github.com/go-martini/martini) | Go |123638 |1.01ms |33.22MB |1.62s |
| [Buffalo](https://github.com/gobuffalo/buffalo) | Go |56722 |2.20ms |15.56MB |3.53s |
| [Koa](https://github.com/koajs/koa) | Javascript |47089 |2.66ms |13.91MB |4.26s |
| [Express](https://github.com/expressjs/express) | Javascript |41018 |3.29ms |13.59MB |5.28s |


参考資料:https://github.com/kataras/server-benchmarks/blob/master/README.md

### 2.2 Go - Iris vs .NetCore - Kestrel


#### システムスペック

|    |    |
|----|:---|
| Processor | Intel(R) Core(TM) i7–4710HQ CPU @ 2.50GHz 2.50GHz |
| RAM | 8GB GB |
| OS | Microsoft Windows 10 Pro for Workstations |
| [Bombardier](https://github.com/codesenberg/bombardier) | v1.2.4 |
| [Go](https://golang.org) | go1.8.3 |
| [.Net Core](https://dotnet.microsoft.com/) | 2.0 |

#### HTTPリクエスト
<code>Kestrel</code>
<pre>
$ bombardier -c 125 -n 5000000 http://localhost:5000/api/values/5
Bombarding http://localhost:5000/api/values/5 with 5000000 requests using 125 connections
 5000000 / 5000000 [=====================================================] 100.00% 2m3s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec     40226.03    8724.30     161919
  Latency        3.09ms     1.40ms   169.12ms
  HTTP codes:
    1xx - 0, 2xx - 5000000, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:     8.91MB/s
</pre>

<code>Iris</code>
<pre>
$ bombardier -c 125 -n 5000000 http://localhost:5000/api/values/5
Bombarding http://localhost:5000/api/values/5 with 5000000 requests using 125 connections
 5000000 / 5000000 [======================================================] 100.00% 47s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec    105643.81    7687.79     122564
  Latency        1.18ms   366.55us    22.01ms
  HTTP codes:
    1xx - 0, 2xx - 5000000, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:    19.65MB/s
</pre>


#### MVCテンプレート

<code>Kestrel</code>
<pre>
Bombarding http://localhost:5000 with 1000000 requests using 125 connections
 1000000 / 1000000 [====================================================] 100.00% 1m20s
Done!
Statistics Avg Stdev Max
 Reqs/sec 11738.60 7741.36 125887
 Latency 10.10ms 22.10ms 1.97s
 HTTP codes:
 1xx — 0, 2xx — 1000000, 3xx — 0, 4xx — 0, 5xx — 0
 others — 0
 Throughput: 89.03MB/s
</pre>

<code>Iris</code>
<pre>
Bombarding http://localhost:5000 with 1000000 requests using 125 connections
 1000000 / 1000000 [======================================================] 100.00% 37s
Done!
Statistics Avg Stdev Max
 Reqs/sec 26656.76 1944.73 31188
 Latency 4.69ms 1.20ms 22.52ms
 HTTP codes:
 1xx — 0, 2xx — 1000000, 3xx — 0, 4xx — 0, 5xx — 0
 others — 0
 Throughput: 192.51MB/s
</pre>

参考資料:https://medium.com/hackernoon/go-vs-net-core-in-terms-of-http-performance-7535a61b67b8