package main

import (
	"fmt"

	"di.com/m/v2/handlers/a"
	"di.com/m/v2/handlers/b"
)

func main() {

	h1 := a.NewHandler()
	user := h1.GetData(123)
	fmt.Println(user)

	h2 := b.NewHandler()
	user = h2.GetData(456)
	fmt.Println(user)
}
