package handling

import (
	"strings"

	"github.com/google/uuid"
	"github.com/inquizarus/gosebus/pkg/event"
)

const (
	defaultPattern        = "*"
	defaultWildcardSymbol = "*"
	defaultRunOnce        = false
)

type Handler interface {
	Match(e event.Event) bool
	Handle(e event.Event) error
	ShouldRunOnce() bool
	Equals(h Handler) bool
}

type EventHandler func(event.Event)

type standardHandler struct {
	id             string
	pattern        string
	handle         EventHandler
	wildcardSymbol string
	runOnce        bool
}

func (h *standardHandler) ShouldRunOnce() bool {
	return h.runOnce
}

func (h *standardHandler) Equals(h2 Handler) bool {
	return h.id == h2.(*standardHandler).id
}

func (h *standardHandler) Match(e event.Event) bool {
	wildcards := strings.Count(h.pattern, "*")

	if wildcards == 0 {
		return e.Name() == h.pattern
	}

	if wildcards == 1 {
		if strings.HasPrefix(h.pattern, h.wildcardSymbol) {
			return strings.HasSuffix(e.Name(), string(h.pattern[1:]))
		}
		if strings.HasSuffix(h.pattern, h.wildcardSymbol) {
			return strings.HasPrefix(e.Name(), string(h.pattern[:len(h.pattern)-1]))
		}
		parts := strings.Split(h.pattern, h.wildcardSymbol)
		return strings.HasPrefix(e.Name(), parts[0]) && strings.HasSuffix(e.Name(), parts[1])
	}

	if wildcards > 1 {
		parts := strings.Split(h.pattern, h.wildcardSymbol)
		for _, part := range parts {
			if !strings.Contains(e.Name(), part) {
				return false
			}
		}
		return true
	}

	return false
}

func (h *standardHandler) Handle(e event.Event) error {
	h.handle(e)
	return nil
}

func NewStandardEventHandler(handler EventHandler, options ...HandlerOption) Handler {
	h := standardHandler{
		id:             uuid.New().String(),
		pattern:        defaultPattern,
		handle:         handler,
		wildcardSymbol: defaultWildcardSymbol,
		runOnce:        defaultRunOnce,
	}

	for _, option := range options {
		option(&h)
	}

	return &h
}
