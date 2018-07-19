// Create by Yale 2018/7/19 18:47
package proxy

import "fmt"

type IA interface {
	a()
}

type A struct {
}

func (a *A) a() {
	fmt.Println("a")
}

type Proxy struct {
	ia IA
}

//和decorator模式比，被代理对象是在proxy对象里生成的
func NewProxy() IA {
	p := &Proxy{}
	p.ia = &A{}
	return p
}

func (p *Proxy) a() {
	fmt.Println("proxy a")
	p.ia.a()
}
