package main

import "fmt"
type order struct {
	id int
	amount float64
}
// constructor function to create a new order
func newOrder(id int, amount float64) *order {
	return &order{
		id: id,
		amount: amount,
	}
}	
// receiver function to change the amount of an order
func (o *order) changeamount(newAmount float64){
	o.amount=newAmount
}
func main(){

	order := newOrder(1, 100.50)
	order.changeamount(150.75)
	fmt.Printf("Order ID: %d, Amount: %.2f\n", order.id, order.amount)

}
