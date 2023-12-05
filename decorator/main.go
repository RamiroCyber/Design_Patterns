package main

import (
	"fmt"
)

type Coffee interface {
	Cost() float64
}

type SimpleCoffee struct{}

func (c SimpleCoffee) Cost() float64 {
	return 10
}

type CoffeeDecorator struct {
	Coffee Coffee
}

func (d CoffeeDecorator) Cost() float64 {
	return d.Coffee.Cost()
}

type MilkDecorator struct {
	CoffeeDecorator
}

func (m MilkDecorator) Cost() float64 {
	return m.Coffee.Cost() + 2
}

type SugarDecorator struct {
	CoffeeDecorator
}

func (s SugarDecorator) Cost() float64 {
	return s.Coffee.Cost() + 1
}

func main() {
	simpleCoffee := SimpleCoffee{}
	milkCoffee := MilkDecorator{CoffeeDecorator{simpleCoffee}}
	sugarMilkCoffee := SugarDecorator{CoffeeDecorator{milkCoffee}}

	fmt.Printf("Cost of Simple Coffee: $%.2f\n", simpleCoffee.Cost())
	fmt.Printf("Cost of Milk Coffee: $%.2f\n", milkCoffee.Cost())
	fmt.Printf("Cost of Sugar Milk Coffee: $%.2f\n", sugarMilkCoffee.Cost())
}
