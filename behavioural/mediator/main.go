package main

import "fmt"

type mediator interface {
	notify(action, param string)
}

type elementA struct {
	m    *myMediator
	data string
}

type elementB struct {
	m    *myMediator
	data string
}

type elementC struct {
	m    *myMediator
	data string
}

type myMediator struct {
	mediator
	a *elementA
	b *elementB
	c *elementC
}

func (m *myMediator) notify(action, param string) {
	switch action {
	case "SETA":
		m.a.data = param
	case "SETB":
		m.b.data = param
	case "SETC":
		m.c.data = param
	case "CLEARALL":
		m.a.data = ""
		m.b.data = ""
		m.c.data = ""
	}
}

func (elem elementA) setA(param string) {
	elem.m.notify("SETA", param)
}
func (elem elementA) setB(param string) {
	elem.m.notify("SETB", param)
}
func (elem elementA) setC(param string) {
	elem.m.notify("SETC", param)
}
func (elem elementA) clear() {
	elem.m.notify("CLEARALL", "")
}

func (elem elementB) setA(param string) {
	elem.m.notify("SETA", param)
}
func (elem elementB) setB(param string) {
	elem.m.notify("SETB", param)
}
func (elem elementB) setC(param string) {
	elem.m.notify("SETC", param)
}
func (elem elementB) clear() {
	elem.m.notify("CLEARALL", "")
}

func (elem elementC) setA(param string) {
	elem.m.notify("SETA", param)
}
func (elem elementC) setB(param string) {
	elem.m.notify("SETB", param)
}
func (elem elementC) setC(param string) {
	elem.m.notify("SETC", param)
}
func (elem elementC) clear() {
	elem.m.notify("CLEARALL", "")
}

func main() {

	var med myMediator
	var a elementA
	a.m = &med
	var b elementB
	b.m = &med
	var c elementC
	c.m = &med
	med.a = &a
	med.b = &b
	med.c = &c

	a.setA("content1 set by a")
	a.setB("content1 set by a")
	a.setC("content1 set by a")
	fmt.Printf("1. abc = %s, %s, %s\n", med.a.data, med.b.data, med.c.data)
	a.clear()
	b.setA("content2 set by b")
	b.setB("content2 set by b")
	b.setC("content2 set by b")
	fmt.Printf("2. abc = %s, %s, %s\n", med.a.data, med.b.data, med.c.data)
	b.clear()
	c.setA("content3 set by c")
	c.setB("content3 set by c")
	c.setC("content3 set by c")
	fmt.Printf("3. abc = %s, %s, %s\n", med.a.data, med.b.data, med.c.data)
	c.clear()
}
