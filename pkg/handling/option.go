package handling

type HandlerOption func(h *standardHandler)

func HandlerOptionWithID(id string) HandlerOption {
	return func(h *standardHandler) {
		h.id = id
	}
}

func HandlerOptionWithPattern(pattern string) HandlerOption {
	return func(h *standardHandler) {
		h.pattern = pattern
	}
}

func HandlerOptionWithWildcardSymbol(symbol string) HandlerOption {
	return func(h *standardHandler) {
		h.wildcardSymbol = symbol
	}
}

func HandlerOptionRunOnce() HandlerOption {
	return func(h *standardHandler) {
		h.runOnce = true
	}
}
