package main

import "fmt"


type human struct {
	name string
	age int
}

type robot struct {
	name string
	material string
}

type speaker interface {
	speak() string
}

func (g human) speak() string {
	if g.age < 60 {
		return g.name + " tells some words of foolishness"
	}
	return g.name + " tells some words of wisdom"
}

func main() {
	humansList := getHumans()

	for _, person := range humansList {
		fmt.Println(person.speak())
	}
}


func getHumans() []*human {
	person1 := &human{name: "Alex", age: 20}
	person2 := &human{name: "Anton", age: 70}
	person3 := &human{name: "Valera", age: 15}

	list := []*human{person1, person2, person3}
	return list
}