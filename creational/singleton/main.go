package main

import "fmt"

type someClass struct {
	someField int
}

type singleton struct {
	privateData someClass
}

func (s *singleton) getInstance() *singleton {
	if nil == s {
		return &singleton{}
	}
	return s
}

func (s *singleton) printInfo() {
	fmt.Printf("pointer == %p\r\n", s)
}

func main() {
	var b *singleton
	b.printInfo()
	s1 := b.getInstance()
	s1.printInfo()
	s2 := b.getInstance()
	s2.printInfo()
}
