package main

import "fmt"

type implementation interface {
	method1()
	method2()
}

type abstraction struct {
	imp implementation
}

func (ab abstraction) feature1() {
	ab.imp.method1()
}
func (ab abstraction) feature2() {
	ab.imp.method2()
}

type myImplementation1 struct {
	implementation
}

func (imp myImplementation1) method1() {
	fmt.Println("Feature 1, Type I implementation")
}
func (imp myImplementation1) method2() {
	fmt.Println("Feature 2, Type I implementation")
}

type myImplementation2 struct {
	implementation
}

func (imp myImplementation2) method1() {
	fmt.Println("Feature 1, Type II implementation")
}

func (imp myImplementation2) method2() {
	fmt.Println("Feature 2, Type II implementation")
}

func main() {

	var imp1 myImplementation1
	var imp2 myImplementation2
	abs1 := abstraction{imp: &imp1}
	abs2 := abstraction{imp: &imp2}

	abs1.feature1()
	abs1.feature2()
	abs2.feature1()
	abs2.feature2()
}
