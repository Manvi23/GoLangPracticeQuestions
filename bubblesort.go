//bubble sort
package main

import "fmt"

//no reference array pass, as we are using slice
//if we modify slice , it will get modified
//last element wld be biggest
//bubble will go till last first
func bubblesort(s []int) {

	//not adjascent are not compared
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				fmt.Printf("Swapping :%d %d\n", s[j], s[j+1])
				temp := s[j]
				s[j] = s[i]
				s[i] = temp
			}
		}
	}
}

func printSlice(s []int) {
	for _, v := range s {
		fmt.Printf(" %d", v)
	}
	fmt.Println("end")
}

func main() {
	fmt.Println("In main function")
	s := []int{44, 11, 88, 33, 99, 66, 77, 55, 22}
	printSlice(s)
	bubblesort(s)
	printSlice(s)
}
