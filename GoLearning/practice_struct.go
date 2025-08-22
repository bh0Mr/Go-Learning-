package main

import "fmt"


type Bank struct{
	accounts []BankAccount
}

type Atm struct{
	currentAmount uint32
	accounts []BankAccount
}


type FreeFire struct{
	userAccounts []Account
	skins []Skin
	weapons []Weapon
}


type LockerStore struct{

	lockers map[string]Locker // key=ownerName value=Locker

}



type GameAccount struct{

	Money uint32
	Diamond uint32

	Skills []Skill

	Skins []Skin

	Level uint32
	Username string
	Email string
	AccountId string
}

type Skill struct{

	Name string

	AttributeType string


}


type Skin struct{

	RateOfFire uint32
	Range uint32
	Damage uint32

}

type BankAccount struct {

	name string 
	balance float32
	branch string
	accountType string 


}


func main(){

	account := BankAccount { 
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
