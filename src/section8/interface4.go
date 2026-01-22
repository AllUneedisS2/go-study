package main

import "fmt"

type dog struct {
	name   string
	weight int
}

func (d dog) run() {
	fmt.Println(d.name, "- dog go!")
}

type cat struct {
	name   string
	weight int
}

func (c cat) run() {
	fmt.Println(c.name, "- cat go!")
}

// 익명 인터페이스
func act(animal interface{ run() }) {
	animal.run()
}

func main() {

	dog1 := dog{name: "Buddy", weight: 30}
	act(dog1)

	cat1 := cat{name: "Kitty", weight: 10}
	act(cat1)

}
