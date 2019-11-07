package main

import (
	"errors"
	"fmt"
)

type myCollection struct {
	numbers []int
}

func (col myCollection) getOddIterator() iterator {
	return &oddIterator{col: &col, pos: 0}
}

func (col myCollection) getHundredIterator() iterator {
	return &hundredIterator{col: &col, pos: 0}
}

type iterator interface {
	getNext() (int, error)
}

type oddIterator struct {
	iterator
	col *myCollection
	pos int
}

func (it *oddIterator) getNext() (result int, err error) {
	if it.col == nil {
		err = errors.New("collection nil")
		return
	}
	if it.pos >= len(it.col.numbers) {
		err = errors.New("no more")
		return
	}

	for it.pos < len(it.col.numbers) {
		val := it.col.numbers[it.pos]
		it.pos = it.pos + 1
		if val%2 == 0 {
			result = val
			return
		}
	}

	err = errors.New("no more")
	return
}

////
type hundredIterator struct {
	iterator
	col *myCollection
	pos int
}

func (it *hundredIterator) getNext() (result int, err error) {
	if it.col == nil {
		err = errors.New("collection nil")
		return
	}
	if it.pos >= len(it.col.numbers) {
		err = errors.New("no more")
		return
	}

	for it.pos < len(it.col.numbers) {
		val := it.col.numbers[it.pos]
		it.pos = it.pos + 1
		if val%100 == 0 && val >= 100 {
			result = val
			return
		}
	}

	err = errors.New("no more")
	return
}

func main() {
	var col myCollection
	for i := 0; i < 1000; i++ {
		col.numbers = append(col.numbers, i)
	}

	it := col.getHundredIterator()

	for {
		x, err := it.getNext()
		if err == nil {
			fmt.Printf("%d,", x)
		} else {
			break
		}
	}
}
