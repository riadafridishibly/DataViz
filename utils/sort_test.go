package utils

import (
	"math/rand"
	"testing"
)

func TestSortInts(t *testing.T) {
	ints := []int{}
	ints = append(ints, 4)
	ints = append(ints, 1)
	ints = append(ints, 2)
	ints = append(ints, 3)

	Sort(ints, IntComparator)

	for i := 1; i < len(ints); i++ {
		if ints[i-1] > ints[i] {
			t.Errorf("Not sorted!")
		}
	}

}

func TestSortStrings(t *testing.T) {

	strings := []string{}
	strings = append(strings, "d")
	strings = append(strings, "a")
	strings = append(strings, "b")
	strings = append(strings, "c")

	Sort(strings, StringComparator)

	for i := 1; i < len(strings); i++ {
		if strings[i-1] > strings[i] {
			t.Errorf("Not sorted!")
		}
	}
}

func TestSortStructs(t *testing.T) {
	type User struct {
		id   int
		name string
	}

	byID := func(a, b any) int {
		c1 := a.(User)
		c2 := b.(User)
		switch {
		case c1.id > c2.id:
			return 1
		case c1.id < c2.id:
			return -1
		default:
			return 0
		}
	}

	// o1,o2,expected
	users := []any{
		User{4, "d"},
		User{1, "a"},
		User{3, "c"},
		User{2, "b"},
	}

	Sort(users, byID)

	for i := 1; i < len(users); i++ {
		if users[i-1].(User).id > users[i].(User).id {
			t.Errorf("Not sorted!")
		}
	}
}

func TestSortRandom(t *testing.T) {
	ints := []any{}
	for i := 0; i < 10000; i++ {
		ints = append(ints, rand.Int())
	}
	Sort(ints, IntComparator)
	for i := 1; i < len(ints); i++ {
		if ints[i-1].(int) > ints[i].(int) {
			t.Errorf("Not sorted!")
		}
	}
}

func BenchmarkGoSortRandom(b *testing.B) {
	b.StopTimer()
	ints := []any{}
	for i := 0; i < 100000; i++ {
		ints = append(ints, rand.Int())
	}
	b.StartTimer()
	Sort(ints, IntComparator)
	b.StopTimer()
}
