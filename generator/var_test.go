package generator

import (
	"testing"
)

func Test_NewVar(t *testing.T) {
	v := NewVar()
	if len(v.Data) != 0 {
		t.Error("Data length not equal")
	}
}
