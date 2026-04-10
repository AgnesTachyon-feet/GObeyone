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
