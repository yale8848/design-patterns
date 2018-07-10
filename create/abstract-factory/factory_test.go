package abstract_factory

import "testing"

func TestFactory(t *testing.T) {

	xmf:=&XiaoMiFactory{}
	xmf.CreatePhone().Play()
	xmf.CreateComputer().Office()

	ssf:=&SamsungFactory{}
	ssf.CreatePhone().Play()
	ssf.CreateComputer().Office()
}

