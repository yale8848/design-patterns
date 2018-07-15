package adapter

import "fmt"

type V220 interface {
	Connect()
}

type V220Impl struct {

}

func (v*V220Impl)Connect()  {

	fmt.Println("220v")

}


type Adapter struct {

	}
func (v*Adapter)Connect()  {
	fmt.Println("110v")
}

type Cooker struct {

}

func (c*Cooker)Cook(v* Adapter)  {

	v.Connect()
}