// Create by Yale 2018/7/20 18:01
package bridge

import "testing"

func TestBridge(t *testing.T) {

	b := NewJDBCBridge()
	b.SetJDBC(&MySQL{})
	b.Connect()

	b.SetJDBC(&Oracle{})
	b.Connect()
}
