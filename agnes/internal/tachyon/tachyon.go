package tachyon

import (
	"fmt"
	"reflect"
)

func SayHelloTachyon() {
	fmt.Println("Hello Agnes Tachyon")
}

type Uma struct {
	Name   string
	Weight int
	Height int
	Race   string
	Racing Racing
}

func Newuma() {
	uma1 := Uma{
		Name:   "Agnes Tachyon",
		Weight: 55,
		Height: 165,
		Race:   "Long",
		Racing: Racing{
			Name:       "Tokyo Racecourse",
			Distance:   2000,
			RaceType:   "G1",
			GroundType: "Turf",
		},
	}
	fmt.Printf("Uma: %+v\n", uma1)
}

func NewUmaArrey() {
	uma := [3]Uma{}
	uma[0] = Uma{
		Name:   "Agnes Tachyon",
		Weight: 55,
		Height: 165,
		Race:   "Long",
		Racing: Racing{
			Name:       "Tokyo Racecourse",
			Distance:   2000,
			RaceType:   "G1",
			GroundType: "Turf",
		},
	}
	uma[1] = Uma{
		Name:   "Manhattan Cafe",
		Weight: 50,
		Height: 160,
		Race:   "Medium",
		Racing: Racing{
			Name:       "Osaka Racecourse",
			Distance:   1600,
			RaceType:   "G2",
			GroundType: "Turf",
		},
	}
	uma[2] = Uma{
		Name:   "Gold Ship",
		Weight: 60,
		Height: 170,
		Race:   "Long",
		Racing: Racing{
			Name:       "Kyoto Racecourse",
			Distance:   2400,
			RaceType:   "G3",
			GroundType: "Turf",
		},
	}
	for i := 0; i < len(uma); i++ {
		fmt.Println("Uma", (i + 1))
		fmt.Println("Name: ", uma[i].Name)
		fmt.Println("Weight: ", uma[i].Weight)
		fmt.Println("Height: ", uma[i].Height)
		fmt.Println("Race: ", uma[i].Race)
		v := reflect.ValueOf(uma[i].Racing)
		typeOfS := v.Type()

		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("%s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
		}
		fmt.Println("--------------------")
	}
}

func NewUmaMap() {
	uma := make(map[string]Uma)
	uma["u1"] = Uma{
		Name:   "Agnes Tachyon",
		Weight: 55,
		Height: 165,
		Race:   "Long",
		Racing: Racing{
			Name:       "Tokyo Racecourse",
			Distance:   2000,
			RaceType:   "G1",
			GroundType: "Turf",
		},
	}
	uma["u2"] = Uma{
		Name:   "Manhattan Cafe",
		Weight: 50,
		Height: 160,
		Race:   "Medium",
		Racing: Racing{
			Name:       "Osaka Racecourse",
			Distance:   1600,
			RaceType:   "G2",
			GroundType: "Turf",
		},
	}
	uma["u3"] = Uma{
		Name:   "Gold Ship",
		Weight: 60,
		Height: 170,
		Race:   "Long",
		Racing: Racing{
			Name:       "Kyoto Racecourse",
			Distance:   2400,
			RaceType:   "G3",
			GroundType: "Turf",
		},
	}
	fmt.Println("--- Uma with Map ---")
	for key, v := range uma {
		fmt.Printf("ID: %s\n", key)
		fmt.Printf("  Name: %s\n", v.Name)
		fmt.Printf("  Weight: %d kg\n", v.Weight)
		fmt.Printf("  Racing at: %s (%d m)\n", v.Racing.Name, v.Racing.Distance)
		fmt.Println("--------------------")
	}
}

type Racing struct {
	Name       string
	Distance   int
	RaceType   string
	GroundType string
}
