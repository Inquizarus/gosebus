package handler

type HandlerOption func(h *standardHandler)

func WithID(id string) HandlerOption {
	return func(h *standardHandler) {
		h.id = id
	}
}

func WithPattern(pattern string) HandlerOption {
	return func(h *standardHandler) {
		h.pattern = pattern
	}
}

func WithWildcardSymbol(symbol string) HandlerOption {
	return func(h *standardHandler) {
		h.wildcardSymbol = symbol
	}
}

func RunOnce() HandlerOption {
	return func(h *standardHandler) {
		h.runOnce = true
	}
}
