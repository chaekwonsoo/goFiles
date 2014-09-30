package main

import "fmt"

func main() {
	var a []int = []int{1,2,3,4,5}

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(&(a[0]))
	fmt.Println(&(a[1]))
	fmt.Println(&(a[2]))

}
