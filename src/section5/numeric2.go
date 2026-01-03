package main

import (
	"fmt"
)

func main() {

	// ASCII
	var char1 byte = 72
	var char2 byte = 0110
	var char3 byte = 0x48

	// Unicode
	var char4 rune = 50556
	var char5 rune = 0142574
	var char6 rune = 0xC57C

	fmt.Printf("char1: %c, char2: %c, char3: %c\n", char1, char2, char3)
	fmt.Printf("char1: %d, char2: %d, char3: %d\n", char1, char2, char3)
	fmt.Printf("char1: %d, char2: %o, char3: %x\n", char1, char2, char3)

	fmt.Printf("char4: %c, char5: %c, char6: %c\n", char4, char5, char6)
	fmt.Printf("char4: %d, char5: %d, char6: %d\n", char4, char5, char6)
	fmt.Printf("char4: %d, char5: %o, char6: %x\n", char4, char5, char6)

}
