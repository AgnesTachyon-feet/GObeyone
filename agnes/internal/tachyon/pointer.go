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

type Mama struct {
	Name string
}

func changeName(p *Mama) {
	p.Name = "Glock"
}

func ChangePersonName() {
	person := Mama{Name: "Agnes"}
	changeName(&person)
	fmt.Println(person.Name)
}

type Employee struct {
	Name   string
	Salary int
}

func giveRaise(e *Employee, raise int) {
	e.Salary += raise
}

func CallRaise() {
	emp := Employee{
		Name:   "Agnes Tachyon",
		Salary: 444,
	}
	giveRaise(&emp, 100)
	fmt.Printf("Employee: %s, Salary: %d\n", emp.Name, emp.Salary)
}
