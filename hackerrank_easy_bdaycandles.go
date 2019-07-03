//hacker rank birthday cake candles
package main

import "fmt"

func candleBlowOutCount(candles []int,n int )int{
	max := 0
	count := 0
	for _, v := range candles{
		if v >max {
			max = v
		}
	}
	//count how many times max exist
	for _, v :=range candles{
		if v == max{
			count +=1
		}
	}
	return count
}


func main(){
	fmt.Println("In Main Function")
	candles := []int{3,2,1,3,4}
	candlesblownout := candleBlowOutCount(candles,5)
	fmt.Println("No Of candles ",candlesblownout)
}
