package main

import (
	"log"
	"net/http"

	"github.com/sillydong/go-generic-router/ggr"
	"github.com/sillydong/go-generic-router/internal/logx"
	"go.uber.org/zap"
)

type XContext struct {
	ggr.ReqContext

	User string
}

func Hello(ctx *XContext) any {
	ctx.Logger.Info("in function now", zap.String("user", ctx.User))
	logx.Info("asdfasdf") // to verify if caller in log is right
	return "hello: " + ctx.User
}

func SetUser() ggr.Middleware[*XContext] {
	return func(h ggr.Handler[*XContext]) ggr.Handler[*XContext] {
		return func(ctx *XContext) any {
			ctx.User = "hello"
			ctx.Logger.Info("in middleware") // to verify if caller in log is right
			return h(ctx)
		}
	}
}

func SetUserGroup() ggr.Middleware[*XContext] {
	return func(h ggr.Handler[*XContext]) ggr.Handler[*XContext] {
		return func(ctx *XContext) any {
			ctx.User = "hello group"
			ctx.Logger.Info("in group middleware") // to verify if caller in log is right
			return h(ctx)
		}
	}
}

func main() {
	r := ggr.NewRouter(func() *XContext {
		return &XContext{}
	})
	r.Get("/hello", Hello, SetUser())
	g := r.Group("/test", SetUserGroup())
	g.Get("/hellotest", Hello)

	h, err := r.Handler()
	if err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe(":8080", h))
}
