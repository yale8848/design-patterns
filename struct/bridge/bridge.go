// Create by Yale 2018/7/20 17:55
package bridge

import "fmt"

type Drawer struct {
}
type Shape interface {
	draw()
}

type Rect struct {
}

func (r *Rect) draw() {

}

type Color interface {
	paint(shape string)
}

type Red struct {
}

func (r *Red) paint(shape string) {

	fmt.Println("red " + shape)
}
