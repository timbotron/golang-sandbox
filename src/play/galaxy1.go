package main

import "fmt"
import "time"
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
	t0 := time.Now()
	var starCount = int(float32(x) * float32(y) * 0.05)
	fmt.Printf("Star Count: %d\n", starCount)
	rand.Seed(seed)
	starChart := make([]Pin, 0)
	for ii := 0; ii < starCount; ii++ {
		var newPin = genPin(&starChart, x, y)
		starChart = append(starChart, newPin)
		// starChart[ii] = newPin
	}
	 starJSON, _ := json.Marshal(starChart)
	fmt.Printf("Star JSON: %s\n", starJSON)
	t1 := time.Now()
	fmt.Printf("Galaxy birth took %v to run.\n", t1.Sub(t0))
	return true
}
func genPin(pins *[]Pin,x int,y int) Pin { 
	var newPin = Pin{randInt(-x, x), randInt(-y, y), 1}
	if isPinUsed(pins,newPin) {
		newPin = genPin(pins, x, y)
	}
	
	return newPin
}


func isPinUsed(pins *[]Pin,pin Pin) bool {
	for ii := 0; ii < len(*pins); ii++ {
		if (*pins)[ii].X == pin.X && (*pins)[ii].Y == pin.Y {
			return true
		}
	}
	return false
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {

	born(1000, 1000)
}
