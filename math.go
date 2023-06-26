package main

import "fmt"

func main() {
	fmt.Println(sum(112, 10))
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}

func sumX(a int, b int) int {
	return a + b + a
}

func sumY(a int, b int) int {
	return a + b + b
}

func div(a int, b int) int {
	return a / b
}

func avg(a int, b int, c int) int {
	return (a + b + c) / 3
}