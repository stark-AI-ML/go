package main

import (
	"encoding/json"
	"fmt"
)

// Decoding arbitary data from byte (Buffer) to json

func main() {

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	fmt.Println(b)

	var f interface{}
	err := json.Unmarshal(b, &f) // this will map through interface and store the value in it
	defer fmt.Print(err)

	// At this point the Go value in f would be a map whose keys are strings and whose
	// values are themselves stored as empty interface values:

	// f = map[string]interface{}{
	// 	"Name": "Wednesday",
	// 	"Age":  6,
	// 	"Parents": []interface{}{
	// 		"Gomez",
	// 		"Morticia",
	// 	},
	// }

	fmt.Println("f : ", f)

	m := f.(map[string]interface{})

	fmt.Println("m : ", m)

	// 	f :  map[Age:6 Name:Wednesday Parents:[Gomez Morticia]]
	//  m :  map[Age:6 Name:Wednesday Parents:[Gomez Morticia]]

	// so fmt.unmarshal did it's work and stored each json value to f itterating over
	// and stored value in form of map
}
