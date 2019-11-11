package main

import "fmt"

type visitor interface {
	visit(element)
}

type element interface {
	accept(visitor)
}

type myVisitor struct {
	visitor
}

type elementA struct {
	element
	aval int
}
type elementB struct {
	element
	bval int
}
type elementC struct {
	element
	cval int
}

func (m *elementA) accept(v visitor) {
	v.visit(m)
}

func (m *elementB) accept(v visitor) {
	v.visit(m)
}

func (m *elementC) accept(v visitor) {
	v.visit(m)
}

func (v myVisitor) visit(m element) {
	var prev int
	switch x := m.(type) {
	case (*elementA):
		prev = x.aval
		x.aval *= 3
		fmt.Println("visiting A, ", prev, " changed to ", x.aval)
	case (*elementB):
		prev = x.bval
		x.bval *= 5
		fmt.Println("visiting B, ", prev, " changed to ", x.bval)
	case (*elementC):
		prev = x.cval
		x.cval *= 7
		fmt.Println("visiting C, ", prev, " changed to ", x.cval)
	}
}

func main() {

	var elements []element
	var v myVisitor

	elements = append(elements, &elementA{aval: 111})
	elements = append(elements, &elementB{bval: 222})
	elements = append(elements, &elementC{cval: 333})

	for _, x := range elements {
		v.visit(x)
	}
}
