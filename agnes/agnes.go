package agnes

import (
	"fmt"
	"github.com/AgnesTachyon/Gobeyone/agnes/internal/tachyon"
)

func SayHelloAgnes() {
	fmt.Println("Uma musume")
	//generateTest()
	tachyon.SayHelloTachyon()
}

func Agnesforloop() {
	var myslice []int
	for i := 0; i < 3; i++ {
		myslice = append(myslice, i)
	}
	fmt.Println(myslice)
}

func AgnesDowhile() {
	i := 22
	var myslice []int
	for {
		myslice = append(myslice, i)
		i--
		if i <= 11 {
			fmt.Println(myslice)
			break
		}
	}
}

func AgnesWhile() {
	i := 5
	var myslice []int
	for i <= 8 {
		myslice = append(myslice, i)
		i++
	}
	fmt.Println(myslice)
}
