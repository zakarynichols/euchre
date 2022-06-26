package main

import "fmt"

func main() {
	sum := add(10, 20)
	fmt.Print(sum)
}

func add(n1 int, n2 int) int {
	return n1 + n2
}
