package main
/*
import "fmt"

func main() {
	ch1 := make(chan chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	done := make(chan bool)

	g1 := func() {
		for {
			whichchan := <-ch1
			whichchan <- "This is "
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

	ch1 <- ch2
	<-done
	ch1 <- ch3
	<-done
}
*/
