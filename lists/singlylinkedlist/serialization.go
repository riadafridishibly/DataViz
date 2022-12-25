package singlylinkedlist

import (
	"encoding/json"

	"github.com/Arafatk/Dataviz/containers"
)

var _ containers.JSONSerializer = (*List)(nil)
var _ containers.JSONDeserializer = (*List)(nil)

// ToJSON outputs the JSON representation of list's elements.
func (list *List) ToJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}

// FromJSON populates list's elements from the input JSON representation.
func (list *List) FromJSON(data []byte) error {
	elements := []any{}
	err := json.Unmarshal(data, &elements)
	if err == nil {
		list.Clear()
		list.Add(elements...)
	}
	return err
}
