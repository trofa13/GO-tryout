package main

import "fmt"


type human struct {
	name string
	age int
}

func (g human) speak() string {
	if g.age < 60 {
		return g.name + " tells some words of foolishness"
	}
	return g.name + " tells some words of wisdom"
}

func main() {
	person1 := human{name: "Alex", age: 20}
	person2 := human{name: "Anton", age: 70}

	fmt.Println(person1.speak())
	fmt.Println(person2.speak())
}