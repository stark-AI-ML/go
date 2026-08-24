package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// well this(code eg) encoding decoding has a problem as we have to know exact json
// struct to create a struct before hand, or you are sure about what you want

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {

	m := Message{"Alice", "Hello", 1294706395881547000}

	// encoding ---------
	// func Marshal(v interface{}) ([]byte, error)

	b, err := json.Marshal(m)

	fmt.Println(b)
	fmt.Println(err)

	// b == []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)

	// decoding ---------
	// func Unmarshal(data []byte, v interface{}) error

	// To decode JSON data we use the Unmarshal function.

	// We must first create a place where the decoded data will be stored

	var k Message

	// and call json.Unmarshal, passing it a []byte of JSON data and a pointer to m

	err = json.Unmarshal(b, &k)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(k.Name) // Alice
	fmt.Println(k.Body) // Hello
	fmt.Println(k.Time) // 1294706395881547000
}
