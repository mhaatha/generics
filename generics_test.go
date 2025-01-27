package generics_test

import (
	"testing"

	"github.com/mhaatha/generics"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		generics.AssertEqual(t, 1, 1)
		generics.AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		generics.AssertEqual(t, "hello", "hello")
		generics.AssertNotEqual(t, "hello", "Grace")
	})

	// AssertEqual(t, 1, "1") // uncomment to see the error
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := generics.NewStack[int]()

		// check stack is empty
		generics.AssertTrue(t, myStackOfInts.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfInts.Push(123)
		generics.AssertFalse(t, myStackOfInts.IsEmpty())

		// add another thing, pop it back again
		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		generics.AssertEqual(t, value, 456)
		value, _ = myStackOfInts.Pop()
		generics.AssertEqual(t, value, 123)
		generics.AssertTrue(t, myStackOfInts.IsEmpty())

		// can get the numbers we put in as numbers, not untyped interface{}
		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()
		generics.AssertEqual(t, firstNum+secondNum, 3)
	})
}
