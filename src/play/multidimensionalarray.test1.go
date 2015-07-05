package main

import "fmt"

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

func main() {
	fmt.Println("Multi-dimensional Arrays Ahoy!")

  var scratch [32][32]uint8

  scratch[0][5] = 4

  fmt.Println(scratch)

  x := 32

  s2 :=matrix(x,x)

  s2[0][8] = 5
  fmt.Println(s2)

  s2 = resetMatrix(s2,x)

  fmt.Println(s2)

}
