package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randNum() {
	rand.Seed(time.Now().UnixNano())
	s1 := rand.Intn(46) + 5
	var s2 int

	for s1 != s2 {
		fmt.Println("Unesi neki broj: ")
		fmt.Scanln(&s2)
		if s2 == s1 {
			fmt.Println("Brojevi su isti")
		} else if s2 > s1 {
			fmt.Println("Broj koji si unio je veci od random broja")
		} else {
			fmt.Println("Broj koji si unio je manji od random broja")
		}
	}
}

func main() {
	randNum()
}
