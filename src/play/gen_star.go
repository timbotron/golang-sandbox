package main

import "fmt"
//import "math/rand"

type st struct {
	cls string
	min_t uint16
	max_t uint16
	color string
	min_m float32
	max_m float32
	min_r float32
	max_r float32
	min_l float32
	max_l float32
}

func gen_map() map[string]st {
	// we are going to generate all of the startypes

	r := make(map[string]st)

	r["g"] = st{cls: "a",
				min_t: 52,
				max_t: 60,
				color: "yellowish white",
				min_m: 0.8,
				max_m: 1.04,
				min_r: 0.96,
				max_r: 1.15,
				min_l: 0.6,
				max_l: 1.5}
	return r
}

func gen_star(m map[string]st) {
	//var r uint16
	//r = rand.Intn(10000)
	fmt.Printf("%v\n",m)
	fmt.Printf("%v\n",m["g"])
	fmt.Printf("%v\n",m["g"].color)
}

func main() {

	m := gen_map()

	gen_star(m)

	//fmt.Println("hello",s)
}
