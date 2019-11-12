package main

import "fmt"

type abstractStringBuilder interface {
	setFirstLetter(rune)
	setSecondLetter(rune)
	setThirdLetter(rune)
	setFourthLetter(rune)
	getResult() string
}

type fourLetterBuilder struct {
	abstractStringBuilder
	a, b, c, d  rune
	finalString string
}

func (f *fourLetterBuilder) setFirstLetter(r rune) {
	f.a = r
}
func (f *fourLetterBuilder) setSecondLetter(r rune) {
	f.b = r
}
func (f *fourLetterBuilder) setThirdLetter(r rune) {
	f.c = r
}
func (f *fourLetterBuilder) setFourthLetter(r rune) {
	f.d = r
}
func (f fourLetterBuilder) getResult() string {
	return fmt.Sprintf("%c%c%c%c", f.a, f.b, f.c, f.d)
}

type zxcvDirector struct {
	bu *fourLetterBuilder
}

type asdfDirector struct {
	bu *fourLetterBuilder
}

func (dir zxcvDirector) getResult() string {
	dir.bu.setFirstLetter('z')
	dir.bu.setSecondLetter('x')
	dir.bu.setThirdLetter('c')
	dir.bu.setFourthLetter('v')
	return dir.bu.getResult()
}

func (dir asdfDirector) getResult() string {
	dir.bu.setFirstLetter('a')
	dir.bu.setSecondLetter('s')
	dir.bu.setThirdLetter('d')
	dir.bu.setFourthLetter('f')
	return dir.bu.getResult()
}

func main() {
	dir1 := zxcvDirector{bu: &fourLetterBuilder{}}
	dir2 := asdfDirector{bu: &fourLetterBuilder{}}
	fmt.Println(dir1.getResult())
	fmt.Println(dir2.getResult())
}
