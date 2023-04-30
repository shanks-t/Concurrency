package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const numberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= numberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recieved order number %d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("makeing pizza number %d will take %d seconds\n", pizzaNumber, delay)

		// delay
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("**** we ran out of ingredients for pizza #%d\n", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("****  the cook quit while make order for pizza #%d\n", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order %d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p

	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are trying to make
	var i = 0
	// run forever until we recieve quite message

	// try to make pizzas
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried to make a pizza (sent something to the data channel)
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}

	}
}

// Pizza makers and pizza consumers -- solution to producer consumer problem
func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("------------------------------------")

	// create a producer data structure
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in background as go routine
	go pizzeria(pizzaJob)
	// create and run consumer aka customer

	// print out ending message

}
