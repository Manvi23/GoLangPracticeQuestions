//hackerrank staircase
package main

import "fmt"

func PrintStairCase(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for k := n - i - 1; k < n; k++ {
			fmt.Print("#")
		}
		fmt.Println("")
	}
}

func main() {
	fmt.Println("In Main Function")
	PrintStairCase(3)
}
