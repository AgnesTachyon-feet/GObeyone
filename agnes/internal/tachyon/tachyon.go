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

func NewUmaArrey() {
	uma := [3]Uma{}
	uma[0] = Uma{
		Name:   "Agnes Tachyon",
		Weight: 55,
		Height: 165,
		Race:   "Long"}
	uma[1] = Uma{
		Name:   "Manhattan Cafe",
		Weight: 50,
		Height: 160,
		Race:   "Medium"}
	uma[2] = Uma{
		Name:   "Gold Ship",
		Weight: 60,
		Height: 170,
		Race:   "Long"}
	fmt.Println("Uma List:")
	for _, v := range uma {
		fmt.Printf("%+v\n", v)
	}
}

func NewUmaMap() {
	uma := make(map[string]Uma)
	uma["u1"] = Uma{
		Name:   "Agnes Tachyon",
		Weight: 55,
		Height: 165}
	uma["u2"] = Uma{
		Name:   "Manhattan Cafe",
		Weight: 50,
		Height: 160,
		Race:   "Medium"}
	uma["u3"] = Uma{
		Name:   "Gold Ship",
		Weight: 60,
		Height: 170,
		Race:   "Long"}
	fmt.Println("Uma Map:")
	for key, value := range uma {
		fmt.Printf("%s: %+v\n", key, value)
	}
	fmt.Println("uma u1:", uma["u1"])
}
