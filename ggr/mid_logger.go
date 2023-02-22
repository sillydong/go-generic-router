package ggr

import (
	"github.com/sillydong/go-generic-router/internal/logx"
	"go.uber.org/zap"
)

func MidLogger[T Context]() Middleware[T] {
	return func(h Handler[T]) Handler[T] {
		return func(ctx T) any {
			reqid := "1234"
			l := logx.Clone().With(zap.String("reqid", reqid))
			ctx.SetLogger(l)
			res := h(ctx)
			l.Info("accesslog", zap.Any("response", res))
			return res
		}
	}
}
