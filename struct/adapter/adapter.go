package adapter

import "fmt"

type Usb interface {
	ConnectUsb()
}
type PS2 interface {
	ConnectPS2()
}
type Usber struct {
}

func (u *Usber) ConnectUsb() {
	fmt.Println("ConnectUsb")
}

//类的适配，继承已有的类，实现要用的接口
type AdapterClass struct {
	Usber
}

func (a *AdapterClass) ConnectPS2() {
	fmt.Println("ConnectPS2")
	a.Usber.ConnectUsb()
}

type AdapterObj struct {
	Usber Usb
}

//对象适配，set已有的对象，实现要用的接口
func (a *AdapterObj) ConnectPS2() {
	fmt.Println("ConnectPS2")
	a.Usber.ConnectUsb()
}

type A interface {
	a()
	b()
	c()
}

//接口适配，先实现已有的接口，继承并并覆盖需要的方法
type AdapterInterface struct {
}

func (A *AdapterInterface) a() {
	fmt.Println("a")
}
func (A *AdapterInterface) b() {

}
func (A *AdapterInterface) c() {

}

type B struct {
	AdapterInterface
}

func (b *B) a() {
	fmt.Println("b>a")

}
