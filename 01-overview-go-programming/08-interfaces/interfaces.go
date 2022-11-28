package main

import "fmt"

type Product interface {
	MadeIn() string
	Quantity() int
}

type Wallet struct {
	Name  string
	Price int
}

type Smartphone struct {
	Name             string
	Size             float64
	YearOfProduction int
}

func main() {
	wallet := Wallet{
		Name:  "RinosBlack",
		Price: 100000,
	}

	printInfoWallet(wallet)

	smartphone := Smartphone {
		Name: "Nevermind Phone",
		Size: 6.7,
		YearOfProduction: 2022,
	}

	printInfoSmartphone(smartphone)
}

func (w Wallet) MadeIn() string {
	return "Wallet made from rhino skin"
}

func (w Wallet) Quantity() int {
	return 11
}

func (s Smartphone) MadeIn() string {
	return "Smartphone made from recycle material"
}

func (s Smartphone) Quantity() int {
	return 5
}

func printInfoWallet(w Wallet) {
	fmt.Println("The product", w.MadeIn(), "and the quantity is", w.Quantity(), "#ImJustKidding it made from potato skin")
}

func printInfoSmartphone(s Smartphone) {
	fmt.Println("The product", s.MadeIn(), "and the quantity is", s.Quantity())
}