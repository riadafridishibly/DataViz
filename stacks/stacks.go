// Package stacks provides an abstract Stack interface.
//
// In computer science, a stack is an abstract data type that serves as a collection of elements, with two principal operations: push, which adds an element to the collection, and pop, which removes the most recently added element that was not yet removed. The order in which elements come off a stack gives rise to its alternative name, LIFO (for last in, first out). Additionally, a peek operation may give access to the top without modifying the stack.
//
// Reference: https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29
package stacks

import "github.com/Arafatk/Dataviz/containers"

// Stack interface that all stacks implement
type Stack interface {
	Push(value any)
	Pop() (value any, ok bool)
	Peek() (value any, ok bool)

	containers.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
