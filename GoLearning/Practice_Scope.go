package main
import "fmt"


var BankName = "Real_bank"


func main(){


	var accountBalance = 50000

	fmt.Println("Account Holder Enter Your name:")

	var name string
	fmt.Scanln(&name)


	fmt.Println("Welcome Your account balance is:",name,accountBalance)


    fmt.Println("Enter Amount to Withdraw")
	var debitMoney int
	fmt.Scanln(&debitMoney)


	fmt.Println("Enter Your pin")
	var pin int
	fmt.Scanln(&pin)

	if pin == 0047 {

		accountBalance -= debitMoney

		fmt.Println("Your money has been withdrawn")
		fmt.Println("Your remaning balance is:",accountBalance)

	}

	

	



	











}