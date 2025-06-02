package main

import (
	"fmt"
	"strings"
)

func MaxKey(m map[int]string) int {
	var c_max int
	for k, _ := range m {
		if k >= *&c_max {
			*&c_max = k
		}
	}

	return c_max
}

func printTrans(m map[uint32]transaction) {
	fmt.Printf("%5s | %-10s | %-10s | %20s | %-15s | %-15s | %20s | %-10s\n", "S/N", "Date", "Type", "Amount(£)", "From", "To", "Balance(£)", "Status")
	fmt.Println(strings.Repeat("-", 125))
	for k, v := range m {
		fmt.Printf("%5d | %-10s | %-10s | %20.2f | %-15s | %-15s | %20.2f | %-15v\n", k, v.date, v.trans_type, v.amnt, v.source, v.dest, v.bal, v.success)
	}
}

// S/N(5) date(-10) type(-10) amnt(	20, 2dp)
// source(-15) destination(-15) balance(20, 2dp) status(-10)
