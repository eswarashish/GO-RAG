package objects

import (
	"errors"
	"fmt"
)

type Vault interface{
	Deposit(amount float64)
	Withdraw(amount float64) error
}

type Account struct{
	Owner string
	balance float64
}

func (a Account) GetBalance() float64{
	return a.balance
}

type PersonalAccount struct{
	Account
}

func (p *PersonalAccount) Deposit(amount float64){
	p.balance += amount
}

func (p *PersonalAccount) Withdraw(amount float64) error{
	if ((p.balance - amount) >= 0){
		p.balance -= amount
		return nil
	}
	return errors.New("Personal Accounts cannot go below 0")
	
}

func PerfomTransaction(v Vault, amount float64, isDeposit bool){
	if isDeposit{
		v.Deposit(amount)
	}else{ err := v.Withdraw(amount)
	if err != nil{
		fmt.Println("error occured", err)
	}}
}



