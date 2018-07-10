package factory

import "testing"

func TestFactory(t *testing.T)  {

	sf:=&RectangleFactory{}
	sf.createShape().Draw()

}
