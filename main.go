package main

import (
	"fmt"

	"github.com/AgnesTachyon/Gobeyone/agnes"
)

func main() {
	fullname := "Agnes Tachyon"
	var age int = 11
	agnes.SayHelloAgnes()
	fmt.Printf("Welcome to the world of feet %s Yay! age: %d\n", fullname, age)
	fullname = "Manhattan" + " Cafe"
	age = age * 2
	fmt.Println(fullname, age)
}
