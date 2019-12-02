package main

import (
	"fmt"
	"image/color"
)

type forest struct {
	trees []tree
}

type tree struct {
	t *treeType
}

type treeType struct {
	color   color.Color
	texture [1024]byte
}

type treeFactory struct {
	treeTypes map[color.Color]*treeType
}

func (tf *treeFactory) getTreeType(cl color.Color) *treeType {

	_, ok := tf.treeTypes[cl]

	if !ok {
		fmt.Println("allocate")
		var newTree treeType
		newTree.color = cl
		tf.treeTypes[cl] = &newTree
	} else {
		fmt.Println("reuse")
	}
	return tf.treeTypes[cl]
}

func (f *forest) plantTrees(count int, cl color.Color) {
	for i := 0; i < count; i++ {
		f.trees = append(f.trees, tree{t: tf.getTreeType(cl)})
	}
}

var tf treeFactory

func main() {

	tf.treeTypes = make(map[color.Color]*treeType)

	var f forest
	f.plantTrees(10, color.RGBA{G: 255})
	f.plantTrees(10, color.RGBA{G: 255, R: 255})
	f.plantTrees(2, color.RGBA{})
	fmt.Println("Tree colors: ", len(tf.treeTypes))
	fmt.Println("Number of trees: ", len(f.trees))
}
