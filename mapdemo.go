package main

import "fmt"

func main() {

	names := make(map[string]int)

	names["Ram"] = 1
	names["John"] = 3
	names["jack"] = 4

	fmt.Println("Printing original map")
	fmt.Println(names)

	delete(names, "Ram") //delete one element

	fmt.Println("After Deleting one element ")
	for i, value := range names {
		fmt.Println("Names: ", i, "Id: ", value)
	}

}
