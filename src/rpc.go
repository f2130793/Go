package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	fmt.Println(2)
	client, err := rpc.Dial("", "localhost:8500")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call("carno.tests.hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
