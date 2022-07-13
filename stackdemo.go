package main

import "fmt"

type Stack struct {
	price []int
}

func main() {

	mystack := Stack{}

	fmt.Println(mystack)
	mystack.push(100)
	mystack.push(200)
	mystack.push(300)
	fmt.Println(mystack)

	mystack.pop()
	fmt.Println(mystack)
}

//push
func (s *Stack) push(i int) {
	s.price = append(s.price, i)
}

//pop
func (s *Stack) pop() int {

	x := s.price[len(s.price)-1]

	s.price = s.price[:len(s.price)-1]
	return x
}
