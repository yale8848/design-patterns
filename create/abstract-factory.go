package create

type AbstractFactory interface {
	CreateProductA()
	CreateProductB()
}

type ConcreteFactory1 struct {

}
type ConcreteFactory2 struct {

}
