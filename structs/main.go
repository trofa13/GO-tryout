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

func (h human) speak() string {
	if h.age < 60 {
		return h.name + " tells some words of foolishness"
	}
	return h.name + " tells some words of wisdom"
}

func (r robot) speak() string {
	return "Hello, I'm robot " + r.name
}

func main() {
	speakersList := getSpeakers()

	for _, speaker := range speakersList {
		fmt.Println(speaker.speak())
	}
}


func getSpeakers() []speaker {
	person1 := &human{name: "Alex", age: 20}
	person2 := &human{name: "Anton", age: 70}
	person3 := &human{name: "Valera", age: 15}
	robot1 := &robot{name: "Bender", material: "stainless steel"}

	list := []speaker{person1, person2, person3, robot1}
	return list
}