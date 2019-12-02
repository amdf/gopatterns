package main

import "fmt"

type treeLeaf interface {
	execute(string)
}

type simple struct {
	treeLeaf
}
type compound struct {
	treeLeaf
	children []treeLeaf
}

func (leaf simple) execute(padding string) {
	fmt.Println(padding + "i'm just a leaf!")
}

func (leaf compound) execute(padding string) {
	fmt.Println(padding + "compound:")

	for _, x := range leaf.children {
		x.execute(padding + "----")
	}
}

func main() {

	var topDownTree compound

	var a, b compound
	a.children = append(a.children, simple{})
	a.children = append(a.children, simple{})
	a.children = append(a.children, simple{})
	b.children = append(b.children, simple{})
	b.children = append(b.children, simple{})
	b.children = append(b.children, simple{})

	topDownTree.children = append(topDownTree.children, a)
	topDownTree.children = append(topDownTree.children, b)

	topDownTree.execute("")
}
