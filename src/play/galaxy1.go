package main

import "fmt"
import "math/rand"
import "encoding/json"

type Pin struct {
	X int
	Y int
	T int
}

var seed = int64(5185441982)

func born(x int, y int) bool {
	fmt.Println("running..")
	var starCount = int(float32(x) * float32(y) * 0.05)
	fmt.Printf("Star Count: %d\n", starCount)
	rand.Seed(seed)
	starChart := make([]Pin, starCount)
	starPointer := &starChart
	for ii := 0; ii < starCount; ii++ {
		var newPin = genPin(starPointer, x, y)
		//var newPin = Pin{randInt(-x, x), randInt(-y, y), 1}
		starChart[ii] = newPin
	}
	fmt.Printf("%v\n", starChart)
	fmt.Printf("%d\n", len(starChart))
	starJSON, err := json.Marshal(starChart)
	fmt.Printf("Errors: %v\n", err)
	fmt.Printf("Star JSON: %s\n", starJSON)
	return true
}
func genPin(pins *[]Pin,x int,y int) Pin { // TODO: figure out slice pointers
	var newPin = Pin{randInt(-x, x), randInt(-y, y), 1}
	if isPinUsed(pins,newPin) {
		newPin = genPin(pins, x, y)
	}
	
	return newPin
}


func isPinUsed(pins *[]Pin,pin Pin) bool {
	for ii := 0; ii < len(pins); ii++ {
		if pins[ii].X == pin.X && pins[ii].Y == pin.Y {
			return false
		}
	}
	return true
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {

	born(15, 15)
}
