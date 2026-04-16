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

type ListNode struct {
	Value int
	Next  *ListNode
}

func prepend(head **ListNode, value int) {
	newNode := ListNode{Value: value}
	*head = &newNode
}

func CallPrepend() {
	var head *ListNode
	prepend(&head, 10)
	prepend(&head, 20)

	current := head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

type Config struct {
	LogLevel string
	Port     int
}

func UpdateConfig(c *Config, logLevel string, port int) {
	c.LogLevel = logLevel
	c.Port = port
}

func CallUpdateConfig() {
	// Initial configuration
	appConfig := &Config{
		LogLevel: "info",
		Port:     8080,
	}

	fmt.Println("Initial Config:", appConfig)

	// Update configuration
	UpdateConfig(appConfig, "debug", 9000)
	fmt.Println("Updated Config:", appConfig)
}
