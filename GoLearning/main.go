package main
import "fmt"


func Test() {

	var x = 5
  
	fmt.Println(x <= 5 &&  x < 10) 
  }

func main(){

	fmt.Println("We start Go today our Objective is to become capabale of making api till the end of this week:)")


	fmt.Println("Operators:")

	num1 := 14
	num2 := 15
	sum := num1 + num2
	mod := num1%num2
	mult := num1 * num2
	num1++
	




	fmt.Println("sum of a num1 and sum2:",sum)
	fmt.Println("modulo of a num1 and sum2:",mod)
	fmt.Println("Multipilication of a num1 and sum2:",mult)
	fmt.Println("Increment of a num1 and sum2:",num1)

	Test()




fmt.Println("Conditions:")


num3 := 20

if num3 > 20 {
	fmt.Println("Number is greater than 20")
} else if num3 == 20 {
	fmt.Println("Number is 20:")

}else {
	fmt.Println("Number is less than 20")
}


fmt.Println("switch Case:")

num4 := 5

switch num4 {

	case 1:
	fmt.Print("Number is 1:")
	case 2:
	fmt.Println("Number is 2:")
	case 3:
	fmt.Println("Number is 3:")
	case 4:
	fmt.Println("Number is 4:")
	case 5:
	fmt.Println("Number is 5:")	
}


for i:= 0;i < 5;i++{
	fmt.Println(i)
}

arr := [5]int{1,2,3,4}
fmt.Println(arr)

for i := 0;i<5;i++{
	fmt.Println(arr[i])
}










    }



