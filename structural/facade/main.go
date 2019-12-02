package main

import "fmt"

type subsystem1 struct {
}

func (s subsystem1) doFeature1() {
	fmt.Println("feature 1")
}

type subsystem2 struct {
}

func (s subsystem2) doFeature2() {
	fmt.Println("feature 2")
}

type subsystem3 struct {
}

func (s subsystem3) doFeature3() {
	fmt.Println("feature 3")
}

type facade struct {
}

var s1 subsystem1
var s2 subsystem2
var s3 subsystem3

func (f facade) doEverything() {
	s1.doFeature1()
	s2.doFeature2()
	s3.doFeature3()
}

func main() {
	var f facade
	f.doEverything()
}
