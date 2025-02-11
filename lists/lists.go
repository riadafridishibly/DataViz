// Package lists provides an abstract List interface.
//
// In computer science, a list or sequence is an abstract data type that represents an ordered sequence of values, where the same value may occur more than once. An instance of a list is a computer representation of the mathematical concept of a finite sequence; the (potentially) infinite analog of a list is a stream.  Lists are a basic example of containers, as they contain other values. If the same value occurs multiple times, each occurrence is considered a distinct item.
//
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
package lists

import (
	"github.com/Arafatk/Dataviz/containers"
	"github.com/Arafatk/Dataviz/utils"
)

// List interface that all lists implement
type List[T any] interface {
	Get(index int) (T, bool)
	Remove(index int) T
	Add(values ...T)
	Contains(values ...T) bool
	Sort(comparator utils.Comparator)
	Swap(index1, index2 int)
	Insert(index int, values ...T)

	containers.Container[T]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
