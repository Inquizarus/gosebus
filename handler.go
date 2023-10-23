package gosebus

import "strings"

type Handler interface {
	Match(e Event) bool
	Handle(e Event) error
}

type EventHandler func(Event)

type standardHandler struct {
	pattern        string
	handle         EventHandler
	wildcardSymbol string
	runOnce        bool
}

func (h *standardHandler) Match(e Event) bool {

	wildcards := strings.Count(h.pattern, "*")

	if wildcards == 0 {
		// If there are no wildcards, return straight comparison to the pattern
		return e.Name() == h.pattern
	}

	if wildcards == 1 {

		if strings.HasPrefix(h.pattern, h.wildcardSymbol) {
			// If there is a wildcard at the beginning of the pattern, check for the remainder as a suffix
			return strings.HasSuffix(e.Name(), string(h.pattern[1:]))
		}
		if strings.HasSuffix(h.pattern, h.wildcardSymbol) {
			// If there is a wildcard at the end of the pattern, check for the remainder as a prefix
			return strings.HasPrefix(e.Name(), string(h.pattern[:len(h.pattern)-1]))
		}
		// If there is a wildcard in the middle of the pattern, split and check for prefix and suffix
		parts := strings.Split(h.pattern, h.wildcardSymbol)

		return strings.HasPrefix(e.Name(), parts[0]) && strings.HasSuffix(e.Name(), parts[1])

	}

	if wildcards > 1 {
		// If there are multiple wildcards, split by wildcard symbol and check each
		parts := strings.Split(h.pattern, h.wildcardSymbol)

		for i := 0; i < len(parts); i++ {
			part := parts[i]
			if !strings.Contains(e.Name(), part) {
				return false
			}
		}

		return true
	}

	return false
}

func (h *standardHandler) Handle(e Event) error {
	h.handle(e)
	return nil
}

func NewStandardEventHandler(pattern string, handler EventHandler) Handler {
	return &standardHandler{
		pattern:        pattern,
		handle:         handler,
		wildcardSymbol: "*",
		runOnce:        false,
	}
}
