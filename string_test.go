package conversion

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	v := "11"
	x := String(&v)
	fmt.Println(x)
}
