package main

import "fmt"

//while文はないがfor 式で実現ができる
func main() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}
