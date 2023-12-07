package gosebus

import (
	"github.com/inquizarus/gosebus/pkg/event"
	"github.com/inquizarus/gosebus/pkg/handling"
)

type Bus interface {
	On(pattern string, handler handling.EventHandler) error
	Handle(handler handling.Handler) error
	Publish(e event.Event) error
}

type standardBus struct {
	handlers []handling.Handler
}

func (b *standardBus) On(pattern string, eh handling.EventHandler) error {
	return b.Handle(handling.NewStandardEventHandler(eh, handling.HandlerOptionWithPattern(pattern)))
}

func (b *standardBus) Handle(h handling.Handler) error {
	b.handlers = append(b.handlers, h)
	return nil
}

func (b *standardBus) Publish(e event.Event) error {
	for _, h := range b.handlers {
		if h.Match(e) {
			go func(h handling.Handler, e event.Event) {
				h.Handle(e)
				if h.ShouldRunOnce() {
					b.handlers = b.removeHandler(b.handlers, h)
				}
			}(h, e)
		}
	}
	return nil
}

func (b *standardBus) removeHandler(handlers []handling.Handler, h handling.Handler) []handling.Handler {
	for i, h2 := range handlers {
		if h.Equals(h2) {
			return append(handlers[:i], handlers[i+1:]...)
		}
	}
	return handlers
}

func New() Bus {
	return &standardBus{
		handlers: []handling.Handler{},
	}
}

var DefaultBus = New()
