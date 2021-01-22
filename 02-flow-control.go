package main

import "fmt"

func main() {
	sum := 0 // Equivalent to `var sum int = 0`
	// initialization; condition; post-increment (all are optional)
	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Println("Whoa! We got to", sum)

	// Introducing the for loop with range clause
	var twoPowers = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, val := range twoPowers {
		if val > 50 {
			break
		}
		fmt.Printf("2**%d = %d\n", i, val)
	}
}
