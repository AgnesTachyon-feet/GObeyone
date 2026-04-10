package tachyon

import (
	"fmt"
)

func SayFunk(name string) {
	fmt.Printf("Hellyeah %s\n", name)
}

func add(a int, b int) int {
	return a + b
}

func AddSum() {
	number1 := 3
	number2 := 5
	sumNumbers := add(number1, number2)
	fmt.Println(sumNumbers)
}
