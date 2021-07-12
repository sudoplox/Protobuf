package main

import (
	"Protobuf/src/simplepb"
	"fmt"
)
func main() {
	doSimple()
}

func doSimple(){
	sm := simplepb.SimpleMessage{
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