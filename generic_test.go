package conversion

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringConvert(t *testing.T) {
	stringTests := []struct {
		name    string
		args    any
		wantR   any
		wantErr bool
	}{
		{
			name:    "get blank string",
			args:    nil,
			wantR:   "",
			wantErr: false,
		},
		// ----------------------------------------------------------------------------- signed integer
		{
			name:    "convert int to string",
			args:    1,
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "convert intPtr to string",
			args:    Intptr(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "convert int8 to string",
			args:    int8(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "*int8 to string",
			args:    Int8ptr(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "convert int16 to string",
			args:    int16(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "*int16 to string",
			args:    Int16ptr(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "int32 to string",
			args:    int32(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "*int32 to string",
			args:    Int32ptr(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "int64 to string",
			args:    int64(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "*int64 to string",
			args:    Int64ptr(1),
			wantR:   "1",
			wantErr: false,
		},
		// ----------------------------------------------------------------------------- unsigned integer
		{
			name:    "convert uint to string",
			args:    uint(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "convert uintPtr to string",
			args:    Uintptr(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "convert uint8 to string",
			args:    uint8(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "*uint8 to string",
			args:    Uint8ptr(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "convert uint16 to string",
			args:    uint16(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "*uint16 to string",
			args:    Uint16ptr(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "uint32 to string",
			args:    uint32(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "*uint32 to string",
			args:    Uint32ptr(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "uint64 to string",
			args:    uint64(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "*uint64 to string",
			args:    Uint64ptr(1),
			wantR:   "1",
			wantErr: false,
		},

		// ----------------------------------------------------------------------------- float
		{
			name:    "float32 to string",
			args:    float32(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "*float32 to string",
			args:    Float32ptr(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "float64 to string",
			args:    float64(1),
			wantR:   "1",
			wantErr: false,
		},
		{
			name:    "*float64 to string",
			args:    Float64ptr(1),
			wantR:   "1",
			wantErr: false,
		},
		// ----------------------------------------------------------------------------- bool
		{
			name:    "bool to string",
			args:    false,
			wantR:   "false",
			wantErr: false,
		},
		{
			name:    "bool to string",
			args:    true,
			wantR:   "true",
			wantErr: false,
		},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := Convert[string](tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("Convert() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestIntConvert(t *testing.T) {
	intTests := []struct {
		name    string
		args    any
		wantR   int
		wantErr bool
	}{
		{
			name:    "get blank string",
			args:    nil,
			wantR:   0,
			wantErr: false,
		},
		// ----------------------------------------------------------------------------- signed integer
		{
			name:    "convert int to int",
			args:    1,
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "convert intPtr to int",
			args:    Intptr(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "convert int8 to int",
			args:    int8(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "*int8 to string",
			args:    Int8ptr(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "convert int16 to int",
			args:    int16(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "*int16 to int",
			args:    Int16ptr(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "int32 to int",
			args:    int32(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "*int32 to int",
			args:    Int32ptr(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "int64 to int",
			args:    int64(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "*int64 to int",
			args:    Int64ptr(1),
			wantR:   1,
			wantErr: false,
		},
		// ----------------------------------------------------------------------------- unsigned integer
		{
			name:    "convert uint to int",
			args:    uint(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "convert uintPtr to int",
			args:    Uintptr(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "convert uint8 to int",
			args:    uint8(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "*uint8 to int",
			args:    Uint8ptr(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "convert uint16 to int",
			args:    uint16(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "*uint16 to int",
			args:    Uint16ptr(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "uint32 to int",
			args:    uint32(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "*uint32 to int",
			args:    Uint32ptr(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "uint64 to int",
			args:    uint64(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "*uint64 to int",
			args:    Uint64ptr(1),
			wantR:   1,
			wantErr: false,
		},

		// ----------------------------------------------------------------------------- float
		{
			name:    "float32 to int",
			args:    float32(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "*float32 to int",
			args:    Float32ptr(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "float64 to int",
			args:    float64(1),
			wantR:   1,
			wantErr: false,
		},
		{
			name:    "*float64 to int",
			args:    Float64ptr(1),
			wantR:   1,
			wantErr: false,
		},
		// ----------------------------------------------------------------------------- bool
		{
			name:    "bool to int",
			args:    false,
			wantR:   0,
			wantErr: false,
		},
		{
			name:    "bool to int",
			args:    true,
			wantR:   1,
			wantErr: false,
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := Convert[int](tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("Convert() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestArray(t *testing.T) {
	//x := [3]int{1, 2, 3}
	convert, err := Convert[[]int]([]string{"1"})
	assert.Nil(t, err)
	fmt.Println(convert)

	convert, err = Convert[[]int](false)
	assert.Nil(t, err)
	fmt.Println(convert)
}

func TestMap(t *testing.T) {
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
