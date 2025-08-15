package main

import (
	"fmt"
)

func main() {
	introOne := "welcome ojas to xist"
	introTwo := "welcome sanjay to xist"
	introThree := "welcome vinu to xist"

	finalize := []string{introOne, introTwo, introThree}

	order()
	totaluser(finalize)
}

func totaluser(finalize []string) {
	for index, final := range finalize {
		fmt.Printf("%d. %s \n", index+1, final)
	}
}

func order() {
	var cName = "JACK DORSEY"
	var items1 = "samosa"
	var item2 = "biryani"
	var item3 = "chicken curry"
	var item4 = "chocolate cake"
	var item5 = "ice cream"
	var contact = "xistapp@gmail.com"
	fmt.Println("WELCOME TO", cName)
	fmt.Println("Please enter your name")

	var userName string
	var itemsneed string
	var paymentOption string
	var address string

	fmt.Scan(&userName)
	// fmt.Printf("Hello %v, please select the item you want to order \n", userName)
	fmt.Printf("Hello %v, Please select an item from the menu %v %v %v %v %v for more details contact %v\n", userName, items1, item2, item3, item4, item5, contact)

	fmt.Scan(&itemsneed)
	fmt.Printf("You have selected %v, please select your payment option\n", itemsneed)

	fmt.Scan(&paymentOption)
	fmt.Printf("You have selected %v as your payment option, please enter your address\n", paymentOption)

	fmt.Scan(&address)
	fmt.Printf("Your order for %v has been placed successfully. It will be delivered to %v. Thank you for choosing %v!\n", itemsneed, address, cName)
	fmt.Println("For any queries, please contact us at", contact)

}
