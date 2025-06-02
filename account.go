package main

import (
	"fmt"
	"strings"
)

type transaction struct {
	date       string
	trans_type string
	amnt       float64
	source     string
	dest       string
	bal        float64
	success    bool
}

type account struct {
	name           string
	phone_num      string
	acct_num       string
	acct_type      string
	bal            float64
	limit          float64
	trans_hist     map[uint32]transaction
	__trans_size__ uint32
}

// Create account

func create_acct(name string, phone_num string, acct_type string) account {
	acct := account{
		name:           strings.TrimSpace(name),
		phone_num:      phone_num,
		acct_num:       phone_num[len(phone_num)-10:],
		acct_type:      acct_type,
		bal:            0.00,
		limit:          0.00,
		trans_hist:     map[uint32]transaction{},
		__trans_size__: 0,
	}

	return acct
}

// View account

func (a *account) print_acct() {
	fmt.Printf("Account Name: %v \n", a.name)
	fmt.Printf("Phone Number: %v \n", a.phone_num)
	fmt.Printf("Account Number: %v \n", a.acct_num)
	fmt.Printf("Account Type: %v \n", a.acct_type)
	fmt.Printf("Account Balance: %v \n", a.bal)
}

// Deposit

func (a *account) deposit(amnt float64) {
	// last_key
	a.__trans_size__ += 1
	a.bal += amnt
	a.trans_hist[a.__trans_size__] = transaction{
		date:       "dd/mm/yyyy",
		trans_type: "credit",
		amnt:       amnt,
		source:     "",
		dest:       a.acct_num,
		bal:        a.bal,
		success:    true,
	}

	fmt.Printf("Dear %v,\n a sum of %v has been deposited into your account %v. Your updated balance is %v", a.name, amnt, a.acct_num, a.bal)
}

// Withdraw

func (a *account) withdraw(amnt float64) {
	a.__trans_size__ += 1
	a.bal -= amnt
	a.trans_hist[a.__trans_size__] = transaction{
		date:       "dd/mm/yyyy",
		trans_type: "debit",
		amnt:       amnt,
		source:     a.acct_num,
		dest:       "",
		success:    true,
	}
}

// Get account balance

func (a *account) check_bal() {
	fmt.Printf("Your account balance is: %v \n", a.bal)
}

// Check transaction history

func (a *account) get_trans_hist() {
	fmt.Println("Your account history is:")
	printTrans(a.trans_hist)
	// fmt.Printf("Type of trans_hist is: %T", a.trans_hist)
}

// Transfer to another account

func (a *account) transfer(acct *account, amnt float64) {
	if a.bal >= amnt {
		a.__trans_size__ += 1
		a.bal -= amnt
		a.trans_hist[a.__trans_size__] = transaction{
			date:       "dd/mm/yyyy",
			trans_type: "debit",
			amnt:       amnt,
			source:     a.acct_num,
			dest:       acct.acct_num,
			bal:        a.bal,
			success:    true,
		}

		// update reciepient history & balance
		acct.bal += amnt
		acct.__trans_size__ += 1
		acct.trans_hist[acct.__trans_size__] = transaction{
			date:       "dd/mm/yyyy",
			trans_type: "credit",
			amnt:       amnt,
			source:     a.acct_num,
			dest:       acct.acct_num,
			bal:        acct.bal,
			success:    true,
		}
	} else {
		fmt.Println("In sufficient funds!!!")
	}
}

// Print account statement

func (a *account) acct_statement() {
	fmt.Printf("Name: %-125s\nPhone Number: %-125s\n Account Number: %20v\t Account Type: %20v\t Limit:%20.2f\n", a.name, a.phone_num, a.acct_num, a.acct_type, a.limit)
	fmt.Println(strings.Repeat("-", 125))
	printTrans(a.trans_hist)
	fmt.Println(strings.Repeat("=", 125))

}
