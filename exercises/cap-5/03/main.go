package main

import "fmt"

const a = 10
const b int = 20

func main() {

	fmt.Printf("%T\t%v\n", a, a)
	fmt.Printf("%T\t%v", b, b)
}
