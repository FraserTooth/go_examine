package main

import (
	"fmt"

	"github.com/FraserTooth/go-examine/cmd/helloworld"
)

func main() {
	greetMessageEmpty := helloworld.Hello("")
	fmt.Println(greetMessageEmpty)

	greetMessageJohn := helloworld.Hello("John")
	fmt.Println(greetMessageJohn)

}
