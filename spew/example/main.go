package main

import (
	"fmt"

	"github.com/micronuths/gotools/spew"
)

type HandlerFunc func(in string) string
type Details struct {
	Weight int
	Height int
	Method HandlerFunc
}
type Person struct {
	Name    string
	Age     int
	Details *Details
}

func ReturnFunc() HandlerFunc {
	return func(in string) string {
		out := in
		return out
	}
}
func main() {
	vStrct1 := Details{
		Weight: 55,
		Height: 165,
		Method: ReturnFunc(),
	}
	person := Person{
		Name:    "hesong",
		Age:     27,
		Details: &vStrct1,
	}
	// fmt.Println("person.Details.Method(person.Name)", person.Details.Method(person.Name))
	// spew.Println(vStr)
	// spew.Println(vFloat)
	// spew.Dump(person)
	fmt.Println(spew.Sdump(person))
}
