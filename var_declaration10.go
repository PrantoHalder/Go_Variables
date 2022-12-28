package main

import (
	"fmt"
	"math"
)

func var10(){
	a,b := 100,150
	c := math.Min(float64(a),float64(b))
	fmt.Println(c)
}