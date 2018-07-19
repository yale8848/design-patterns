// Create by Yale 2018/7/19 18:23
package decorator

import "testing"

func TestDecorator(t *testing.T) {

	de := &Decorator{}
	de.ao = &A{}
	de.a()
}
