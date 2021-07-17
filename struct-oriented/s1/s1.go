package main

import "fmt"

//小文字のxyにすると他のパッケージからは見れないようになる
type Vertex struct {
	X, Y int
}

// Value(値) Receiver
func(v Vertex) Area() int {
	return v.X * v.Y
}

// PointerReceiver: ポインタを使って変数の中身を書き換えれる
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

//x, y, z の３方向--------------------------
type Vertex3D struct {
	//Embedded
	Vertex
	Z int
}


func(v Vertex3D) Area3D() int {
	return v.X * v.Y * v.Z
}


func (v *Vertex3D) Scale3D(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
	v.Z = v.Z * i
}
//---------------------------

func Area(v Vertex) int {
	return v.X * v.Y
}

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	v :=New(3, 4, 5)
	// v := Vertex{3, 4}
	//fmt.Println(Area(v))
	v.Scale(10)
	//struct指向の書き方
	fmt.Println(v.Area())
	fmt.Println(v.Area3D())
}