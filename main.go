package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// helper function to get input from the user
// we want to pass in a prompt and a reader, which is a pointer to a bufio.Reader
// we want to return a string and an error
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

// create new bill using the newBill function
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - Add item, s - Save bill, t - Add tip): ", reader)

	switch opt {
	case "a":
		// we want to ask for item name and price
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		// parse the price as a float
		p, err := strconv.ParseFloat(price, 64)

		// if error is not nil, there is an error
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(name, p)
		fmt.Println("Item added -", name, price)
		promptOptions(b)

	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip added:", tip)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("You saved the file -", b.name)

	default:
		fmt.Println("That is not a valid option")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}
