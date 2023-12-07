package gosebus

import (
	"github.com/inquizarus/gosebus/pkg/event"
	"github.com/inquizarus/gosebus/pkg/handler"
)

type Bus interface {
	On(pattern string, handler handler.EventHandler) error
	Handle(handler handler.Handler) error
	Publish(e event.Event) error
}

type standardBus struct {
	handlers []handler.Handler
}

func (b *standardBus) On(pattern string, eh handler.EventHandler) error {
	return b.Handle(handler.New(eh, handler.WithPattern(pattern)))
}

func (b *standardBus) Handle(h handler.Handler) error {
	b.handlers = append(b.handlers, h)
	return nil
}

func (b *standardBus) Publish(e event.Event) error {
	for _, h := range b.handlers {
		if h.Match(e) {
			go func(h handler.Handler, e event.Event) {
				h.Handle(e)
				if h.ShouldRunOnce() {
					b.handlers = b.removeHandler(b.handlers, h)
				}
			}(h, e)
		}
	}
	return nil
}

func (b *standardBus) removeHandler(handlers []handler.Handler, h handler.Handler) []handler.Handler {
	for i, h2 := range handlers {
		if h.Equals(h2) {
			return append(handlers[:i], handlers[i+1:]...)
		}
	}
	return handlers
}

func New() Bus {
	return &standardBus{
		handlers: []handler.Handler{},
	}
}

var DefaultBus = New()
