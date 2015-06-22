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
	// Data: https://en.wikipedia.org/wiki/Stellar_classification#Harvard_spectral_classification

	r := make(map[string]st)

	
	r["b"] = st{cls: "b",
				min_t: 100,
				max_t: 300,
				color: "deep blue white",
				min_m: 2.1,
				max_m: 16,
				min_r: 1.8,
				max_r: 6.6,
				min_l: 25,
				max_l: 30000}
	r["a"] = st{cls: "a",
				min_t: 75,
				max_t: 100,
				color: "blue white",
				min_m: 1.4,
				max_m: 2.1,
				min_r: 1.4,
				max_r: 1.8,
				min_l: 5,
				max_l: 25}
	r["f"] = st{cls: "f",
				min_t: 60,
				max_t: 75,
				color: "white",
				min_m: 1.04,
				max_m: 1.4,
				min_r: 1.15,
				max_r: 1.4,
				min_l: 1.5,
				max_l: 5}
	r["g"] = st{cls: "g",
				min_t: 52,
				max_t: 60,
				color: "yellowish white",
				min_m: 0.8,
				max_m: 1.04,
				min_r: 0.96,
				max_r: 1.15,
				min_l: 0.6,
				max_l: 1.5}
	r["k"] = st{cls: "k",
				min_t: 37,
				max_t: 52,
				color: "pale yellow orange",
				min_m: 0.45,
				max_m: 0.8,
				min_r: 0.7,
				max_r: 0.96,
				min_l: 0.08,
				max_l: 0.6}
	r["m"] = st{cls: "m",
				min_t: 24,
				max_t: 37,
				color: "light orange red",
				min_m: 0.08,
				max_m: 0.45,
				min_r: 0.01,
				max_r: 0.7,
				min_l: 0.001,
				max_l: 0.08}
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
