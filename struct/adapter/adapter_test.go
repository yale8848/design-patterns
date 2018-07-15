package adapter

import "testing"

func TestCooke(t *testing.T) {

	c:=&Cooker{}
	adap:=&Adapter{}
	c.Cook(adap)

}
