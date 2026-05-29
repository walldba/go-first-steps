package main

import "fmt"

const (
	_ = 1998 + iota
	b
	c
)

func main() {

	fmt.Println(b, c)

	// iota é um contador que começa em 0 e incrementa a cada linha
}
