package tachyon

import (
	"fmt"
)

func SayHelloTachyon() {
	fmt.Println("Hello Agnes Tachyon")
}

type Uma struct {
	Name   string
	Weight int
	Height int
	Race   string
}

func Newuma() {
	uma1 := Uma{
		Name:   "Agnes Tachyon",
		Weight: 55,
		Height: 165,
		Race:   "Long",
	}
	fmt.Printf("Uma: %+v\n", uma1)
}
