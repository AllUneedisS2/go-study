package main

import "fmt"

func Enumeration2() {

	const (
		_       = iota + 0.75*2 // 0
		DEFAULT                 // 1
		SILVER                  // 2
		GOLD                    // 3
		_                       // PLATINUM 멤버십 건너띔
		DIAMOND                 // 5
	)

	fmt.Println("DEFAULT:", DEFAULT)
	fmt.Println("SILVER:", SILVER)
	fmt.Println("GOLD:", GOLD)
	fmt.Println("DIAMOND:", DIAMOND)

}
