package ggr

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/naoina/denco"
)

type Handler[T Context] func(ctx T) error

type Router[T Context] struct {
	mux      *denco.Mux
	handlers []denco.Handler
	mw       []Middleware[T]
	new      func() T
}

func NewRouter[T Context](new func() T) *Router[T] {
	return &Router[T]{
		mux:      denco.NewMux(),
		handlers: []denco.Handler{},
		mw: []Middleware[T]{
			MidLogger[T](),
			MidRecover[T](),
		},
		new: new,
	}
}

func (ro *Router[T]) Handler() (http.Handler, error) {
	return ro.mux.Build(ro.handlers)
}

func (ro *Router[T]) Get(path string, handler Handler[T], mw ...Middleware[T]) {
	ro.Handle(http.MethodGet, path, handler, mw...)
}

func (ro *Router[T]) Post(path string, handler Handler[T], mw ...Middleware[T]) {
	ro.Handle(http.MethodPost, path, handler, mw...)
}

func (ro *Router[T]) Put(path string, handler Handler[T], mw ...Middleware[T]) {
	ro.Handle(http.MethodPut, path, handler, mw...)
}

func (ro *Router[T]) Delete(path string, handler Handler[T], mw ...Middleware[T]) {
	ro.Handle(http.MethodDelete, path, handler, mw...)
}

func (ro *Router[T]) Options(path string, handler Handler[T], mw ...Middleware[T]) {
	ro.Handle(http.MethodOptions, path, handler, mw...)
}

func (ro *Router[T]) Head(path string, handler Handler[T], mw ...Middleware[T]) {
	ro.Handle(http.MethodHead, path, handler, mw...)
}

func (ro *Router[T]) Handle(method, path string, handler Handler[T], mw ...Middleware[T]) {
	handler = wrapMiddleware(mw, handler)
	handler = wrapMiddleware(ro.mw, handler)

	h := ro.mux.Handler(method, path, ro.wrap(handler))
	ro.handlers = append(ro.handlers, h)
}

func (ro *Router[T]) wrap(handler Handler[T]) denco.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params denco.Params) {
		cx := ro.new()
		cx.SetContext(r.Context())
		cx.SetRequest(r)
		cx.SetResponseWriter(w)
		cx.SetParams(params)
		res := handler(cx)
		switch res.(type) {
		case error:
			fmt.Fprintf(w, res.Error())
		default:
			json.NewEncoder(w).Encode(res)
		}
	}
}
