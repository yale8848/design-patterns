package factory

import "fmt"

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

type ShapeFactory interface {
	createShape() Shape
}

type RectangleFactory struct {

}
func (rf*RectangleFactory)createShape() Shape  {
	return &Rectangle{}
}


