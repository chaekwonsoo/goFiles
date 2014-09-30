package main

/*
import "fmt"

func main() {
	done := make(chan bool)

	intslice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	
	for _, v := range intslice {
		v := v
		go func() {
			fmt.Print(v, " ")
			done <- true
		}()
	}
	
	for _, v := range intslice {
		go func(u int) {
			fmt.Print(u, " ")
			done <- true
		}(v)
	}

	for _, v := range intslice {
		go func() {
			fmt.Print(v, " ")
			done <- true
		}()
	}

	for i := 0; ; i++ {
		if i >= len(intslice) {
			break
		}
		<- done
	}
}
*/
