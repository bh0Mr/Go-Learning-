package main
import "fmt"



func sumOfThreeNumbers(n1 *int ,n2 *int ,n3 *int ) int {

	sum := *n1 + *n2 + *n3

	return sum
}


func main(){
	num1 := 3
	num2 := 4
	num3 := 6

	fmt.Println(sumOfThreeNumbers(&num1,&num2,&num3))



}