package main

import "fmt"

func main() {
	

	Append_prac()


}

func Append_prac(){

	var queue []string
	
	
	queue = append(queue, "First Person \n")
	queue = append(queue, "Second Person \n")
	queue = append(queue, "Third Person\n")
	

	fmt.Println("Orders in the queue:\n", queue)

}



