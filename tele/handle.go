package tele

import (
	"fmt"
	"log"
)

type Handle struct {
	group    *Group
	handlers map[string]HandlerFunc
	OnError  func(error, Context)
}

func NewHandle(b *Handle) *Handle {
	b.group = b.Group()
	b.handlers = make(map[string]HandlerFunc)
	return b
}
func (b *Handle) Use(middleware ...MiddlewareFunc) {
	b.group.Use(middleware...)
}
func (b *Handle) Handle(endpoint interface{}, h HandlerFunc, m ...MiddlewareFunc) {
	if len(b.group.middleware) > 0 {
		m = append(b.group.middleware, m...)
	}
	handler := func(c Context) error {
		return applyMiddleware(h, m...)(c)
	}
	switch end := endpoint.(type) {
	case string:
		b.handlers[end] = handler
	case CallbackEndpoint:
		b.handlers[end.CallbackUnique()] = handler
	default:
		panic("telebot: unsupported endpoint")
	}
}
func (b *Handle) handle(end string, c Context) bool {
	log.Println(end)
	if handler, ok := b.handlers[end]; ok {
		b.runHandler(handler, c)
		return true
	} else if end != OnDefault {
		return b.handle(OnDefault, c)
	}
	return false
}
func (b *Handle) handleText(text string, c Context) bool {
	if handler, ok := b.handlers[text]; ok {
		b.runHandler(handler, c)
		return true
	}
	return false
}
func (b *Handle) runHandler(h HandlerFunc, c Context) {
	f := func() {
		defer b.deferDebug()
		if err := h(c); err != nil {
			b.OnError(err, c)
		}
	}
	go f()
}
func (b *Handle) deferDebug() {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			log.Println(err)
		} else if str, ok := r.(string); ok {
			log.Println(fmt.Errorf("%s", str))
		}
	}
}
func (b *Handle) Group() *Group {
	return &Group{b: b, middleware: make([]MiddlewareFunc, 0)}
}
