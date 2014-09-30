package main

/*
import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	for {
		select {
			case c <- 0:
			case c <- 1:
		}
	}
}


func main() {
	go func() {
		(chan bool)(nil) <- true
	}()

	select {
		case <- (chan bool)(nil):
			fmt.Println("nil (bool)")
		case <- (chan int)(nil):
			fmt.Println("nil (int)")
		case <- (chan string)(nil):
			fmt.Println("nil (string)")
		default:
			fmt.Println("default")
	}
}
*/
