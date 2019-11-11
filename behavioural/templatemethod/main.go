package main

import "fmt"

type compiler interface {
	preprocess()
	compile()
	link()
}

type goCompiler struct {
	compiler
}

type cCompiler struct {
	compiler
}

func (c cCompiler) preprocess() {
	fmt.Println("preprocessing...")
}

func (c goCompiler) preprocess() {
	fmt.Println("skip preprocessing...")
}

func (c goCompiler) compile() {
	fmt.Println("compiling Go source...")
}

func (c cCompiler) compile() {
	fmt.Println("compiling C source...")
}

func (c goCompiler) link() {
	fmt.Println("linking .a files...")
}

func (c cCompiler) link() {
	fmt.Println("linking .obj files...")
}

func main() {

	var compc cCompiler
	var compgo goCompiler

	fmt.Println("C compiler:")

	compc.preprocess()
	compc.compile()
	compc.link()

	fmt.Println("Go compiler:")

	compgo.preprocess()
	compgo.compile()
	compgo.link()

}
