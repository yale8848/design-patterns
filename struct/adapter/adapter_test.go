package adapter

import "testing"

func TestAdapterClass(t *testing.T) {

	var ps2 PS2
	ps2 = &AdapterClass{}
	ps2.ConnectPS2()

}
func TestAdapterObj(t *testing.T) {

	var ps2 AdapterObj
	ps2 = AdapterObj{}
	ps2.Usber = &Usber{}
	ps2.ConnectPS2()

}

func TestAdapterInterface(t *testing.T) {

	var a A
	a = &B{}
	a.a()

}
