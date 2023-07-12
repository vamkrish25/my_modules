package main

import "fmt"

type std struct {
	id   int
	name string
	avg  int
}

func helloworld() {
	fmt.Println("Hello World!!!")
}

type Square struct {
	side float64
}

func (s Square) calculateArea() float64 {
	return s.side * s.side
}

func main() {

	fmt.Println("----------------Struct--------------------")
	var s1 = std{1, "Ben", 67}
	fmt.Println(s1)

	fmt.Println("----------------Pointers------------------")

	mynum := 5
	ptr := &mynum

	fmt.Println(mynum)
	fmt.Println(*ptr)
	fmt.Println(ptr)

	*ptr = 10
	fmt.Println(mynum)
	fmt.Println(&mynum)

	fmt.Println("----------------Function------------------")
	helloworld()

	fmt.Println("------------------Slice-------------------")
	myslice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	fmt.Println("length is :", len(myslice))
	fmt.Println(myslice)
	myslice = append(myslice, "k")
	fmt.Println(myslice)

	fmt.Println("-----------------Method------------------")
	square := Square{side: 5}
	fmt.Println("Area of the square:", square.calculateArea())
}
