# go class by P'Yod

## Basic Syntax - Variable

```go

name := 'tao-isaman'
number := '0942257558'
int_64 := 25

```

## Basic Syntax - Function
```go
func add(a int, b int) int {
    return a + b
}

func add(a, b int) int {
    return a + b
}

func squareArea(a float64) float64 {
    return a * a
}

func swap(a, b int) (int, int){
	return b, a
}
```

## Basic Syntax - control Flow if/else
```go
if ok := IsCoreect; ok {
    println("It' correct")
}

same

ok := IsCorrect()
if ok {
    println("It' correct")
}

```

## Basic Syntax - loop 
```go 

// for loop
for i := 0; i < 10; i++ {

}

// while loop
for i <= 10 {

}

// infinity loop
for {

}

```
## Basic Syntax - get ENV
```go 
var name string = "isaman"

func main() {
	n := os.Getenv("NAME")
	name := "isaman"
	if n != "" {
		name = n
	}
	fmt.Println("Hello,", name)

}
```

## Basic Syntax - Unit Test 
```go
package foobar_test

import (
	"testing"

	"github.com/tao-isaman/hello_golang/foobar"
)

func TestGivenOneWantOne (t *testing.T) {

	//Give
	given := 1
	want := "1"

	//When
	result := foobar.Say(given)

	//Then
	if result != want {
		t.Errorf("Say %d = %q; want %s", given, want, result)
	}
}
```

