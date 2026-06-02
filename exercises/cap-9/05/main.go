package main

import "fmt"

func main() {
	x := make([]string, 5)
	fmt.Println(x)

	x = []string{"Go", "Python", "Java", "Ruby", "C++"}
	fmt.Println(x)

	x = append(x, "JavaScript")
	fmt.Println(x)

	for y := 0; y < len(x); y++ {
		fmt.Println(x[y])
	}

}
