package treemap

import "github.com/Arafatk/Dataviz/containers"

func assertJSONSerializerDeserializer[K comparable, V any]() {
	var _ containers.JSONSerializer = (*Map[K, V])(nil)
	var _ containers.JSONDeserializer = (*Map[K, V])(nil)
}

// ToJSON outputs the JSON representation of list's elements.
func (m *Map[K, V]) ToJSON() ([]byte, error) {
	return m.tree.ToJSON()
}

// FromJSON populates list's elements from the input JSON representation.
func (m *Map[K, V]) FromJSON(data []byte) error {
	return m.tree.FromJSON(data)
}
