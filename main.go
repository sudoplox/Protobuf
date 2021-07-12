package main

import (
	example_simple "Protobuf/src/src.simple"
	"fmt"
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
	sm.Name = "Renamed Name"
	fmt.Println(sm)
	fmt.Println("The id is :", sm.GetId())

}