package main

import (
	"fmt"
	"time"
	"math"
	"runtime"
)

type Big struct {
	list []int
}


func add(in chan *Big, out chan *Big) {
	for {
		data := <- in
		for i, _ := range data.list {
			data.list[i] += data.list[i]
		}
		out <- data
	}
}

func square(int chan *Big, out chan *Big) {
	for {	
		data := <- int
		for i, _ := range data.list {
			data.list[i] *= data.list[i]
		}
		out <- data
	}
}

func sqrt(in chan *Big, out chan *Big) {
	for {
		data := <- in
		for i, _ := range data.list {
			data.list[i] = int(math.Sqrt(float64(data.list[i])))
		}
		out <- data
	}
}

func addone(in chan *Big, out chan *Big) {
	for {	
		data := <- in
		for i := range data.list {
			data.list[i] += 1
		}
		out <- data
	}
}

func pipeline(in chan *Big, out chan *Big, n int) {
	ch1 := make(chan *Big, n)
	ch2 := make(chan *Big, n)
	ch3 := make(chan *Big, n)
	ch4 := make(chan *Big, n)

	go add(in, ch1)
	go square(ch1, ch2)
	go sqrt(ch2, ch3)
	go square(ch3, ch4)
	go addone(ch4, out)
}

func main() {
	t0 := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())

	n := 100000
	in, out := make(chan *Big, n), make(chan *Big, n)
	bigdatas := make([]*Big, n)

	for i := 0; i < n; i++ {
		bigdatas[i] = new(Big)
		bigdatas[i].list = make([]int, 1000)
		for j, _ := range bigdatas[i].list {
			bigdatas[i].list[j] = i + j
		}
	}

	go pipeline(in, out, n)

	go func() {
		for _, bigdata := range bigdatas {
			in <- bigdata
		}
	}()

	for _, _ = range bigdatas {
		fmt.Println((<-out).list[0])
	}

	fmt.Printf("The sequential took %v to run.\n", time.Now().Sub(t0))
}
