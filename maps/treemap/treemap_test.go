package treemap

import (
	"fmt"
	"testing"
)

func TestMapPut(t *testing.T) {
	m := NewWithIntComparator()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	if actualValue := m.Size(); actualValue != 7 {
		t.Errorf("Got %v expected %v", actualValue, 7)
	}
	if actualValue, expectedValue := m.Keys(), []any{1, 2, 3, 4, 5, 6, 7}; !sameElements(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := m.Values(), []any{"a", "b", "c", "d", "e", "f", "g"}; !sameElements(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	// key,expectedValue,expectedFound
	tests1 := [][]any{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{8, nil, false},
	}

	for _, test := range tests1 {
		// retrievals
		actualValue, actualFound := m.Get(test[0])
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}
}

func TestMapRemove(t *testing.T) {
	m := NewWithIntComparator()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite

	m.Remove(5)
	m.Remove(6)
	m.Remove(7)
	m.Remove(8)
	m.Remove(5)

	if actualValue, expectedValue := m.Keys(), []any{1, 2, 3, 4}; !sameElements(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	if actualValue, expectedValue := m.Values(), []any{"a", "b", "c", "d"}; !sameElements(actualValue, expectedValue) {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := m.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}

	tests2 := [][]any{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, nil, false},
		{6, nil, false},
		{7, nil, false},
		{8, nil, false},
	}

	for _, test := range tests2 {
		actualValue, actualFound := m.Get(test[0])
		if actualValue != test[1] || actualFound != test[2] {
			t.Errorf("Got %v expected %v", actualValue, test[1])
		}
	}

	m.Remove(1)
	m.Remove(4)
	m.Remove(2)
	m.Remove(3)
	m.Remove(2)
	m.Remove(2)

	if actualValue, expectedValue := fmt.Sprintf("%s", m.Keys()), "[]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s", m.Values()), "[]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := m.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
	if actualValue := m.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func sameElements(a []any, b []any) bool {
	if len(a) != len(b) {
		return false
	}
	for _, av := range a {
		found := false
		for _, bv := range b {
			if av == bv {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func TestMapEach(t *testing.T) {
	m := NewWithStringComparator()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	count := 0
	m.Each(func(key any, value any) {
		count++
		if actualValue, expectedValue := count, value; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		switch value {
		case 1:
			if actualValue, expectedValue := key, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := key, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 3:
			if actualValue, expectedValue := key, "c"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
	})
}

func TestMapMap(t *testing.T) {
	m := NewWithStringComparator()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	mappedMap := m.Map(func(key1 any, value1 any) (key2 any, value2 any) {
		return key1, value1.(int) * value1.(int)
	})
	if actualValue, _ := mappedMap.Get("a"); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, "mapped: a")
	}
	if actualValue, _ := mappedMap.Get("b"); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, "mapped: b")
	}
	if actualValue, _ := mappedMap.Get("c"); actualValue != 9 {
		t.Errorf("Got %v expected %v", actualValue, "mapped: c")
	}
	if mappedMap.Size() != 3 {
		t.Errorf("Got %v expected %v", mappedMap.Size(), 3)
	}
}

func TestMapSelect(t *testing.T) {
	m := NewWithStringComparator()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	selectedMap := m.Select(func(key any, value any) bool {
		return key.(string) >= "a" && key.(string) <= "b"
	})
	if actualValue, _ := selectedMap.Get("a"); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, "value: a")
	}
	if actualValue, _ := selectedMap.Get("b"); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, "value: b")
	}
	if selectedMap.Size() != 2 {
		t.Errorf("Got %v expected %v", selectedMap.Size(), 2)
	}
}

func TestMapAny(t *testing.T) {
	m := NewWithStringComparator()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	isAny := m.Any(func(key any, value any) bool {
		return value.(int) == 3
	})
	if isAny != true {
		t.Errorf("Got %v expected %v", isAny, true)
	}
	isAny = m.Any(func(key any, value any) bool {
		return value.(int) == 4
	})
	if isAny != false {
		t.Errorf("Got %v expected %v", isAny, false)
	}
}

func TestMapAll(t *testing.T) {
	m := NewWithStringComparator()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	all := m.All(func(key any, value any) bool {
		return key.(string) >= "a" && key.(string) <= "c"
	})
	if all != true {
		t.Errorf("Got %v expected %v", all, true)
	}
	all = m.All(func(key any, value any) bool {
		return key.(string) >= "a" && key.(string) <= "b"
	})
	if all != false {
		t.Errorf("Got %v expected %v", all, false)
	}
}

func TestMapFind(t *testing.T) {
	m := NewWithStringComparator()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	foundKey, foundValue := m.Find(func(key any, value any) bool {
		return key.(string) == "c"
	})
	if foundKey != "c" || foundValue != 3 {
		t.Errorf("Got %v -> %v expected %v -> %v", foundKey, foundValue, "c", 3)
	}
	foundKey, foundValue = m.Find(func(key any, value any) bool {
		return key.(string) == "x"
	})
	if foundKey != nil || foundValue != nil {
		t.Errorf("Got %v at %v expected %v at %v", foundValue, foundKey, nil, nil)
	}
}

