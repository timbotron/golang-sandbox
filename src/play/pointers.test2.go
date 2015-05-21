package main

import "fmt"

type Foo struct {
	X int
	Y int
}

func main() {
	a := make([]Foo,5)
	a[0].X = 1
	a[0].Y = 1
	pa := &a
	fmt.Println("a len:",len(a))
	fmt.Println("pa len:",len(*pa))
	fmt.Println("orig:",a)
	fmt.Println("pointer:",pa)
	fmt.Println("orig_len:",len(a))
	fmt.Println("pointer_len:",len(*pa))
	calcLen := lenfunc(pa)
	fmt.Println("lenfunc:",calcLen)
}

func lenfunc(thing *[]Foo) int {
	return len(*thing)
}