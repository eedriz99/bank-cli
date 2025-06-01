package main

import "fmt"

func main() {

	user1 := create_acct("First User", "07123456789", "Current")

	user1.deposit(1000.99)
	user1.print_acct()

	user2 := create_acct("Second User", "07987654321", "Savings")
	user2.print_acct()

	fmt.Println("<============================= POST TRANSFER =================================>")

	user1.transfer(&user2, 350.90)

	user1.check_bal()
	user2.check_bal()

	user1.get_trans_hist()
	println("\n \n")
	user2.get_trans_hist()

	// fmt.Println(user1.trans_hist)
}
