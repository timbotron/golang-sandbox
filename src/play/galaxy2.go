package main

import "fmt"
import "time"
import "math/rand"
import "encoding/json"
import "io/ioutil"
//
type Pin struct {
	SX int16
	SY int16
	X uint8
	Y uint8
	C uint16
}
//
// // NEEDS LOTSA WORK
//
var seed = int64(19820514)
//

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func matrix(x, y int) [][]uint8 {
	mat := make([]uint8, x*y)
	rows := make([][]uint8, y) // Assuming row-based
	for i := range rows {
		rows[i] = mat[i*x : (i+1)*x]
	}
	return rows
}

func resetMatrix(t [][]uint8,l int) [][]uint8 {
	for xx := 0; xx < l; xx++ {
    for yy := 0; yy < l; yy++ {
    t[xx][yy] = 0
    }
  }
  return t
}

func genPin(scratch *[][]uint8,l int,md int) Pin {
	var newPin = Pin{0,0,uint8(rand.Intn(l)), uint8(rand.Intn(l)), 1}

	for xx := (int(newPin.X) - md); xx < (int(newPin.X) + md); xx++ {
		for yy := (int(newPin.Y) - md); yy < (int(newPin.Y) + md); yy++ {
			if(xx < 0 && xx > l && yy < 0 && yy > l && (*scratch)[xx][yy] == 1) {
				return genPin(scratch,l,md);
			}
		}
	}

	return newPin
}

func born(config map[string]interface{}) bool {
	fmt.Println("running..")
	ss := int(config["sector_size"].(float64))
	cs := int(config["cradle_size"].(float64))
	min := int(config["min_stars_per_sector"].(float64))
	max := int(config["max_stars_per_sector"].(float64))
	md := int(config["min_star_distance"].(float64))
	var ns int
	scratch := matrix(ss,ss)

	fmt.Println(scratch)
	t0 := time.Now()
	rand.Seed(seed)
	starChart := make([]Pin, 0)
	for xx := -cs; xx <= cs; xx++ {
		for yy := -cs; yy <= cs; yy++ {
			// Now we are creating a sector!
			// We create a blank multideminsional array to hold the data in addition to storing it in starChart, so we can easily make sure neighboring stars are not an issue
			scratch = resetMatrix(scratch,ss)
			// rand to find the num of stars in the sector, between the min and max
			ns = randInt(min,max)

			fmt.Println(ns)

			// loop x times, discarding star and retrying if it fails the neighbor test.
			for aa := 0; aa < ns; aa++ {
				var newPin = genPin(&scratch,ss,md)
				newPin.SX = int16(xx)
				newPin.SY = int16(yy)
				scratch[newPin.X][newPin.Y] = 1
				starChart = append(starChart,newPin)

			}

		}
	}
	fmt.Println("Last Sector: ",scratch)
	//fmt.Println("All Stars: ",starChart)
	// starJSON, _ := json.Marshal(starChart)
	// fmt.Printf("Star JSON: %s\n", starJSON)
	fmt.Println("Number of stars: ",len(starChart))
	t1 := time.Now()
	fmt.Printf("Galaxy birth took %v to run.\n", t1.Sub(t0))
	return true
}


func main() {
	// read in config
	byt, _ := ioutil.ReadFile("config.mk2.json")
  //check(err)
  var config map[string]interface{}

  _ = json.Unmarshal(byt, &config)
  fmt.Printf("%v\n",config)
  fmt.Println(config["sector_size"])

	// call born with sector size to gen sector_size
	born(config)
}
