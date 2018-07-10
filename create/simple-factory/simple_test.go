package simple_factory

import "testing"

func TestShapeFactory_GetShape(t *testing.T) {

	sp:=ShapeFactory{}
	shape:=sp.GetShape(RECT)
	shape.Draw()
	shape=sp.GetShape(SQUARE)
	shape.Draw()
	shape=sp.GetShape(CIRCLE)
	shape.Draw()
}
