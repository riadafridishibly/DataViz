// Package maps provides an abstract Map interface.
//
// In computer science, an associative array, map, symbol table, or dictionary is an abstract data type composed of a collection of (key, value) pairs, such that each possible key appears just once in the collection.
//
// Operations associated with this data type allow:
// - the addition of a pair to the collection
// - the removal of a pair from the collection
// - the modification of an existing pair
// - the lookup of a value associated with a particular key
//
// Reference: https://en.wikipedia.org/wiki/Associative_array
package maps

import "github.com/Arafatk/Dataviz/containers"

// Map interface that all maps implement
type Map interface {
	Put(key any, value any)
	Get(key any) (value any, found bool)
	Remove(key any)
	Keys() []any

	containers.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
