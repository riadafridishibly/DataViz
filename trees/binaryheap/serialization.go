package binaryheap

import "github.com/Arafatk/Dataviz/containers"

var _ containers.JSONSerializer = (*Heap)(nil)
var _ containers.JSONDeserializer = (*Heap)(nil)

// ToJSON outputs the JSON representation of list's elements.
func (heap *Heap) ToJSON() ([]byte, error) {
	return heap.list.ToJSON()
}

// FromJSON populates list's elements from the input JSON representation.
func (heap *Heap) FromJSON(data []byte) error {
	return heap.list.FromJSON(data)
}
