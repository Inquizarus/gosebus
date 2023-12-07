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
				// TODO: add something that removes this handler if runOnce is true
			}(h, e)
		}
	}
	return nil
}

func New() Bus {
	return &standardBus{
		handlers: []handling.Handler{},
	}
}

var DefaultBus = New()
