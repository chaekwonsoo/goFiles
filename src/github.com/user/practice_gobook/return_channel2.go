package main

/*
import "fmt"

type emptyI interface{}
type myVector struct {
	a []emptyI
}

func (p *myVector) iterate(c chan emptyI) {
	for _, v := range p.a {
		c <- v
	}
	close(c)
}

func (p *myVector) Iter() chan emptyI {
	c := make(chan emptyI)
	
	go p.iterate(c)
	return c
}

func main() {
	vec := new(myVector)
	vec.a = make([]emptyI, 10)

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			vec.a[i] = "HoHoHo"
		} else {
			vec.a[i] = i * i
		}
	}

	i := 0
	for x := range vec.Iter() {
		switch x.(type) {
			case string:
				fmt.Printf("vec[%d] is %q\n", i, x)
			case int:
				fmt.Printf("vec[%d] is %d\n", i, x)
		}
		i++
	}
}
*/
