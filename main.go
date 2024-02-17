package main

import "fmt"

func ttt(i ...int) {
	fmt.Println(i)
}

func main() {
	ttt(1)
	ttt([]int{2, 3, 4, 5}...)
}
