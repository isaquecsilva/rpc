package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	defer client.Close()

	arg := 0
	var files []string	
	err = client.Call("Worker.ListDirFiles", &arg, &files)
	if err != nil {
		client.Close()
		panic(err)
	}

	fmt.Println("RPC-REPLY:")
	if len(files) == 0 {
		fmt.Println("empty folder")
	}

	for _, filename := range files {
		fmt.Println(filename)
	}	
}