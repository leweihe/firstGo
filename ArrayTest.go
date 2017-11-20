package main

import "fmt"

func main() {

	var a [9]int
	var b = [9]int{}
	var c = [...]int{3: 1, 8: 2, 23: 1}
	for i := 0; i < 9; i++ {
		a[i] = i
		fmt.Println(b[i] - a[i])
	}
	fmt.Println(c)
	fmt.Println(b)

	fmt.Println(a == b)

	//x := [1]int{}
	//y := new([1]int)

	//fmt.Println(x == y)

	//for i := 0; i < a.len(); i++ {
	//
	//}

	xx := [...]int{2, 123, 421, 12, 3, 41, 5, 12}
	fmt.Println(xx)
	num := len(xx)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if xx[i] < xx[j] {
				tmp := xx[i]
				xx[i] = xx[j]
				xx[j] = tmp
			}
		}
	}
	fmt.Println(xx)

}
