package main

import "fmt"

type product interface {
	doStuff()
}

type productA struct {
	product
}

type productB struct {
	product
}

func (prod productA) doStuff() {
	fmt.Println("AAAAAAAAAAA")
}

func (prod productB) doStuff() {
	fmt.Println("BBBBBBBBBBB")
}

///////////////////////////////////

type creator interface {
	createProduct() product
}

type creatorA struct {
	creator
}

type creatorB struct {
	creator
}

func (cr creatorA) createProduct() product {
	return productA{}
}

func (cr creatorB) createProduct() product {
	return productB{}
}

func main() {

	var aaa creatorA
	var bbb creatorB
	prodA := aaa.createProduct()
	prodA.doStuff()
	prodB := bbb.createProduct()
	prodB.doStuff()
}
