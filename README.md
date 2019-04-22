## httpserver

httpserver is sample http server swapper of fasthttp and fasthttproute, you can use it to easily create http server. it's 
dependence fasthttp and fasthttproute.

## usage 

```
go get github.com/huangjiasingel/httpserver
```

## example

```
import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/huangjiasingle/httpserver"
	"github.com/valyala/fasthttp"
)

func main() {
	route := fasthttprouter.New()
	server := httpserver.New(route, "/api/v1")
	registry(server)
	log.Fatal(server.ServerHTTP(":8080"))
}

func Ping(ctx *fasthttp.RequestCtx) {
	name := ctx.QueryArgs().Peek("name")
	fmt.Fprintf(ctx, "Pong! %s\n", string(name))
}

func Pong(ctx *fasthttp.RequestCtx) {
	name := ctx.QueryArgs().Peek("name")
	fmt.Fprintf(ctx, "ping! %s\n", string(name))
}

func registry(server *httpserver.Server) {
	server.Registry("/ping", "GET", Ping).Registry("/pong", "POST", Pong)
}

```