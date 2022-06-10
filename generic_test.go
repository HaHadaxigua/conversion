package conversion

import (
	"reflect"
	"testing"
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
