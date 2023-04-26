package main

import "fmt"

type Animal interface {
	Name() string
	Speak() string
}

type Dog struct{}

func (d Dog) Name() string {
	return "Dog"
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Name() string {
	return "Cat"
}
func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Println(animal.Name(), "-->", animal.Speak())
	}
}
