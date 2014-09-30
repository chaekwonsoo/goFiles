package main

/*
import (
	"fmt"
)

func pass1(rch <-chan int, sch chan<- int) {
	if msg := <-rch; msg == 1 {
		sch <- msg + 1
	}
}

func pass2(rch <-chan int, sch chan<- int) {
	if msg := <-rch; msg == 2 {
		sch <- msg + 1
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go pass1(ch1, ch2)
	go pass2(ch2, ch3)

	ch1 <- 1
	fmt.Println(<-ch3)
}
*/
