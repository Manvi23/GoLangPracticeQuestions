package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("I am in Main Function")
	fmt.Println("No Of Cores in the System :%d", runtime.NumCPU())

}
