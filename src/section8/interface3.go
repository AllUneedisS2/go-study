package main

import "fmt"

type dog struct {
	name   string
	weight int
}

func (d dog) bite() {
	fmt.Println(d.name, "- dog bite!")
}
func (d dog) sound() {
	fmt.Println(d.name, "- dog bark!")
}
func (d dog) run() {
	fmt.Println(d.name, "- dog go!")
}

type cat struct {
	name   string
	weight int
}

func (c cat) bite() {
	fmt.Println(c.name, "- cat bite!")
}
func (c cat) sound() {
	fmt.Println(c.name, "- cat cry!")
}
func (c cat) run() {
	fmt.Println(c.name, "- cat go!")
}

type behavior interface {
	bite()
	sound()
	run()
}

func act(animal behavior) {
	animal.bite()
	animal.sound()
	animal.run()
}

func main() {

	dog1 := dog{name: "Buddy", weight: 30}
	act(dog1)

	cat1 := cat{name: "Kitty", weight: 10}
	act(cat1)

}
