package main

type Item struct {
	price int
}

func (i *Item) Price() int {
	return i.price
}
