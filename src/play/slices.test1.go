package main

import "fmt"

type Pin struct {
	X int
	Y int
	T int
}

func main() {
	s := make([]int,0)

	fmt.Println("empty:", s)

	// Lets add some elements

	s = append(s,55)
	s = append(s,66)
	s = append(s,77)
	fmt.Println("with some:", s)

	pin := make([]Pin,0)
	fmt.Println("empty pins:", pin)

	for ii := 0; ii < 5; ii++ {
		var newPin = Pin{ii,ii,1}
		pin = append(pin,newPin)
		//var newPin = Pin{randInt(-x, x), randInt(-y, y), 1}
	}
	fmt.Println("with some:", pin)
}
