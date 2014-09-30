package main

/*
import(
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "every 1 second"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		c2 <- "every 2 second"
		time.Sleep(time.Second * 2)
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			default:
				fmt.Println("Default")
				time.Sleep(time.Second * 1)
			}
		}
	}()

	time.Sleep(time.Second * 5)
}
*/
