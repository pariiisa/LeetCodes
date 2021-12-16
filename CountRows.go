package main

import (
	"fmt"
)

func main() {

	fmt.Println(getRows(5))
}

func getRows(n int) int {

	var i int = 1

	if n < 0 {
		fmt.Println("The number should br greater than 0")
	} else {

		for i*(i+1)/2 <= n {
			i++
		}

	}
	return i - 1

}
