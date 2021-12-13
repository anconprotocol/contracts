package main

import "github.com/wasmerio/wasmer-go/wasmer"

//	"syscall/js"

//export add
func Add() wasmer.Value {


        return wasmer.NewValue(
			"a", wasmer.I64,
		)
}

func main() {
	//	cli := graphqlclient.NewClient(http.DefaultClient, "http://localhost:7788/v0/query")
}
