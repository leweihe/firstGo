package main

import "fmt"

func main() {
	fmt.Println(A)
	value := B(1, 3, 8, -1)
	fmt.Println(value)

	C(1, 2, 3, 4, 5)

	//传递slice地址拷贝
	s := []int{1, 2, 3, 4}
	D(s)
	fmt.Println(s[0])

	//传递值拷贝
	a, b := 1, 2
	C(a, b)
	fmt.Println(a)

	//闭包
	closure := closure(10)
	x := closure(4)
	fmt.Println(x)

	//defer
	fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("")
	fmt.Println("")

	//defer in loop XXXXXX
	for i := 0; i < 4; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func A(a, b, c int) (d, e, f int) {

	return a, b, c
}

func B(a ...int) (int) {
	bigger := a[0]
	for _, v := range a {
		if v > bigger {
			bigger = v
		}
	}
	fmt.Println(bigger)
	return bigger
}

func C(a ...int) {
	a[0] = 111

}

func D(s []int) {
	s[0] = 122
}

func closure(x int) (func(int) int) {
	fmt.Println()
	return func(y int) int {
		return x + y
	}
}
