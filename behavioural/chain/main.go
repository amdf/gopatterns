package main

import "fmt"

type stringHandler interface {
	Handle(s string)
	SetNext(n stringHandler)
}

type firstHandler struct {
	stringHandler
	next stringHandler
}

type secondHandler struct {
	stringHandler
	next stringHandler
}

type thirdHandler struct {
	stringHandler
	next stringHandler
}

func (h *firstHandler) SetNext(n stringHandler) {
	if h != nil && n != nil {
		h.next = n
	}
}

func (h *firstHandler) Handle(s string) {
	if h == nil {
		return
	}
	fmt.Println("RUN first handler")
	if s == "first" {
		fmt.Println("OK first handler")
	} else {
		if h.next != nil {
			h.next.Handle(s)
		}
	}
}

func (h *secondHandler) SetNext(n stringHandler) {
	if h != nil && n != nil {
		h.next = n
	}
}

func (h *secondHandler) Handle(s string) {
	if h == nil {
		return
	}
	fmt.Println("RUN second handler")
	if s == "second" {
		fmt.Println("OK second handler")
	} else {
		if h.next != nil {
			h.next.Handle(s)
		}
	}
}

func (h *thirdHandler) SetNext(n stringHandler) {
	if h != nil && n != nil {
		h.next = n
	}
}

func (h *thirdHandler) Handle(s string) {
	if h == nil {
		return
	}
	fmt.Println("RUN third handler")
	if s == "third" {
		fmt.Println("OK third handler")
	} else {
		if h.next != nil {
			h.next.Handle(s)
		}
	}
}

func main() {

	var h1 firstHandler
	var h2 secondHandler
	var h3 thirdHandler

	h1.SetNext(&h2)
	h2.SetNext(&h3)

	h1.Handle("third")
}
