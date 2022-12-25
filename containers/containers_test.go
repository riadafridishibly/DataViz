// All data structures must implement the container structure

package containers

import (
	"testing"

	"github.com/Arafatk/Dataviz/utils"
)

// For testing purposes
type ContainerTest struct {
	values []any
}

func (container ContainerTest) Empty() bool {
	return len(container.values) == 0
}

func (container ContainerTest) Size() int {
	return len(container.values)
}

func (container ContainerTest) Clear() {
	container.values = []any{}
}

func (container ContainerTest) Values() []any {
	return container.values
}

func TestGetSortedValuesInts(t *testing.T) {
	container := ContainerTest{}
	container.values = []any{5, 1, 3, 2, 4}
	values := GetSortedValues(container, utils.IntComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1].(int) > values[i].(int) {
			t.Errorf("Not sorted!")
		}
	}
}

func TestGetSortedValuesStrings(t *testing.T) {
	container := ContainerTest{}
	container.values = []any{"g", "a", "d", "e", "f", "c", "b"}
	values := GetSortedValues(container, utils.StringComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1].(string) > values[i].(string) {
			t.Errorf("Not sorted!")
		}
	}
}
