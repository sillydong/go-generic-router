package ggr

import (
	"net/http"
	"net/url"
)

type Group[T Context] struct {
	ro   *Router[T]
	name string
	mw   []Middleware[T]
}

func (g *Group[T]) Get(path string, handler Handler[T], mw ...Middleware[T]) {
	g.Handle(http.MethodGet, path, handler, mw...)
}

func (g *Group[T]) Post(path string, handler Handler[T], mw ...Middleware[T]) {
	g.Handle(http.MethodPost, path, handler, mw...)
}

func (g *Group[T]) Put(path string, handler Handler[T], mw ...Middleware[T]) {
	g.Handle(http.MethodPut, path, handler, mw...)
}

func (g *Group[T]) Delete(path string, handler Handler[T], mw ...Middleware[T]) {
	g.Handle(http.MethodDelete, path, handler, mw...)
}

func (g *Group[T]) Options(path string, handler Handler[T], mw ...Middleware[T]) {
	g.Handle(http.MethodOptions, path, handler, mw...)
}

func (g *Group[T]) Head(path string, handler Handler[T], mw ...Middleware[T]) {
	g.Handle(http.MethodHead, path, handler, mw...)
}

func (g *Group[T]) Handle(method, path string, handler Handler[T], mw ...Middleware[T]) {
	handler = wrapMiddleware(mw, handler)
	handler = wrapMiddleware(g.mw, handler)
	handler = wrapMiddleware(g.ro.mw, handler)

	path, _ = url.JoinPath(g.name, path)
	h := g.ro.mux.Handler(method, path, g.ro.wrap(handler))
	g.ro.handlers = append(g.ro.handlers, h)
}
