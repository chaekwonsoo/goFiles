package main

/*
import "fmt"

func main() {

	type T struct {
		fn func() string
		ch chan string
	}

	ch1 := make(chan T)
	ch2 := make(chan string)
	ch3 := make(chan string)
	done := make(chan bool)

	g1 := func() {
		for {
			t := <-ch1
			t.ch <- t.fn()
		}
	}
	go g1()

	g2 := func() {
		fmt.Println(<-ch2, "g2")
		done <-true
	}
	go g2()

	g3 := func() {
		fmt.Println(<-ch3, "g3")
		done <-true
	}
	go g3()

	ch1 <- T{fn: func() string {return "hello"}, ch: ch2}
	<-done
	
	ch1 <- T{fn: func() string {return "hello"}, ch: ch3}
	<-done
}
*/
