package main

import "fmt"

func main() {

	x := []int{1, 2, 3}
	y := []int{6, 7, 8, 9}

	fmt.Println(x, y)

	x = append(x, 4, 5)

	fmt.Println(x, y)

	x = append(x, y...)

	fmt.Println(x)
	x = append(x[:8], x[9:]...)

	fmt.Println(x)

}
