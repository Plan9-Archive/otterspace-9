package main

import "fmt"

func Example212() {
	fmt.Println(ex212([]int{1, 4, 2, 3}))
	// Output:
	// [4 3 2 1]
}
 
func Example213_true() {
	fmt.Println(ex213([]int{1, 2, 3, 4}, 3))
	// Output:
	// 2
}

func Example213_false() {
	fmt.Println(ex213([]int{1, 2, 3, 4}, 9))
	// Output:
	// -1
}

func Example214() {
	fmt.Println(ex214([]int{0, 1, 1, 0}, []int{1, 1, 0, 1}))
	// Output:
	// [1 0 0 0 1]
}
