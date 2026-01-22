package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

func (d Dog) bite() {
	fmt.Println(d.name, "bites!")
}

type behavior interface {
	bite()
}

func main() {

	// 1
	dog1 := Dog{name: "Buddy", weight: 30}
	var inter1 behavior
	inter1 = dog1
	inter1.bite()

	// 2
	dog2 := Dog{name: "Max", weight: 25}
	inter2 := behavior(dog2)
	inter2.bite()

	// 3
	inters := []behavior{dog1, dog2}
	for idx, _ := range inters {
		inters[idx].bite()
	}

	// 4
	for _, inter := range inters {
		inter.bite()
	}

}
