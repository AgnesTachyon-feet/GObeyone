package tachyon

import (
	"fmt"
)

func PointerAgnes() {
	x := 10
	var p *int = &x
	fmt.Println("Value of x: ", x)
	fmt.Println("Value at p: ", *p)
	*p = 20
	fmt.Println("New value of x: ", x)

}

func changeValue(ptr *int) {
	*ptr = 50
}

func CallchangeValue() {
	x := 20
	changeValue(&x)
	fmt.Println("Value of x after changeValue: ", x)
}
