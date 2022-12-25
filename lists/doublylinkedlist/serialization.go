package doublylinkedlist

import (
	"encoding/json"

	"github.com/Arafatk/Dataviz/containers"
)

func assertJSONSerializerDeserializer[T comparable]() {
	var _ containers.JSONSerializer = (*List[T])(nil)
	var _ containers.JSONDeserializer = (*List[T])(nil)
}

// ToJSON outputs the JSON representation of list's elements.
func (list *List[T]) ToJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}

// FromJSON populates list's elements from the input JSON representation.
func (list *List[T]) FromJSON(data []byte) error {
	elements := []T{}
	err := json.Unmarshal(data, &elements)
	if err == nil {
		list.Clear()
		list.Add(elements...)
	}
	return err
}
