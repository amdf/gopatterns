package main

import "fmt"

type someClient struct {
	adapterRef *adapter
}

type adapter struct {
	classRef *someClass
}

type someClass struct {
}

func (sc someClass) method(onlyString string) {
	fmt.Printf("I prefer only strings; the string is %s\r\n", onlyString)
}

func (sc someClient) method(someInteger int) {
	sc.adapterRef.method(someInteger)
}

func (sc adapter) method(someInteger int) {
	converted := fmt.Sprintf("%d", someInteger)
	sc.classRef.method(converted)
}

func main() {

	var c someClass
	adap := adapter{classRef: &c}
	cli := someClient{adapterRef: &adap}
	cli.method(123)
}
