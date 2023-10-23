package gosebus

type Bus interface {
	On(pattern string, handler EventHandler) error
	Handle(handler Handler) error
	Publish(e Event) error
}

type standardBus struct {
	handlers []Handler
}

func (b *standardBus) On(pattern string, handle EventHandler) error {
	return b.Handle(&standardHandler{
		pattern:        pattern,
		handle:         handle,
		wildcardSymbol: "*",
	})
}

func (b *standardBus) Handle(h Handler) error {
	b.handlers = append(b.handlers, h)
	return nil
}

func (b *standardBus) Publish(e Event) error {
	for _, h := range b.handlers {
		if h.Match(e) {
			go func(e Event) {
				h.Handle(e)
				// TODO: add something that removes this handler if runOnce is true
			}(e)
		}
	}
	return nil
}

func New() Bus {
	return &standardBus{
		handlers: []Handler{},
	}
}

var DefaultBus = New()
