package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go run on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}
	fmt.Printf("\n")
	for i := 0; i < 5; i++ {
		fmt.Printf("%s\n", runtime.GOOS)
	}
}
