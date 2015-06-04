package main

import "fmt"
import "time"
import "math/rand"
// import "encoding/json"

type Pin struct {
	X int
	Y int
	T int
}

var seed = int64(19820514)

func born(x int, y int) bool {
	fmt.Println("running..")
	t0 := time.Now()
	rand.Seed(seed)
	starChart := make([]Pin, 0)
	for xx := -x; xx < x; xx++ {
		for yy := -y; yy < y; yy++ {
			if rand.Intn(99) < 1 {
				newStar := genStar(xx,yy)
				//fmt.Printf("%d\t%d\n",newStar.X,newStar.Y)

				starChart = append(starChart,newStar)
			}
		}
	}
	// starJSON, _ := json.Marshal(starChart)
	// fmt.Printf("Star JSON: %s\n", starJSON)
	fmt.Println("Number of stars: ",len(starChart))
	t1 := time.Now()
	fmt.Printf("Galaxy birth took %v to run.\n", t1.Sub(t0))
	return true
}

func genStar(x int, y int) Pin {
	// gen planet type
	t := 1
	return Pin{x, y, t}
}

func main() {

	born(100, 100)
}
