//timeouts
//can be done using two ways
//u can use sleep or time.After
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//f1 waits on channels to have some values
func f1(wg *sync.WaitGroup, ch1, ch2 chan int) {
	defer wg.Done()
	for {
		select {
		case res1 := <-ch1:
			fmt.Println("--------------------Received on Channel 1", res1)
		case res2 := <-ch2:
			fmt.Println("--------------------Received on Channel 2", res2)
		case <-time.After(3 * time.Second):
			fmt.Println("Bye Bye")
			return
			//default:
			//	fmt.Println("Keeping it active")
		}
	}

}

func main() {
	fmt.Println("In Main Function")
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(3)

	//start a anon go routine that put some value on ch1 after 2 sec
	go func(wg *sync.WaitGroup) {
		//it wld sleep for 2 seconds and put some value
		defer wg.Done()
		time.Sleep(1 * time.Second)
		ch1 <- 2
	}(&wg)

	go func(wg *sync.WaitGroup) {
		//it wld sleep for 2 seconds and put some value
		defer wg.Done()
		time.Sleep(2 * time.Second)
		ch2 <- 2
	}(&wg)

	go f1(&wg, ch1, ch2)
	wg.Wait()
}
