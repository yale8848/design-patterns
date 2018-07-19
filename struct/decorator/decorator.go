// Create by Yale 2018/7/19 18:21
package decorator

import "fmt"

type IA interface {
	a()
}
type A struct {
}

func (a *A) a() {
	fmt.Println("a")
}

//和proxy对比，装饰模式的被装饰对象是set进来的
type Decorator struct {
	ao IA
}

func (d *Decorator) a() {
	fmt.Println("Decorator first")
	d.ao.a()
}
