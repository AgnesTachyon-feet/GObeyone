package tachyon

import (
	"fmt"
)

type Student struct {
	Firstname string
	Lastname  string
}

func (s Student) Fullname() string {
	return s.Firstname + " " + s.Lastname
}

func CallStudent() {
	student := Student{
		Firstname: "Agnes",
		Lastname:  "Tachyon",
	}
	fullName := student.Fullname()
	fmt.Println("Full name of the student:", fullName)
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func CallRectangle() {
	rect := Rectangle{
		Length: 10,
		Width:  5,
	}
	area := rect.Area()
	fmt.Println("Area of rectangle: ", area)
}

// polymorphism

type Speaker interface {
	Speak() string
}
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello!"
}

func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func MakeSound() {
	dog := Dog{Name: "Denji"}
	person := Person{Name: "Makima"}

	makeSound(dog)
	makeSound(person)
}
