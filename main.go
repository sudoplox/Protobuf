package main

import (
	"fmt";,
	"protobuf/src/example.simple"
)
func main() {
	doSimple()
}
func doSimple(){
	sm := example_simple.SimpleMessage{
		Id: 12345,
		IsSimple: true,
		Name: "Sudo",
		SampleList: []int32{1,2,3},
	}
	fmt.Println(sm)
}