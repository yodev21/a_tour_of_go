package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	// fmt.Println("My Favorite Number is", rand.Intn(10))
	// fmt.Printf("Now you have %g problems .\n", math.Sqrt(7))
	// 大文字で始まる名前は、外部のパッケージから参照できる
	// fmt.Println(math.Pi)
	fmt.Println(add(1, 4))
}
