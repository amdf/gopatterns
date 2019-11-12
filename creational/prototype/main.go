package main

import (
	"fmt"
	"image/color"
)

type prototype interface {
	clone() prototype
}

type shape struct {
	prototype
	color color.Color
	x, y  int
}

type rectangle struct {
	prototype
	shape
	width, height int
}

type circle struct {
	prototype
	shape
	radius float64
}

func (sh *shape) construct(src shape) {
	sh.color = src.color
	sh.x, sh.y = src.x, src.y
}

func (sh shape) clone() prototype {
	newShape := &shape{}
	newShape.construct(sh)
	return newShape
}

func (rect *rectangle) construct(src rectangle) {
	rect.shape.construct(src.shape)
	rect.width, rect.height = src.width, src.height
}

func (ci *circle) construct(src circle) {
	ci.shape.construct(src.shape)
	ci.radius = src.radius
}

func (rect rectangle) clone() prototype {
	newRect := &rectangle{}
	newRect.construct(rect)
	return newRect
}

func (ci circle) clone() prototype {
	newCircle := &circle{}
	newCircle.construct(ci)
	return newCircle
}

func main() {

	var shapes []prototype
	var copyshapes []prototype

	var newCircle circle
	var newRectangle rectangle
	shapes = append(shapes, newCircle)
	shapes = append(shapes, newRectangle)

	for _, s := range shapes {
		copyshapes = append(copyshapes, s.clone())
	}
	fmt.Println("original: (length ", len(shapes), ")")
	for _, s := range shapes {
		switch x := s.(type) {
		case rectangle:
			fmt.Println("rectangle, height = ", x.height)
		case circle:
			fmt.Println("circle, radius = ", x.radius)
		}
	}
	fmt.Println("copy: (length ", len(copyshapes), ")")
	for _, s := range copyshapes {
		switch x := s.(type) {
		case rectangle:
			fmt.Println("rectangle, height = ", x.height)
		case circle:
			fmt.Println("circle, radius = ", x.radius)
		}
	}
}
