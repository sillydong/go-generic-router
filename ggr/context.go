package ggr

import (
	"context"
	"net/http"

	"github.com/naoina/denco"
	"go.uber.org/zap"
)

type Context interface {
	context.Context

	SetContext(context.Context)
	SetRequest(*http.Request)
	SetResponseWriter(http.ResponseWriter)
	SetParams(params denco.Params)
	SetLogger(*zap.Logger)

	GetRequest() *http.Request
	GetResponseWriter() http.ResponseWriter
	GetParams() denco.Params
	GetLogger() *zap.Logger
}

type ReqContext struct {
	context.Context
	Logger   *zap.Logger
	Request  *http.Request
	Response http.ResponseWriter
	Params   denco.Params
}

func (r *ReqContext) SetLogger(l *zap.Logger) {
	r.Logger = l
}

func (r *ReqContext) SetContext(ctx context.Context) {
	r.Context = ctx
}

func (r *ReqContext) SetRequest(req *http.Request) {
	r.Request = req
}

func (r *ReqContext) SetResponseWriter(rw http.ResponseWriter) {
	r.Response = rw
}

func (r *ReqContext) SetParams(params denco.Params) {
	r.Params = params
}

func (r *ReqContext) GetRequest() *http.Request {
	return r.Request
}

func (r *ReqContext) GetResponseWriter() http.ResponseWriter {
	return r.Response
}

func (r *ReqContext) GetParams() denco.Params {
	return r.Params
}

func (r *ReqContext) GetLogger() *zap.Logger {
	return r.Logger
}
