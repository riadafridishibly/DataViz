package avltree

import (
	"encoding/json"

	"github.com/Arafatk/Dataviz/containers"
	"github.com/Arafatk/Dataviz/utils"
)

var _ containers.JSONSerializer = (*Tree)(nil)
var _ containers.JSONDeserializer = (*Tree)(nil)

// ToJSON outputs the JSON representation of list's elements.
func (tree *Tree) ToJSON() ([]byte, error) {
	elements := make(map[string]any)
	it := tree.Iterator()
	for it.Next() {
		elements[utils.ToString(it.Key())] = it.Value()
	}
	return json.Marshal(&elements)
}

// FromJSON populates list's elements from the input JSON representation.
func (tree *Tree) FromJSON(data []byte) error {
	elements := make(map[string]any)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		tree.Clear()
		for key, value := range elements {
			tree.Put(key, value)
		}
	}
	return err
}
