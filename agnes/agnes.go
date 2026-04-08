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

func Agnesmap() {
	myMap := make(map[string]int)
	myMap["apple"] = 6
	myMap["banana"] = 7
	fmt.Println("Apple:", myMap["apple"])
}

func Agnesloopmap() {
	myMap := make(map[string]int)
	myMap["apple"] = 6
	myMap["banana"] = 7
	myMap["orange"] = 8
	fmt.Println("Orange:", myMap["orange"])
	delete(myMap, "orange")
	for key, value := range myMap {
		fmt.Printf("%s -> %d\n", key, value)
	}
	val, ok := myMap["pear"]
	if ok {
		fmt.Println("Pear's value:", val)
	} else {
		fmt.Println("Pear not found in map")
	}
}

func NewUma() {
	tachyon.Newuma()
	fmt.Println("--------------------")
	tachyon.NewUmaArrey()
	tachyon.NewUmaMap()
}
