//waitgroup usage
package main

import (
	"fmt"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("In F1 and i is %d\n", i)
	}
}
func f2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("In F2 and i is %d\n", i)
	}
}

func main() {
	fmt.Println("In Main Function")
	var wg sync.WaitGroup
	wg.Add(2)
	go f1(&wg)
	go f2(&wg)
	wg.Wait()
}
