package main

import "fmt"

type bankAccount struct {

	name string 
	balance float32
	branch string
	accountType string 

}
func main(){

	account := bankAccount { 
		name: "Tin",
		balance: 12000.3,
		branch: "jammu",
		accountType:"Saving",
	}

	fmt.Println("Name of the Account holder:",account.name)
	fmt.Println("Balance of the Account:",account.balance)
	fmt.Println("Branch of the Account :",account.branch)
	fmt.Println("Type of the Account:",account.accountType)
	

}
