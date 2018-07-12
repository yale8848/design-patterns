package builder

import (
	"testing"
	"fmt"
)

func TestNewStudentBuilder(t *testing.T) {
	b:=NewStudentBuilder()
	s:=b.Name("aa").Sex("man").No(1).Build()

	fmt.Printf("(%T)",s)

}
