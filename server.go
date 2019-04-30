package httpserver

import (
	"fmt"
	"time"

	"k8s.io/klog"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type Server struct {
	router *router.Router
	prefix string
}

func New(router *router.Router, prefix string) *Server {
	return &Server{router, prefix}
}

func (s *Server) Registry(path, method string, handle fasthttp.RequestHandler) *Server {
	path = fmt.Sprintf("%v%v", s.prefix, path)
	switch method {
	case "GET":
		s.router.GET(path, handle)
	case "HEAD":
		s.router.HEAD(path, handle)
	case "POST":
		s.router.POST(path, handle)
	case "PUT":
		s.router.PUT(path, handle)
	case "DELETE":
		s.router.DELETE(path, handle)
	case "OPTIONS":
		s.router.OPTIONS(path, handle)
	case "PATCH":
		s.router.PATCH(path, handle)
	}
	return s
}

func (s *Server) Handler(ctx *fasthttp.RequestCtx) {
	t := time.Now()
	s.router.Handler(ctx)
	sub := time.Now().Sub(t)
	klog.V(2).Infof("%s  %v  %s", ctx.String(), ctx.Response.StatusCode(), sub)
}

func (s *Server) ServerHTTP(address string) error {
	return fasthttp.ListenAndServe(address, s.Handler)
}