func TestMapChaining(t *testing.T) {
	m := NewWithStringComparator()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)
	chainedMap := m.Select(func(key any, value any) bool {
		return value.(int) > 1
	}).Map(func(key any, value any) (any, any) {
		return key.(string) + key.(string), value.(int) * value.(int)
	})
	if actualValue := chainedMap.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, found := chainedMap.Get("aa"); actualValue != nil || found {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	if actualValue, found := chainedMap.Get("bb"); actualValue != 4 || !found {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
	if actualValue, found := chainedMap.Get("cc"); actualValue != 9 || !found {
		t.Errorf("Got %v expected %v", actualValue, 9)
	}
}

func TestMapIteratorNextOnEmpty(t *testing.T) {
	m := NewWithStringComparator()
	it := m.Iterator()
	it = m.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty map")
	}
}

func TestMapIteratorPrevOnEmpty(t *testing.T) {
	m := NewWithStringComparator()
	it := m.Iterator()
	it = m.Iterator()
	for it.Prev() {
		t.Errorf("Shouldn't iterate on empty map")
	}
}

func TestMapIteratorNext(t *testing.T) {
	m := NewWithStringComparator()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)

	it := m.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		value := it.Value()
		switch key {
		case "a":
			if actualValue, expectedValue := value, 1; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case "b":
			if actualValue, expectedValue := value, 2; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case "c":
			if actualValue, expectedValue := value, 3; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
		if actualValue, expectedValue := value, count; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, 3; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestMapIteratorPrev(t *testing.T) {
	m := NewWithStringComparator()
	m.Put("c", 3)
	m.Put("a", 1)
	m.Put("b", 2)

	it := m.Iterator()
	for it.Next() {
	}
	countDown := m.Size()
	for it.Prev() {
		key := it.Key()
		value := it.Value()
		switch key {
		case "a":
			if actualValue, expectedValue := value, 1; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case "b":
			if actualValue, expectedValue := value, 2; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case "c":
			if actualValue, expectedValue := value, 3; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
		if actualValue, expectedValue := value, countDown; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestMapIteratorBegin(t *testing.T) {
	m := NewWithIntComparator()
	it := m.Iterator()
	it.Begin()
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(2, "b")
	for it.Next() {
	}
	it.Begin()
	it.Next()
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestMapTreeIteratorEnd(t *testing.T) {
	m := NewWithIntComparator()
	it := m.Iterator()
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(2, "b")
	it.End()
	it.Prev()
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestMapIteratorFirst(t *testing.T) {
	m := NewWithIntComparator()
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(2, "b")
	it := m.Iterator()
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 1 || value != "a" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 1, "a")
	}
}

func TestMapIteratorLast(t *testing.T) {
	m := NewWithIntComparator()
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(2, "b")
	it := m.Iterator()
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key, value := it.Key(), it.Value(); key != 3 || value != "c" {
		t.Errorf("Got %v,%v expected %v,%v", key, value, 3, "c")
	}
}

func TestMapSerialization(t *testing.T) {
	m := NewWithStringComparator()
	m.Put("a", "1")
	m.Put("b", "2")
	m.Put("c", "3")

	var err error
	assert := func() {
		if actualValue := m.Keys(); actualValue[0].(string) != "a" || actualValue[1].(string) != "b" || actualValue[2].(string) != "c" {
			t.Errorf("Got %v expected %v", actualValue, "[a,b,c]")
		}
		if actualValue := m.Values(); actualValue[0].(string) != "1" || actualValue[1].(string) != "2" || actualValue[2].(string) != "3" {
			t.Errorf("Got %v expected %v", actualValue, "[1,2,3]")
		}
		if actualValue, expectedValue := m.Size(), 3; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if err != nil {
			t.Errorf("Got error %v", err)
		}
	}

	assert()

	json, err := m.ToJSON()
	assert()

	err = m.FromJSON(json)
	assert()
}

func benchmarkGet(b *testing.B, m *Map, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, m *Map, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Put(n, struct{}{})
		}
	}
}

func benchmarkRemove(b *testing.B, m *Map, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			m.Remove(n)
		}
	}
}

func BenchmarkTreeMapGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := NewWithIntComparator()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := NewWithIntComparator()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := NewWithIntComparator()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := NewWithIntComparator()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapPut100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := NewWithIntComparator()
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapPut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := NewWithIntComparator()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapPut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := NewWithIntComparator()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapPut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := NewWithIntComparator()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := NewWithIntComparator()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeMapRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := NewWithIntComparator()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeMapRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := NewWithIntComparator()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeMapRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := NewWithIntComparator()
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}
