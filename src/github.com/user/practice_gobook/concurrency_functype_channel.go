package main

//import "fmt"

// ***********************************************
type Car struct {
	model string
	price float64
	funch chan func()
}

func (c *Car) sequence() {
	for f := range c.funch {
		f()
	}
}

func (c *Car) setPrice(price float64) {
	c.funch <- func() {c.price = price }
}

func (c *Car) getPrice() float64 {
	tmpch := make(chan float64)
	c.funch <- func() { tmpch <- c.price }

	return <-tmpch
}
// ***********************************************

func NewCar(model string, price float64) *Car {
	car := &Car{model, price, make(chan func())}
	go car.sequence()

	return car
}

/*
func main() {
	newcar := NewCar("Mustang", 40000.5)
	fmt.Println(newcar.getPrice())
	newcar.setPrice(25000.99)
	fmt.Println("On Sale:", newcar.getPrice())
}
*/
