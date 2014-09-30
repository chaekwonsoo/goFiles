package main

/*
import "fmt"

func main() {
	ch := make(chan int, 10)

	go func() {
		x, y := 0, 1
		for i := 0; i < cap(ch); i++ {
			ch <- x
			x, y = y, x+y
		}
		close(ch)
	}

	for {
		v, okay :=<- ch
		if !okay {
			break;
		}
		fmt.Println(v)
	}
}
*/
