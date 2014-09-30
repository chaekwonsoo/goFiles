package main

/*
import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	done := make(chan bool)

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		done <- true
	}()
	<- done
}*/
