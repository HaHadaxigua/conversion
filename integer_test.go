package conversion

import (
	"fmt"
	"testing"
)

func TestInt(t *testing.T) {
	x := 111
	v := Int(&x)
	fmt.Println(v)
}
