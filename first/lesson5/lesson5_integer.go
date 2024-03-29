package main

import "fmt"

func main() {
	var (
		u8 uint8      = 255
		i8 int8       = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)

	fmt.Println(u8, i8, f32, c64)

	//%v： 値を表示
	fmt.Printf("type=%T value=%v", u8, u8)

	fmt.Println(10/3)
	// => 3 (小数点以下切り捨て)

	fmt.Println(10.0/3)
	fmt.Println(10/3.0)
	// => 3.33333...

	x := 0
	x++
	// => x = x + 1
	x--
	// => x = x - 1

	fmt.Println(1 << 0) // 0001 0001 => 1
	fmt.Println(1 << 1) // 0001 0010 => 2
  fmt.Println(1 << 2) // 0001 0100 => 4
	fmt.Println(1 << 3) // 0001 1000 => 8


}