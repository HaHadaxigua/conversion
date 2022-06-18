# Conversion

Just for simplifying type-conversion code.

You can use it convert struct to struct, struct to map, map to struct.

And any other common case of map and slice.

```Go
package main

import "github.com/HaHadaxigua/conversion"

func main() {
	//x := [3]int{1, 2, 3}
	convert, err := Convert[[]int]([]string{"1"})
	assert.Nil(t, err)
	fmt.Println(convert)

	convert, err = Convert[[]int](false)
	assert.Nil(t, err)
	fmt.Println(convert)
	
	type Person struct {
		Name string
	}
	convert, err := Convert[map[string]any](struct {
		A string
	}{
		A: "hello",
	})
	assert.Nil(t, err)
	fmt.Println(convert)

	v2, err := Convert[Person](map[string]any{
		"Name": "hello",
	})
	assert.Nil(t, err)
	fmt.Println(v2)

	v3, err := Convert[Person](struct {
		Name string
	}{
		Name: "xxx",
	})
	assert.Nil(t, err)
	fmt.Println(v3)
}
```