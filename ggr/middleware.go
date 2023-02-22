package ggr

type Middleware[T Context] func(Handler[T]) Handler[T]

func wrapMiddleware[T Context](mw []Middleware[T], handler Handler[T]) Handler[T] {
	for i := len(mw) - 1; i >= 0; i-- {
		h := mw[i]
		if h != nil {
			handler = h(handler)
		}
	}

	return handler
}
