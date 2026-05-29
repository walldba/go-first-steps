package main

import "fmt"

type Wall int

var x Wall

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)

}
