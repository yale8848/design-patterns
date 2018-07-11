package factory_method

import "testing"

func TestFactory(t *testing.T)  {

	sf:=&RectangleFactory{}
	sf.createShape().Draw()

}
