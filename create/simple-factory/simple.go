package simple_factory

import "fmt"

type ShapeType int

const (
	RECT ShapeType = iota
	SQUARE
	CIRCLE
)
type Shape interface {
	Draw()
}
type Rectangle struct {

}
func (rect*Rectangle)Draw()  {
	fmt.Println("Inside Rectangle::draw() method.")
}

type Square struct {

}
func (rect*Square)Draw()  {
	fmt.Println("Inside Square::draw() method.")
}

type Circle struct {

}
func (rect*Circle)Draw()  {
	fmt.Println("Inside Circle::draw() method.")
}

type ShapeFactory struct {
}
func (sf*ShapeFactory) GetShape(shapeType ShapeType)Shape  {

	switch shapeType {
	case RECT:{
		return &Rectangle{}

	}
	case SQUARE:{
		return &Square{}
	}
	case CIRCLE:{
		return &Circle{}
	}
	}

	return nil
}

