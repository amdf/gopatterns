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

	var clang cCompiler
	var golang goCompiler

	fmt.Println("C compiler:")

	clang.preprocess()
	clang.compile()
	clang.link()

	fmt.Println("Go compiler:")

	golang.preprocess()
	golang.compile()
	golang.link()

}
