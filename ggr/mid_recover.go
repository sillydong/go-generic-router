package ggr

import "go.uber.org/zap"

func MidRecover[T Context]() Middleware[T] {
	return func(h Handler[T]) Handler[T] {
		return func(ctx T) error {
			defer func() {
				if err := recover(); err != nil {
					ctx.GetLogger().WithOptions(zap.AddCallerSkip(2)).Panic("HELP", zap.Any("error", err))
				}
			}()
			return h(ctx)
		}
	}
}
