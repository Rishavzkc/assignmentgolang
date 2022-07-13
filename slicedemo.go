package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 34, 67, 59, 90}

	fmt.Println("Original Slice : ", slice1)

	//Editing slice element
	slice1[3] = 89
	slice1[4] = 100

	fmt.Println("After editing original slice: ", slice1)

	//deleting element from slice
	slice1 = append(slice1[:3], slice1[4:]...)
	fmt.Println("After deleting element: ", slice1)

}
