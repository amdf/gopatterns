package main

import (
	"fmt"
	"sync"
)

type someClass struct {
	someField string
}

type singleton struct {
	privateData someClass
}

var instance *singleton
var once sync.Once

func (s singleton) getInstance() *singleton {
	once.Do(func() {
		instance = &singleton{someClass{"initial value"}}
		return
	})

	return instance
}

func (s *singleton) printInfo() {
	fmt.Printf("pointer == %p, data == %s\r\n", s, s.privateData.someField)
}

var mySingleton singleton

func main() {
	var s1, s2 *singleton

	//make one instance
	s1 = mySingleton.getInstance()
	s1.printInfo()

	//change data
	s1.privateData.someField = "new value"
	s1.printInfo()

	//check another instance
	s2 = mySingleton.getInstance()
	s2.printInfo()
}
