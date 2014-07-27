package main

import (
	"./moke"
	"fmt"
)

func main() {
	v := moke.Hello()
	fmt.Printf("value is (%s)\n", v)
}
