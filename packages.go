package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// fmt.Println("My Favorite Number is", rand.Intn(10))
	// fmt.Printf("Now you have %g problems .\n", math.Sqrt(7))
	// 大文字で始まる名前は、外部のパッケージから参照できる
	// fmt.Println(math.Pi)
	// fmt.Println(add(1, 4))
	fmt.Println(split(17))
}
