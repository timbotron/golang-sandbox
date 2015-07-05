package main

import "fmt"
//import "time"
// import "math/rand"
import "encoding/json"
import "io/ioutil"
//
// type Pin struct {
// 	X int16
// 	Y int16
// 	T uint8
// }
//
// // NEEDS LOTSA WORK
//
// var seed = int64(19820514)
//

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
func born(config map[string]interface{}) bool {
	fmt.Println("running..")
	s := int(config["cradle_size"].(float64))
	scratch := matrix(s,s)

	fmt.Println(scratch)
	// t0 := time.Now()
	// rand.Seed(seed)
	// starChart := make([]Pin, 0)
	// for xx := -config["cradle_size"]; xx < config["cradle_size"]; xx++ {
	// 	for yy := -config["cradle_size"]; yy < config["cradle_size"]; yy++ {
	// 		// Now we are creating a sector!
	// 			fmt.Println("Generating Sector: %i, %i",xx,yy)
	//
	// 		// We create a blank multideminsional array to hold the data in addition to storing it in starChart, so we can easily make sure neighboring stars are not an issue
	//
	// 		// rand to find the num of stars in the sector, between the min and max
	//
	// 		// loop x times, discarding star and retrying if it fails the neighbor test.
	//
	// 		// if rand.Intn(99) < 1 {
	// 		// 	newStar := genStar(xx,yy)
	// 		// 	//fmt.Printf("%d\t%d\n",newStar.X,newStar.Y)
	// 		//
	// 		// 	starChart = append(starChart,newStar)
	// 		// }
	// 	}
	// }
	// starJSON, _ := json.Marshal(starChart)
	// fmt.Printf("Star JSON: %s\n", starJSON)
	// fmt.Println("Number of stars: ",len(starChart))
	// t1 := time.Now()
	// fmt.Printf("Galaxy birth took %v to run.\n", t1.Sub(t0))
	return true
}
//
// func genStar(x int16, y int16) Pin {
// 	// gen planet type
// 	var t uint8
// 	t = 1
// 	return Pin{x, y, t}
// }

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
