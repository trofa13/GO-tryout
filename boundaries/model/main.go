package model


type human struct {
	name string
	age int
}


func (h human) Speak() string {
	if h.age < 60 {
		return h.name + " tells some words of foolishness"
	}
	return h.name + " tells some words of wisdom"
}

type robot struct {
	name string
	material string
}

func (r robot) Speak() string {
	return "Hello, I'm robot " + r.name
}

type speaker interface {
	Speak() string
}

func GetSpeakers() []speaker {
	person1 := &human{name: "Alex", age: 20}
	person2 := &human{name: "Anton", age: 70}
	person3 := &human{name: "Valera", age: 15}
	robot1 := &robot{name: "Bender", material: "stainless steel"}

	list := []speaker{person1, person2, person3, robot1}
	return list
}