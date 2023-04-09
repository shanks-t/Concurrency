package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// var for bank balance
	var bankBalance int
	var balance sync.Mutex
	// print out starting values
	fmt.Printf("initial account balance %d.00", bankBalance)
	fmt.Println()

	// define weekly revenue
	incomes := []Income{
		{Source: "Investments", Amount: 500},
		{Source: "Infor", Amount: 2500},
		{Source: "Theft", Amount: 300},
		{Source: "Youtube", Amount: 5500},
	}

	wg.Add(len(incomes))
	// loop through 52 week and calc how much is made
	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("on week %d you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}
	wg.Wait()

	// print final balance
	fmt.Printf("final balance for the year is $%d.00", bankBalance)
	fmt.Println()
}
