package main

import "fmt"

type decorator interface {
	execute()
}

type middleClass struct {
	decorator
	super baseClass
}

type finalClass struct {
	decorator
	super middleClass
}

type baseClass struct {
	decorator
}

func (b baseClass) execute() {
	fmt.Println("base function execute")
}

func (b middleClass) execute() {
	b.super.execute()
	fmt.Println("middle function execute")
}

func (b finalClass) execute() {
	b.super.execute()
	fmt.Println("final function execute")
}

func main() {
	var b finalClass
	var ptr decorator
	ptr = &b
	ptr.execute()
}
