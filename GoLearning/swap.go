package main

import "fmt"

func swapOfTONumbers(num1 int,num2 int){

	fmt.Println("Before Swaping num1:",num1)
	fmt.Println("Before Swaping num2:",num2)

	num1,num2 = num2,num1

	fmt.Println("After  Swaping num1:",num1)
	fmt.Println("After  Swaping num2:",num2)
	
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
	x, y = y, x+y
	}
	return x
	}




func main(){

	swapOfTONumbers(2,3)

	fmt.Println(fib(10))





}