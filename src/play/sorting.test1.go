// This uses sort and a divide and conquer search; could be handy for reference

package main

import "fmt"
import "sort"
import "time"
import "math/rand"
// import "encoding/json"

type Pin struct {
	X int
	Y int
	T int
}

type ByX []Pin

func (a ByX) Len() int           { return len(a) }
func (a ByX) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByX) Less(i, j int) bool { return a[i].X < a[j].X }

var seed = int64(5185441982)

func born(x int, y int) bool {
	fmt.Println("running..")
	t0 := time.Now()
	var starCount = int(float32(x) * float32(y) * 0.01)
	fmt.Printf("Star Count: %d\n", starCount)
	rand.Seed(seed)
	starChart := make([]Pin, 0)
	for ii := 0; ii < starCount; ii++ {
		fmt.Println("Creating Star # ",ii)
		var newPin = genPin(&starChart, x, y)
		starChart = append(starChart, newPin)
		  sort.Sort(ByX(starChart))
	}
	//  starJSON, _ := json.Marshal(starChart)
	// fmt.Printf("Star JSON: %s\n", starJSON)
	t1 := time.Now()
	fmt.Printf("Galaxy birth took %v to run.\n", t1.Sub(t0))
	return true
}
func genPin(pins *[]Pin,x int,y int) Pin { 
	var newPin = Pin{randInt(-x, x), randInt(-y, y), 1}
	//if isPinUsed(pins,newPin) {
	if isPinUsed(pins,newPin,0,len(*pins)) {
		newPin = genPin(pins, x, y)
	}
	
	return newPin
}

// WORKING ONE, want to try w/Divide & Conquer
// func isPinUsed(pins *[]Pin,pin Pin) bool {
// 	for ii := 0; ii < len(*pins); ii++ {
// 		if (*pins)[ii].X == pin.X && (*pins)[ii].Y == pin.Y {
// 			return true
// 		}
// 	}
// 	return false
// }

func isPinUsed(pins *[]Pin,pin Pin, min int, max int) bool {
	diff := (max - min) / 2
	mid := min + diff
	//fmt.Printf("%v %d %d %d %d\n", pin, min, max, diff, mid)
	if diff < 100 {
		for ii := min; ii < max; ii++ {
			if (*pins)[ii].X == pin.X && (*pins)[ii].Y == pin.Y {
				return true
			}
		}
		return false
	} else if pin.X > (*pins)[mid].X {
		return isPinUsed(pins,pin,mid,max)
	} else {
		return isPinUsed(pins,pin,min,mid)
	}

	return true
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {

	born(10000, 10000)
}
