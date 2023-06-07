package main

import "fmt"

func main() {
	handlerA := InitializeHandlerA()

	user := handlerA.GetData(123)
	fmt.Println(user)

	handlerB := InitializeHandlerB()
	user = handlerB.GetData(456)
	fmt.Println(user)

	handlerC := InitializeHandlerC()
	user = handlerC.GetData(789)
	fmt.Println(user)

}
