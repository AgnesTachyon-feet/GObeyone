package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/AgnesTachyon/Gobeyone/agnes"

)

func main() {
	id := uuid.New()
	fmt.Println("Hello, World")
	fmt.Printf("UUID: %s\n", id)
	agnes.SayHelloAgnes()
}