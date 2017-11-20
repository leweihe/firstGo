package main

import (
	"fmt"
	"math"
)

type (
	ByteSize int64
	byte int8
)

var outInt, outInt2 int64

var (
	outBool bool
	outInt4 int
)

func main() {
	var a bool
	var abc ByteSize = 1

	ax := 112

	var bx, cx int8 = 123, 123
	a = true
	fmt.Println(a)
	fmt.Println(math.MaxInt32)
	fmt.Println(abc)

	ax = 999999999999
	//运行错误 bx 为int8
	//bx = 999999999999
	fmt.Println(ax)
	fmt.Println(bx)
	fmt.Println(cx)
	fmt.Println(outInt)
	fmt.Println(outInt2)
	fmt.Println(outBool)
	outInt4 = 999
	//fmt.Println(outInt4)


	changeA := 1
	changeB := float32(changeA)
	fmt.Println(changeB)
	fmt.Print("hello world")
}
