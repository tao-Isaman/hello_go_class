package foobar

import "strconv"


var Title string = "foobar"

func Say(a int) string{
	if a % 3 == 0 && a % 5 == 0 {
		return "FooBar"
	} else if a % 3 == 0 {
		return "Foo"
	} else if a % 5 == 0 {
		return "Bar"
	} else {
		return strconv.Itoa(a)
	}
}

