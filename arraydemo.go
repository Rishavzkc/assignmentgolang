package main

import "fmt"

func main() {

	arr := [6]int{13, 23, 34, 45, 87, 89}

	//printing array
	fmt.Println("Original Array :", arr)

	//After editing
	arr[2] = 100

	//printing Array
	fmt.Println("After Editing original Array: ", arr)

	//index to be remove
	i := 3

	if i < 0 || i >= len(arr) {
		fmt.Println("The given index is out of bound")
	} else {
		arr2 := 0

		for index := range arr {
			if i != index {
				arr[arr2] = arr[index]
				arr2++
			}
		}
		newarr := arr[:arr2]
		fmt.Println("THe Array is :", newarr)

	}

}
