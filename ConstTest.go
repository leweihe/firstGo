package main

import "fmt"

const a = 1.3
const sss = "123"
const (
	b       = a - 1
	c       = "hello world"
	d, e, f = 1, 23, 4
	g       = len(sss)
	h       = iota
)
const (
	i = iota
)
const (
	j = iota
	k
	l
	m
)

func main() {
	fmt.Println(a, b, c, d, e, f, g, h, i)
	fmt.Println(j, k, l, m)
}
