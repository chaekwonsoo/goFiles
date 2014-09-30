package main

/*
import "fmt"

func square(x int) int {
	return x * x
}

func sum(a, b, c int) int {
	return a + b + c
}


func squaresum(a, b, c int) int {
	// **********************
	// We can do the parallel computing here! (a, b, and c is independent.)

	asq := square(a)
	bsq := square(b)
	csq := square(c)

	// **********************
	return sum(asq, bsq, csq)
}


func consquare(x int) chan int {
	conch := make(chan int)
	go func() {
		conch <- square(x)
	}()

	return conch
}

func consquaresum(a, b, c int) int {
	asynch, bsynch, csynch := consquare(a), consquare(b), consquare(c)
	as, bs, cs := <-asynch, <-bsynch, <-csynch

	return sum(as, bs, cs)
}


func main() {
	fmt.Println(consquaresum(1, 2, 3))
}
*/
