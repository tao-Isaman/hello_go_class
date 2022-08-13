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

## Basic Syntax - Type
```go 

bool

string

int int8 int16 int32 int64 

uint uint8 uint16 uint32 uint64

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

```

## Basic Syntax - Array
```go 

var array [5]string // name[0] - name[4]

primes := [...]int{2, 3, 5, 7, 11, 13}

// for loop
for i, prime := range primes {
    fmt.Println(i, prime)
}

// under score requirer
for _, prime := range primes {
    fmt.Println(prime)
}

```

## Basic Syntax - Type Cast with Rune 
```go
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
```
## Basic Syntax - Map easy 
```go
	m := map[string]int {
		"G" : 71,
		"O" : 79,
		"P" : 90,
		"H" : 72,
		"E" : 69,
		"R" : 82,
	}

	var keys []string
	var vals []int

	for k, v := range m {
		keys := append(keys, k)
		vals := append(vals, v)
		fmt.Println(keys, vals)
	}
```

## Basic Syntax - Create new Type
```go
type Int int 
func (i Int) toString() string {
	return strconv.Itoa(int(i))
}

type String string
func (s *String) toUpperCase() {
	*s = String(strings.ToUpper(string(*s)))
}

func main(){
	var test Int = 2
	fmt.Println(test.toString())

	var s String = "asdsd"
	s.toUpperCase()
	fmt.Println(s)
}
```

## Basic Syntax - Simple struc
```go
type rectangle struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

var jsonString = `{"width":10,"height":20}`
func main() {
	rec1 := rectangle{Width: 10, Height: 20}
	b, err := json.Marshal(&rec1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	var rec2 rectangle
	if err := json.Unmarshal([]byte(jsonString), &rec2); err != nil {
		log.Panic(err)
	}

	fmt.Printf("%#v \n", rec2)
}
```

## Basic Syntax - defer 
```go
func main(){
	printInt(9)
}


func printInt(n int){
	defer fmt.Println(n)

	n += 1
	fmt.Println(n)
}
```

## Basic Syntax - Interface Type
```go
func main(){
	var i interface{}

	i = 10 
	fmt.Printf("type is %T, value is %v \n", i,i)

	i = "ten"
	fmt.Printf("type is %T, value is %v \n", i,i)

	i = struct {
		number int
		text string
	}{
		number: 10,
		text: "ten",
	}
	fmt.Printf("type is %T, value is %v \n", i,i)

	i = func() string {
		return "10"
	}
	fmt.Printf("type is %T, value is %v \n", i,i)
}
```



