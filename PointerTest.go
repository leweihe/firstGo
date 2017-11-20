package main

import (
	"fmt"
)

func main() {
	a := -1
	var p = &a
	fmt.Println(p)
	fmt.Println(*p)

	if a < 2 {
		fmt.Println(a)
	}

	if b := 1; b > 2 {
		fmt.Println(b)
	}
	for i := 0; i < 3; i++ {
		a++
		fmt.Println(a)
	}
	fmt.Println("over")

	switch a {
	case 0:
		fmt.Println("a==0")
	case 1:

	}

	for {
		for i := 0; i < 3; i++ {
			if (i > 3) {
				break
			}
		}
		break
	}

	for a := 1; a <= 9; a++ {
		for b := 1; b <= 9; b++ {
			if b > a {
				continue
			}
			fmt.Printf(" %d * %d = %d ", a, b, a*b)
		}
		fmt.Println()
	}


}
