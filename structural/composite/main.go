package main

import "fmt"

func main() {
	item1 := Item{price: 1}
	item2 := Item{price: 2}
	item3 := Item{price: 3}
	item4 := Item{price: 4}

	Bundle := Bundle{}
	Bundle.Add(&item2)
	Bundle.Add(&item3)
	Bundle.Add(&item4)

	fmt.Println(Checkout(&item1))
	fmt.Println(Checkout(&Bundle))
}

func Checkout(c Component) int {
	return c.Price()
}
