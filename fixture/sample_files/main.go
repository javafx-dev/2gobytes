//go:generate ../../2gobytes --index Files --input ../sample.txt --input ../../README.md --output data.go

package main

import "fmt"

func main() {
	fmt.Println(string(*Files["../sample.txt"]))
	fmt.Println(Files)
}
