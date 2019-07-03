//example of a closure
package main

import "fmt"

func c1(n int) func(int) int {
	a := n
	//complete def of function is added here
	//they are anonymous function
	return func(x int) int {
		return x + a
	}
}

func main() {
	fmt.Println("I am in the main function")
	f1 := c1(5)
	f2 := c1(8)
	for i := 0; i < 2; i++ {
		fmt.Printf("\nf1 value for %d is %d", i, f1(i))
		fmt.Printf("\nf2 value for %d is %d", i, f2(i))
	}
}
