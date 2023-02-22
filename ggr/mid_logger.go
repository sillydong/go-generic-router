package ggr

import (
	"github.com/sillydong/go-generic-router/internal/logx"
	"go.uber.org/zap"
)

func MidLogger[T Context]() Middleware[T] {
	return func(h Handler[T]) Handler[T] {
		return func(ctx T) error {
			reqid := "1234"
			l := logx.Clone().With(zap.String("reqid", reqid))
			ctx.SetLogger(l)
			err := h(ctx)
			if err != nil {
				l.Info("visit error", zap.String("result", err.Error()))
			} else {
				l.Info("visit success")
			}
			return nil
		}
	}
}
