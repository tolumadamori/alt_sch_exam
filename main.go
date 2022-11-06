package main

import "fmt"

/*-------------------------- Car Struct -----------------------------*/
//cars have brands, models, class, status and year of manufacture fields.
type car struct {
	brand  string
	model  string
	class  string
	status string
	year   int
}

/* -------- car attributes --------------*/
//The main attribute of a car is that it can move from point a to b.
//This functions moves the car by the distance provided.
func (c car) move(distance int) {
	if distance > 0 {
		for i := 0; i < distance; i++ {
			fmt.Println("car is moving...")

		}
		fmt.Printf("\n")
		fmt.Println("car has stopped")
		fmt.Printf("\n")

	} else {
		fmt.Printf("\nThe distance provided: %v  is too short. Pass a distance greater than zero", distance)
	}
}

/*------------ Product Struct ---------------------*/
// Cars and other objects can be found anywhere and can belong to anyone.
//However, for a product to belong in our store, it needs to be in our inventory with an assigned inventory number and other product details(class, price etc).
type product struct {
	invNumber     int
	productClass  string //a product can be a car, an accessory e.g car chargers dashcams etc or a consumable e.g tires, engine oil etc.
	productPrice  float32
	productObject // a product object is any object that implements the productObject interface declared below.
}

//Here we declare the productobject interface.
//We are declaring  this interface so we can embed items with these attributes into our product struct declared above.
//Objects just need to have the display product method to implement our productObject interface
type productObject interface {
	displayProduct() //displays the product
}

//This method prints out the details of the car and cars now implement the productobject interface.
func (c car) displayProduct() {
	fmt.Printf("\nThese are the car details: \n Brand : %v \n Model: %v \n Status: %v \n Year: %v \n\n", c.brand, c.model, c.status, c.year)
}

//Products also have attributes. We declare the product attributes below.

//This function displays the product
func (p product) showProduct() {
	fmt.Printf("\n Displaying the product: %v \n", p)
}

//Displays the product price
func (p product) displayPrice() {
	fmt.Printf("\n\n The price of the product is: %f \n \n", p.productPrice)
}

/*------------------------- Store ---------------------------------*/

//Car shops need warehouses and show rooms to store products.
//Here we declare the store type.
type store []product

// A store is just a room if we can't add items into the room, display(list) our items, sell our items that we've stored and see a history of transactions on our stored items.
//Let's make our stores world class showrooms and define their attributes.

//This funtion adds products to the store. We should only add inventory items to our store.
func (s *store) addToStore(i ...product) {
	*s = append(*s, i...)
	fmt.Println("The products were added to the store successfully")
}

//This functions lists out all the items currently in the store.
func (s store) display() {
	fmt.Printf("\n \n")
	fmt.Printf("\n Displaying items in our store \n")
	for _, v := range s {
		fmt.Printf("\n *** %v", v)
	}
	fmt.Printf("\n \n \n")
}

//This function sells the item from the store. It takes the inventory number, removes the product from our inventory and adds it to our transaction history
func (s *store) sellItem(itemNumber int, transactionHistory *[]product) {
	for i, v := range *s {
		if v.invNumber == itemNumber { //we need to check that the product is indeed in the store.
			fmt.Printf("The %v with inventory number: %v  is now sold", v.productClass, v.invNumber) //If it is in the store, we need to send a notification that it is now sold.
			*s = append((*s)[:i], (*s)[i+1:]...)                                                     //Then we remove it from our store inventory.
			*transactionHistory = append(*transactionHistory, v)                                     //and we add it to our transaction history.

		}
	}
}

//This function shows a transaction history and the total products of items sold.
func (s store) history(transactionHistory []product) {
	var revenue float32
	for _, v := range transactionHistory {
		revenue += v.productPrice
	}
	fmt.Printf("The number of cars sold from the store is: %v \n", len(transactionHistory))
	fmt.Printf("\nTotal revenue from sold products is: %v\n", revenue)
	fmt.Println("Please see details of sold proucts:")
	fmt.Printf("\n %v", transactionHistory)
}

//Creating our store.
var showRoom store

func main() {
	//creating new car objects.
	car1 := car{
		"Toyota",
		"Camry",
		"Sedan",
		"Foreign Used",
		2010,
	}

	car2 := car{
		"Kia",
		"Rio",
		"Compact",
		"Brand New",
		2018,
	}

	car3 := car{
		"Honda",
		"Cross Tour",
		"Cross Over",
		"Foreign Used",
		2010,
	}

	car4 := car{
		"Lexus",
		"RX 350",
		"SUV",
		"Nigerian Used",
		2010,
	}

	//Let's move our car.
	car1.move(5)

	//Let's display one of our cars.
	car2.displayProduct()

	//Adding our cars with their product details as products in our store.
	item1 := product{
		1000001,
		"car",
		6000000,
		car1,
	}

	item2 := product{
		1000002,
		"car",
		7000000,
		car2,
	}

	item3 := product{
		1000003,
		"car",
		5000000,
		car3,
	}

	item4 := product{
		1000004,
		"car",
		6500000,
		car4,
	}

	//Let's display one of our items
	item3.showProduct()

	//Let's display one of our product prices
	item4.displayPrice()

	//Adding items to our store
	showRoom.addToStore(item1, item2, item3, item4)

	//Displaying the items in our store.
	showRoom.display()

	//Let's sell an item from our store.

	//First we create our transaction history.
	records := make([]product, 0)

	//Now we sell.
	fmt.Println("selling...")
	showRoom.sellItem(1000004, &records)

	//selling a second item
	showRoom.sellItem(1000003, &records)

	//Displaying our items again:
	showRoom.display()

	//Let's see our transaction history
	showRoom.history(records)

}
