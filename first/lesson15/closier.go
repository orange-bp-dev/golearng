package main

import "fmt"

func incrementGenerator() (func() int) {
	x:=0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64{
	return func(radius float64) float64{
		return pi * radius * radius
	}
}

func main() {
	counter := incrementGenerator()
	c1 := circleArea(3.14)
	fmt.Println("c1=", c1(2))
	c2 := circleArea(3)
	fmt.Println("c2=", c2(2))


	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}