package main

import "fmt"

func main() {
	numbersSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range numbersSlice {
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}
