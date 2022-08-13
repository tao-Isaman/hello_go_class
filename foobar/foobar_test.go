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

func TestGivenTwoWantTwo (t *testing.T) {
	//Given
	given := 2
	want := "2"

	//When
	result := foobar.Say(given)

	//Then
	if result != want {
		t.Errorf("Say %d = %q; want %s", given, want, result)
	}
}

func TestGivenThreeWantFoo (t *testing.T) {
	//Given
	given := 3
	want := "Foo"

	//When
	result := foobar.Say(given)

	//Then
	if result != want {
		t.Errorf("Say %d = %q; want %s", given, want, result)
	}
}

func TestGivenFiveWantFive (t *testing.T) {
	//Given
	given := 5
	want := "Bar"

	//When
	result := foobar.Say(given)

	//Then
	if result != want {
		t.Errorf("Say %d = %q; want %s", given, want, result)
	}
}

func TestGivenFiftyFiveWantFooBar (t *testing.T) {
	//Given
	given := 15
	want := "FooBar"

	//When
	result := foobar.Say(given)

	//Then
	if result != want {
		t.Errorf("Say %d = %q; want %s", given, want, result)
	}
}