package main

import "fmt"

func main() {
	x := map[string]int{
		"Wallace": 36,
		"Leticia": 33,
		"Enzo":    31,
		"Arthur":  22,
		"Julia":   18,
	}

	for i, v := range x {
		fmt.Println(i, v)
	}

	delete(x, "Leticia")
	fmt.Println(x)

}
