package main

import "fmt"

// custom type
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// function that takes a string and returns a bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// receiver function associated with the bill type - format the bill
// limiting this function to the bill object
func (b *bill) format() string {
	// cycle through the items in the bill
	fs := "Bill breakdown: \n" // fs stands for formatted string
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// add the tip
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "tip:", b.tip)

	// add the total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)
	return fs
}

// update tip - pass in the pointer to the bill
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
