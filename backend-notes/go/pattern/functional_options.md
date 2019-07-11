# Introduction

Voir Dave Cheney - functional options

Pour insérer des options dans un nouveau struct, il existe le pattern des functional options.
Celui ci est à utiliser avec les variadic functions.

# Syntaxe

```go
type handler struct {
	s      Story
	t      *template.Template
	pathFc func(r *http.Request) string
}

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFunc(fc func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFc = fc
	}
}

func NewHandler(s Story, options ...HandlerOption) http.Handler {
	h := handler{s, tpl, defaultPathFc}
	for _, option := range options {
		option(&h)
	}
	return h
}

handler := NewHandler(story, WithTemplate(newTemplate))                                  // Avec les functional options, on peut passer un nouveau template avec WithTemplate
```
