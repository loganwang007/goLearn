package main

import "fmt"

func main() {
	intRange := makeRange(0, 10)

	for _, value := range intRange {
		if value%2 != 0 {
			fmt.Println(value, "is odd")
		} else {
			fmt.Println(value, "is even")
		}
	}

}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
