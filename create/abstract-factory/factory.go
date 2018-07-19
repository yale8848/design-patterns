package abstract_factory

import "fmt"

type Phone interface {
	Play()
}
type Computer interface {
	Office()
}

type XiaoMiPhone struct {

}
func (xmp*XiaoMiPhone)Play()  {

	fmt.Println("XiaoMiPhone phone play good")
}

type SamsungPhone struct {

}
func (ssp*SamsungPhone)Play()  {
	fmt.Println("SamsungPhone phone play bad")
}

type XiaoMiComputer struct {

}

func (xmc*XiaoMiComputer) Office() {
	fmt.Println("XiaoMiComputer Office  bad")
}
type SamSungComputer struct {

}

func (ssc*SamSungComputer) Office() {
	fmt.Println("SamSungComputer Office good")
}
type Factory interface {
	CreatePhone()Phone
	CreateComputer()Computer
}

type EasyFactory struct {

	phone interface{}
	computer interface{}

}

func (xm*EasyFactory) Create(phone interface{},computer interface{}) {

	xm.phone = phone
	xm.computer = computer
}

func (xm*EasyFactory) CreatePhone()Phone {

	eaf:=&EasyFactory{}
	eaf.Create(XiaoMiPhone{},XiaoMiComputer{})
	return &XiaoMiPhone{}
}
func (xm*EasyFactory) CreateComputer()Computer {

	return &XiaoMiComputer{}
}


type XiaoMiFactory struct {

}

func (xm*XiaoMiFactory) CreatePhone()Phone {

	return &XiaoMiPhone{}
}
func (xm*XiaoMiFactory) CreateComputer()Computer {

	return &XiaoMiComputer{}
}

type SamsungFactory struct {

}
func (sf*SamsungFactory) CreatePhone()Phone {

	return &SamsungPhone{}
}
func (sf*SamsungFactory) CreateComputer()Computer {

	return &SamSungComputer{}
}
