package httpserver

import (
	"fmt"
	"time"

	"k8s.io/klog"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type Server struct {
	router *fasthttprouter.Router
	prefix string
}

func New(router *fasthttprouter.Router, prefix string) *Server {
	return &Server{router, prefix}
}

func (s *Server) Registry(path, method string, handle fasthttp.RequestHandler) {
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
}

func (s *Server) Handler(ctx *fasthttp.RequestCtx) {
	t := time.Now()
	s.router.Handler(ctx)
	sub := time.Now().Sub(t)
	klog.V(3).Infof("%s  %s  request from %s  %v  %v", ctx.Method(), ctx.RequestURI(), ctx.Host(), ctx.Response.StatusCode(), sub)
}

func (s *Server) ServerHTTP(address string) error {
	return fasthttp.ListenAndServe(address, s.Handler)
}