package main

/*
import "fmt"

func main() {
	x := 100
	fns := []func() {		// slice whose elements are function-closures
		func() {x += x},
		func() {x -= 1},
		func() {x *= x},
		func() {x /= x},
	}

	fch := make(chan func())	// function-type channel

	go func(ch chan func(), n int, fns ...func()) {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- func(fns ...func()) func() {
				return fns[i%len(fns)]
			}(fns...)
		}
	}(fch, 5, fns...)

	for fn := range fch {
		fn()
		fmt.Println(x)
	}
}
*/

/*
func main() {
	x := 100
	fns := []func() {
		func() {x += x},
		func() {x -= 1},
		func() {x *= x},
		func() {x /= x},
	}

	fch := make(chan func())

	go func(ch chan func(), n int, fns ...func()) {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- func(fns ...func()) func() {
				return fns[i%len(fns)]
			}(fns...)
		}
	}(fch , 5, fns...)

	for fn := range fch {
		fn()
		fmt.Println(x)
	}
}
*/
