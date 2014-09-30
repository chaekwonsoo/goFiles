package main
/*
import "fmt"

func checkcond(b bool, c chan bool) chan bool {
	if !b {
		return nil
	}
	return c
}

func main() {	
	cond := 2
	
	cond0ch := make(chan bool)
	cond1ch := make(chan bool)
	cond2ch := make(chan bool)
	cond3ch := make(chan bool)

	go func() {
		cond0ch <- true
	}()

	go func() {
		cond1ch <- true
	}()

	go func() {
		cond2ch <- true
	}()

	go func() {
		cond3ch <- true
	}()

	select {
		case <- checkcond(cond == 0, cond0ch):
			fmt.Println("condition 0")
		case <- checkcond(cond == 1, cond1ch):
			fmt.Println("condition 1")
		case <- checkcond(cond == 2, cond2ch):
			fmt.Println("condition 2")
		case <- checkcond(cond == 3, cond3ch):
			fmt.Println("condition 3")
	}
}
*/
