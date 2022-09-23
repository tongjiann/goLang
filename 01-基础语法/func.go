package main

import "fmt"

func main() {
	//generalFunc()

	//anonymousFunc()

	//var a = []int{1, 2, 3}
	//print(a...)

	testDefer()
	fmt.Println("======")
	testDefer2()
}

func anonymousFunc() {
	var add = func(a, b int) int {
		return a + b
	}
	i := add(1, 2)
	fmt.Println(i)
}

func generalFunc() {
	i := add(1, 2)
	fmt.Println(i)
}

func add(a, b int) int {
	return a + b
}

func print(a ...int) {
	fmt.Println(a)
}

func testDefer() {
	for i := 0; i < 3; i++ {
		i := i
		defer func() { fmt.Println(i) }()
	}
}
func testDefer2() {
	for i := 0; i < 3; i++ {
		defer func() { fmt.Println(i) }()
	}
}
