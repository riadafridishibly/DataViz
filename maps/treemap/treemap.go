// Package treemap implements a map backed by red-black tree.
//
// Elements are ordered by key in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package treemap

import (
	"fmt"
	"strings"

	"github.com/Arafatk/Dataviz/maps"
	rbt "github.com/Arafatk/Dataviz/trees/redblacktree"
	"github.com/Arafatk/Dataviz/utils"
)

func assertMap[K comparable, V any]() {
	var _ maps.Map[K, V] = (*Map[K, V])(nil)
}

// Map holds the elements in a red-black tree
type Map[K comparable, V any] struct {
	tree *rbt.Tree[K, V]
}

// NewWith instantiates a tree map with the custom comparator.
func NewWith[K comparable, V any](comparator utils.Comparator) *Map[K, V] {
	return &Map[K, V]{tree: rbt.NewWith[K, V](comparator)}
}

// NewWithIntComparator instantiates a tree map with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{tree: rbt.NewWithIntComparator[K, V]()}
}

// NewWithStringComparator instantiates a tree map with the StringComparator, i.e. keys are of type string.
func NewWithStringComparator[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{tree: rbt.NewWithStringComparator[K, V]()}
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Put(key K, value V) {
	m.tree.Put(key, value)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	return m.tree.Get(key)
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, V]) Remove(key K) {
	m.tree.Remove(key)
}

// Empty returns true if map does not contain any elements
func (m *Map[K, V]) Empty() bool {
	return m.tree.Empty()
}

// Size returns number of elements in the map.
func (m *Map[K, V]) Size() int {
	return m.tree.Size()
}

// Keys returns all keys in-order
func (m *Map[K, V]) Keys() []K {
	return m.tree.Keys()
}

// Values returns all values in-order based on the key.
func (m *Map[K, V]) Values() []V {
	return m.tree.Values()
}

// Clear removes all elements from the map.
func (m *Map[K, V]) Clear() {
	m.tree.Clear()
}

// Min returns the minimum key and its value from the tree map.
// Returns nil, nil if map is empty.
func (m *Map[K, V]) Min() (key K, value V, ok bool) {
	if node := m.tree.Left(); node != nil {
		return node.Key, node.Value, true
	}
	return key, value, false
}

// Max returns the maximum key and its value from the tree map.
// Returns nil, nil if map is empty.
func (m *Map[K, V]) Max() (key K, value V, ok bool) {
	if node := m.tree.Right(); node != nil {
		return node.Key, node.Value, true
	}
	return key, value, false
}

// String returns a string representation of container
func (m *Map[K, V]) String() string {
	str := "TreeMap\nmap["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"
}

// Visualizer makes a visual image demonstrating the treemap data structure
// using dot language and Graphviz. It first producs a dot string corresponding
// to the treemap and then runs graphviz to output the resulting image to a file.
func (m *Map[K, V]) Visualizer(fileName string) bool {
	return m.tree.Visualizer(fileName)
}
