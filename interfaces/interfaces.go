package main

import "fmt"




type razorpay struct{}

func (r razorpay) pay(amount int) {
	fmt.Printf("Paying %d using Razorpay\n", amount)
}

type paypal struct{}

func (p paypal) pay(amount int) {
	fmt.Printf("Paying %d using PayPal\n", amount)
}

type paymentGateway interface {
	pay(amount int)
}
func main(){
  razorpay := razorpay{}
  paypal := paypal{}

  newPayment:=paymentGateway(razorpay)
  newPayment.pay(100)
  
  newPayment=paymentGateway(paypal)
  newPayment.pay(200)

}
