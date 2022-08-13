package main

import (
	"fmt"
)

func main() {

	fmt.Println(power(2,5))

}

func power(b, x int) int {
	number := b
	for i := 1; i < x; i++{
		number *= b
	}

	return number
}

func mainBasic() {
	println("Hello, World!")
	fmt.Println("Square area of 3 is", squareArea(3))

	a,b := swap(1,2)
	fmt.Println("Swap 1 and 2 : ", a, b)

}

func squareArea(a float64) float64 {
    return a * a
}

func swap(a, b int) (int, int){
	return b,a
}