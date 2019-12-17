package httpserver

import (
	"fmt"
	"time"

	"k8s.io/klog"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type Server struct {
	router      *router.Router
	Interceptor *func(ctx *fasthttp.RequestCtx) error
	prefix      string
}

func New(router *router.Router, prefix string) *Server {
	return &Server{router, nil, prefix}
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

/* func (s *Server) Handler(ctx *fasthttp.RequestCtx) {
	var sub time.Duration

	defer func() {
		klog.V(2).Infof("%s  %v  %s", ctx.String(), ctx.Response.StatusCode(), sub)
	}()

	t := time.Now()
	if !(string(ctx.RequestURI()) == "/api/v1/users/login") {
		if s.Interceptor != nil {
			(*s.Interceptor)(ctx)
			sub = time.Now().Sub(t)
			if ctx.Response.StatusCode() == 403 {
				return
			}
		}

		cluster := string(ctx.Request.Header.Peek("cluster"))
		c, err := serverapi.GetCluster(cluster)
		if err != nil {
			ctx.Error(fmt.Sprintf("query cluster error %v", err.Error()), 400)
			return
		}
		if !c.IsMaster {
			ctx.Request.SetHost(c.Endpoint)
			fasthttp.Do(&ctx.Request, &ctx.Response)
			sub = time.Now().Sub(t)
		} else {
			s.router.Handler(ctx)
			sub = time.Now().Sub(t)
		}
	} else {
		s.router.Handler(ctx)
		sub = time.Now().Sub(t)
	}
} */

func (s *Server) Handler(ctx *fasthttp.RequestCtx) {
	var sub time.Duration

	defer func() {
		klog.V(2).Infof("%s  %v  %s", ctx.String(), ctx.Response.StatusCode(), sub)
	}()

	t := time.Now()
	if s.Interceptor != nil {
		if (*s.Interceptor)(ctx) != nil {
			sub = time.Now().Sub(t)
			return
		}
	}

	s.router.Handler(ctx)
	sub = time.Now().Sub(t)
}

func (s *Server) ServerHTTP(address string) error {
	return fasthttp.ListenAndServe(address, s.Handler)
}

func (s *Server) AddInterceptor(interceptor func(ctx *fasthttp.RequestCtx) error) {
	s.Interceptor = &interceptor
}
