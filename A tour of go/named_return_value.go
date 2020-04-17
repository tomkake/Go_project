package main

import "fmt"

//下記のように返せるけれどもコードの読みやすさが欠如するためやめたほうが良い。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
