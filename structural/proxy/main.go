package main

import "fmt"

type mainLib interface {
	doSomething()
}

type mainClass struct {
	mainLib
}

type proxyClass struct {
	mainLib
	service *mainClass
}

func (m *mainClass) doSomething() {
	fmt.Println("main class activity")
}

func (m *proxyClass) doSomething() {

	if m.service == nil {
		fmt.Println("starting service...")
		var myService mainClass
		m.service = &myService
	}
	m.service.doSomething()
}

func main() {

	var svc proxyClass

	for i := 0; i < 10; i++ {
		svc.doSomething()
	}
}
