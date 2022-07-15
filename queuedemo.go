package main

import "fmt"

type queue struct {
	names []string
}

func main() {
	myqueue := queue{}
	fmt.Println(myqueue)

	myqueue.addelement("John")
	myqueue.addelement("Jack")
	myqueue.addelement("Andrew")

	fmt.Println(myqueue)

	myqueue.removeelement()
	fmt.Println(myqueue)
}

//Add element
func (q *queue) addelement(i string) {
	q.names = append(q.names, i)
}

//Remove element
func (q *queue) removeelement() string {
	x := q.names[0]       //to remove 1st index value
	q.names = q.names[1:] //here we are skipping index 0 and start from 1
	return x
}
