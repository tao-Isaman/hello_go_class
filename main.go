package main

import (
	"fmt"
	"os"
)

var name string = "isaman"


func main(){
	given := "abcefg"

	fmt.Println(matcher(given))
	fmt.Println(matcher("abcef"))
	fmt.Println(matcher("ภาษาไทย"))
}

func matcher(str string) []string {
	var r []string
	s := []rune(str)

	for s = append(s, []rune("*")...); len(s) > 1; s = s[2:] {
		r = append(r, string(s[:2]))
	}
	return r
}

func mainBasic3(){
	matrix := [4][4]int {
		{1,2,3,4},
		{2,3,7,8},
		{1,6,9,2},
		{2,7,3,4},
	}

	result := miltiply(matrix,2)
	for _ , value := range result {
		fmt.Println(value)
	}
}

func miltiply(matrix [4][4]int, n int) [4][4]int {
	result := [4][4]int{}
	for i, arrays := range matrix {
		for j, value := range arrays {
			result[i][j] = value * n
		}
	}

	return result
}

func mainBasic2(){
	primes := [...]int{2, 3, 5, 7, 11, 13}

for _, prime := range primes {
    fmt.Println(prime)
}
}

func mainTemp() {
	n := os.Getenv("NAME")
	name := "isaman"
	if n != "" {
		name = n
	}
	fmt.Println("Hello,", name)

}

func isPalinDrome(a, b, c, d, e int) bool {
	if a == e && b == d {
		return true
	}
	return false
}

func power(b, x int) int {
	number := b
	for i := 1; i < x; i++ {
		number *= b
	}
	return number
}

func mainBasic() {
	println("Hello, World!")
	fmt.Println("Square area of 3 is", squareArea(3))

	a, b := swap(1, 2)
	fmt.Println("Swap 1 and 2 : ", a, b)

}

func squareArea(a float64) float64 {
	return a * a
}

func swap(a, b int) (int, int) {
	return b, a
}
