package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const GENERATED_BY = "// File generated by " + NAME + " v" + VERSION + " (" + URL + ")"

type Generator struct {
	PackageName     string
	VarName         string
	VarData         []byte
	GeneratePackage bool
	GenerateInfo    bool
	GenerateRawData bool
}

func NewGenerator() *Generator {
	g := Generator{}
	g.GeneratePackage = true
	g.GenerateInfo = true
	g.GenerateRawData = false
	return &g
}

func (g *Generator) SetData(data []byte) {
	var totalBytes uint64
	var tmp = ""
	for _, k := range data {
		if totalBytes%12 == 0 {
			tmp += "\n\t"
		}
		tmp += fmt.Sprintf("0x%02x, ", k)
		totalBytes++
	}
	g.VarData = []byte(tmp)
}

func (g *Generator) SetDataFromFile(fPath string) error {
	data, err := ioutil.ReadFile(fPath)
	if err != nil {
		return err
	}
	g.SetData(data)
	g.VarName = FilepathToStructName(fPath)
	return nil
}

// split down to the filename, remove extension and change first char to Uppercase
// http://play.golang.org/p/ic5dxOKdcU
func FilepathToStructName(fPath string) string {
	filename := filepath.Base(fPath)
	fileExt := filepath.Ext(filename)
	filenameWithoutExt := filename[0 : len(filename)-len(fileExt)]
	filenameCleaned := strings.Replace(filenameWithoutExt, " ", "_", -1)
	filenameCleaned = strings.Replace(filenameCleaned, "-", "_", -1)
	filenameCleaned = strings.Replace(filenameCleaned, ".", "_", -1)
	data := []byte(filenameCleaned)
	upper := bytes.ToUpper([]byte{data[0]})
	rest := data[1:]
	return string(bytes.Join([][]byte{upper, rest}, nil))
}

func (g *Generator) GenerateCode() []byte {
	code := ""
	if g.GenerateInfo {
		code += GENERATED_BY + "\n"
	}
	if g.GeneratePackage {
		if g.PackageName == "" {
			g.PackageName = "main"
		}
		code += "package " + g.PackageName + "\n\n"
	}
	code += "var " + g.VarName + " []byte = []byte{" + string(g.VarData) + "\n}\n\n"
	return []byte(code)
}
