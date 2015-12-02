package main

import (
	"fmt"
)

func main() {
	params := Params{}
	over := make(chan bool)
	InitParams(&params)

	fmt.Println(params)
}
