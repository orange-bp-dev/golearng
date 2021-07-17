package main

import "fmt"

type Vertex struct {
	X, Y int
	S string
}

func changeVertex(v Vertex){
	v.X = 1000
}

func changeVertex2(v *Vertex){
	v.X = 1000
}

func main () {
	// v := Vertex{X: 1, Y: 2}
	// v.X = 100

	// v2 := Vertex{100, 200, "test"} 
	// fmt.Println(v2)

	// //v3 == v4 
	// v3 := Vertex{} 
	// fmt.Println(v3)

	// var v4 Vertex
	// fmt.Println(v4)

  // // v5 == v6 ポインタが返ってくる
	// v5 := new(Vertex)
	// fmt.Printf("%T %v\n", v5, v5)

	// v6 := &Vertex{}
	// fmt.Printf("%T %v\n,", v6, v6)

	v := Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println(v)

	v2 := &Vertex{1, 2, "test"}
	changeVertex2(v2)
	fmt.Println(v2) // 1 => 1000
}