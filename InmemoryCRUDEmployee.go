package main

import "fmt"

type Employee struct {
	Name    string
	Age     int
	EmpId   int
	Address Address
}

type Address struct {
	City string
	Pin  int
}

func CrudEmployee() {
	emp1 := Employee{
		Name:  "John",
		Age:   28,
		EmpId: 129013,
		Address: Address{
			City: "Pune",
			Pin:  405906,
		},
	}
	emp2 := Employee{
		Name:  "Andrew",
		Age:   35,
		EmpId: 127809,
		Address: Address{
			City: "Mumbai",
			Pin:  498707,
		},
	}
	fmt.Println("Creating Employee Details")
	fmt.Println(fmt.Sprintf("Employee1 is: %+v\n ", emp1))
	fmt.Println(fmt.Sprintf("Employee2 is: %+v\n ", emp2))

	//Read Employee details
	fmt.Println("Reading name and employeeId of Employees")
	fmt.Println("Employee 1 : ", emp1.Name, "Employee Id: ", emp1.EmpId)
	fmt.Println("Employee 2 : ", emp2.Name, "Employee Id: ", emp2.EmpId)

	//Updating Employee details

	emp1.Age = 30
	emp1.Address.City = "Nagpur"

	emp2.EmpId = 134576
	emp2.Address.Pin = 409876

	fmt.Println("Updating Employee Details")
	fmt.Println(fmt.Sprintf("Employee1 is: %+v\n ", emp1))
	fmt.Println(fmt.Sprintf("Employee2 is: %+v\n ", emp2))

	//Deleting Employee detail
	emp1 = Employee{Name: "John", Age: 28}

	fmt.Println("After Deleting Employee 1 Details")
	fmt.Println(fmt.Sprintf("Employee1 is: %+v\n ", emp1))

}

func main() {
	CrudEmployee()
}
