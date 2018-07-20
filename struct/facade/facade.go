// Create by Yale 2018/7/20 17:43
package facade

import "fmt"

type CPU struct {
}

func (c *CPU) start() {

	fmt.Println("cpu start")
}

type Memery struct {
}

func (m *Memery) work() {
	fmt.Println("Memery work")
}

//外观模式一般不涉及接口，它将所要的构建组装起来，对外统一调用
type Computer struct {
	cpu    *CPU
	memery *Memery
}

func (c *Computer) work() {
	fmt.Println("Computer start work")
	c.cpu.start()
	c.memery.work()
}

func NewComputer() *Computer {
	c := &Computer{}
	c.cpu = &CPU{}
	c.memery = &Memery{}
	return c
}
