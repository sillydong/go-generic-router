package logx_test

import (
	"testing"

	"github.com/sillydong/go-generic-router/internal/logx"
	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	logx.New("xxx", logx.InfoLevel)

	logx.Info("message", zap.String("key", "value"))

	x := New()
	x.Test()
	logx.Info("message2", zap.String("key", "value"))
}

type X struct {
	Logger *zap.Logger
}

func New() *X {
	return &X{
		Logger: logx.Clone().With(zap.String("xid", "aaaa")),
	}
}

func (x *X) Test() {
	x.Logger.Info("in x")
}
