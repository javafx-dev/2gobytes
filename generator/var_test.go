package generator

import (
	"strings"
	"testing"
)

func Test_NewVar(t *testing.T) {
	v := NewVar()
	if len(v.Data) != 0 {
		t.Error("Data length not equal")
	}
}

func Test_Var_SetData(t *testing.T) {
	v := NewVar()
	v.SetData("color", []byte("hello world"))
	code := v.GenerateCode()
	result := []string{
		"// Color store the data as a byte array.",
		"var Color []byte = []byte{",
		"	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, ",
		"}",
		"",
	}
	if code != strings.Join(result, "\n") {
		t.Error("SetData Generated Code not equal")
	}
}

func Test_Var_SetDataFromFile(t *testing.T) {
	v := NewVar()
	v.SetDataFromFile("../fixture/sample.txt")
	code := v.GenerateCode()
	result := []string{
		"// Sample store the data of '../fixture/sample.txt' as a byte array.",
		"var Sample []byte = []byte{",
		"	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x0a, ",
		"}",
		"",
	}
	if code != strings.Join(result, "\n") {
		t.Error("SetDataFromFile Generated Code not equal")
	}
}

func Test_Var_SetDataFromFile_With_Name(t *testing.T) {
	v := NewVar()
	v.SetDataFromFile("../fixture/sample.txt")
	v.Name = "Color"
	code := v.GenerateCode()
	result := []string{
		"// Color store the data of '../fixture/sample.txt' as a byte array.",
		"var Color []byte = []byte{",
		"	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x0a, ",
		"}",
		"",
	}
	if code != strings.Join(result, "\n") {
		t.Error("SetDataFromFile with Name Generated Code not equal")
	}
}
