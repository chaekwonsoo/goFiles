package main

/*
import "fmt"
import "time"
*/
/*
func main() {
	t1 := time.Now().UnixNano()
	
	top := make(chan int)
	north, south := (chan int)(nil), top

	for i := 0; i < 100000; i++ {
		north, south = south, make(chan int)
		go func(sch, nch chan int) {
			nch <- 1 + <-sch
		}(south ,north)
	}
	south <-0

	x := <-top
	fmt.Printf("top is (%d)\n", x)
	t2 := time.Now().UnixNano()
	fmt.Println(float64(t2 - t1))
	
}
*/
/*
func main() {
	t1 := time.Now().UnixNano()
	
	top := make(chan int)
	north, south := (chan int)(nil), top

	for i := 0; i < 100000; i++ {
		north, south = south, make(chan int)
		go func(sch, nch chan int) {
			nch <- 1 + <- sch
		}(south, north)
	}
	
	south <- 0

	x := <- top

	fmt.Println("top is (%d)\n", x)
	t2 := time.Now().UnixNano()
	fmt.Println(float64(t2 - t1))
}*/
