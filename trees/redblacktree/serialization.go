package redblacktree

import (
	"encoding/json"

	"github.com/Arafatk/Dataviz/containers"
	"github.com/Arafatk/Dataviz/utils"
)

func assertJSONSerializerDeserializer[K comparable, V any]() {
	var _ containers.JSONSerializer = (*Tree[K, V])(nil)
	var _ containers.JSONDeserializer = (*Tree[K, V])(nil)
}

// ToJSON outputs the JSON representation of list's elements.
func (tree *Tree[K, V]) ToJSON() ([]byte, error) {
	elements := make(map[string]any)
	it := tree.Iterator()
	for it.Next() {
		elements[utils.ToString(it.Key())] = it.Value()
	}
	return json.Marshal(&elements)
}

type str string

// FromJSON populates list's elements from the input JSON representation.
func (tree *Tree[K, V]) FromJSON(data []byte) error {
	elements := make(map[K]V)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		tree.Clear()
		for key, value := range elements {
			tree.Put(key, value)
		}
	}
	return err
}
