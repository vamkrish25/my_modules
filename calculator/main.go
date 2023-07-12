package main

import "fmt"

func main() {

	//Arithmetic Operators
	var a int
	fmt.Println("Enter Your First Number: ")
	fmt.Scanln(&a)

	var b int
	fmt.Println("Enter Your Second Number: ")
	fmt.Scanln(&b)

	var ch int
	fmt.Println("Enter Your choice (1.ADD  2.Sub  3.mul  4.div 5.greaterthan): ")
	fmt.Scanln(&ch)

	switch ch {
	case 1:
		fmt.Println("the added value is :", add(a, b))
	case 2:
		fmt.Println("the substracted value is :", sub(a, b))
	case 3:
		fmt.Println("the multiplication value is :", mul(a, b))
	case 4:
		fmt.Println("the division value is :", div(a, b))
	case 5:
		fmt.Println("is greater :", great(a, b))

	}

}

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func mul(x int, y int) int {
	return x * y
}

func div(x int, y int) float32 {
	//return float32(x) / float32(y)
	return float32(x / y)
}

func great(x int, y int) bool {
	if x > y {
		return true
	} else {
		return false
	}
}
