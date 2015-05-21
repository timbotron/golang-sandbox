package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	pa := &a
	fmt.Println("orig:",a)
	fmt.Println("pointer:",pa)
	fmt.Println("orig_len:",len(a))
	fmt.Println("pointer_len:",len(pa))
	calcLen := lenfunc(pa)
	fmt.Println("lenfunc:",calcLen)
}

func lenfunc(thing *[5]int) int {
	return len(thing)
}