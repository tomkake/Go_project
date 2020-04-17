// +build ignore
// +debug ignore
package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func pi() float32 {
	return (math.Pi)
}

func main() {
	fmt.Printf("%d\n", add(3, 4))
	fmt.Printf("%f", pi())
}
