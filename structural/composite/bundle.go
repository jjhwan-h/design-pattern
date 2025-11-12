package main

type Bundle struct {
	items []Component
}

func (b *Bundle) Price() int {
	var sum int
	for _, item := range b.items {
		sum += item.Price()
	}
	return sum
}

func (b *Bundle) Add(item Component) {
	b.items = append(b.items, item)
}
