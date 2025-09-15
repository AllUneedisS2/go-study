package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i > 50:
		fmt.Println("i is greater than 50:", i)
	case i <= 20:
		fmt.Println("i is 20 or less:", i)
	default:
		fmt.Println("i is between 21 and 50:", i)
	}

}
