package arraylist

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
	return json.Marshal(list.elements[:list.size])
}

// FromJSON populates list's elements from the input JSON representation.
func (list *List[T]) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &list.elements)
	if err == nil {
		list.size = len(list.elements)
	}
	return err
}
