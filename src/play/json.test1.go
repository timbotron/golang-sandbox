package main

import "fmt"
import "encoding/json"
import "io/ioutil"

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func main() {

	byt, err := ioutil.ReadFile("json.test1.json")
    check(err)
    var dat map[string]interface{}

    err = json.Unmarshal(byt, &dat)
    fmt.Printf("%v\n",dat)
    fmt.Println(dat["sector_size"])
}
