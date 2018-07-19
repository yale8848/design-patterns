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

//类的装饰是要
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

func (a *AdapterObj) ConnectPS2() {
	fmt.Println("ConnectPS2")
	a.Usber.ConnectUsb()
}

type A interface {
	a()
	b()
	c()
}

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
