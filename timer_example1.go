//timer example
//create two timer
//one executes fully
//one we have to stop
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In Main Function")
	timer1 := time.NewTimer(3 * time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	select {
	case <-timer1.C:
		fmt.Println("Timer 1 Ticked")
		timer2.Stop()
		return
	case <-timer2.C:
		fmt.Println("Timer 2 ticked")
		timer1.Stop()
		return
	case <-time.After(5 * time.Second):
		fmt.Println("bye bye! No Timer Ticked")
		return
	}
}
