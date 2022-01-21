package main

import (
	"fmt"

	calcv1 "github.com/kishoredevops55/go_math/calc"
	"github.com/kishoredevops55/go_math/geometry"
	calcv2 "github.com/kishoredevops55/go_math/v2/calc"
)

func main() {
	fmt.Println(geometry.CubeVoulme(5))
	s := calcv2.Add(1, 2, 3, 4)
	fmt.Println(s)
	s1 := calcv1.Add(1, 2)
	fmt.Println(s1)

}
