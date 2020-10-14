package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v, S0 float64) func(float64) float64 {

	fn := func(t float64) float64 {
		return a*math.Pow(t, 2)/2 + v*t + S0
	}
	return fn
}

func main() {
	var a, v, S0, t float64

	fmt.Println("Enter some value for acceleration")
	fmt.Scanln(&a)
	fmt.Println("Enter some value for velocity")
	fmt.Scanln(&v)
	fmt.Println("Enter some value for initial displacement")
	fmt.Scanln(&S0)

	fn := GenDisplaceFn(a, v, S0)

	fmt.Println("Enter some value for time")
	fmt.Scanln(&t)

	fmt.Println("Calculated displacement :", fn(t))

}
