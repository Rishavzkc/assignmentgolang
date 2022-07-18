package main

import "fmt"

type Employee struct {
	EmplId     int
	EmplName   string
	EmplSalary int
}

func main() {
	Employees := make(map[int]Employee)

	var input int

	for {
		fmt.Println("Enter the Employee Details")
		fmt.Println("1 For Create")
		fmt.Println("2 For Read")
		fmt.Println("3 For Update")
		fmt.Println("4 For Delete")
		fmt.Println("5 For Exit")

		fmt.Scanln(&input)

		switch input {
		case 1:
			createEmployee(Employees)
		case 2:
			readEmployee(Employees)
		case 3:
			updateEmployee(Employees)
		case 4:
			deleteEmployee(Employees)
		default:
			fmt.Println("Exit")
		}

	}

}

func createEmployee(employees map[int]Employee) {
	emp := Employee{}
	fmt.Println("Enter employee id")
	fmt.Scanln(&emp.EmplId)

	fmt.Println("Enter Employee Name")
	fmt.Scanln(&emp.EmplName)

	fmt.Println("Enter Employee Salary")
	fmt.Scanln(&emp.EmplSalary)

	employees[emp.EmplId] = emp
}

func readEmployee(employees map[int]Employee) {
	fmt.Printf("%#v \n", employees)
}

func updateEmployee(employees map[int]Employee) {
	var EmplId, EmplSal int
	fmt.Println("Enter Employee id for Salary update")
	fmt.Scanln(&EmplId)

	fmt.Println("Enter Updated salary")
	fmt.Scanln(&EmplSal)

	updatedEmp, _ := employees[EmplId]

	updatedEmp.EmplSalary = EmplSal
	employees[EmplId] = updatedEmp

}

func deleteEmployee(employee map[int]Employee) {
	var id int
	fmt.Println("Enter the Employee id to delete")
	fmt.Scanln(&id)
	delete(employee, id)
}
