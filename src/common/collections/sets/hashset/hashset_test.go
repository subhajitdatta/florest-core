package hashset

import (
	"testing"
)

func TestSetAdd(t *testing.T) {
	set := New()
	set.Add()
	set.Add(1)
	set.Add(2)
	set.Add(2, 3)
	set.Add()
	if isEmpty := set.IsEmpty(); isEmpty {
		t.Errorf("IsEmpty - Got %v expected %v", isEmpty, false)
	}
	if size := set.Size(); size != 3 {
		t.Errorf("Size - Got %v expected %v", size, 3)
	}

	if actualValues, expectedValues := set.Values(), []interface{}{1, 2, 3}; !Equals(actualValues, expectedValues) {
		t.Errorf("Values() - Got %v expected %v", actualValues, expectedValues)
	}
}

func TestSetContainsClear(t *testing.T) {
	set := New()
	set.Add(3, 2, 1)
	set.Add(2, 3)
	set.Add()
	if contains := set.Contains(); !contains {
		t.Errorf("Contains() - Got %v expected %v", contains, true)
	}
	if contains := set.Contains(1); !contains {
		t.Errorf("Contains() - Got %v expected %v", contains, true)
	}
	if contains := set.Contains(1, 2, 3); !contains {
		t.Errorf("Contains() - Got %v expected %v", contains, true)
	}
	if contains := set.Contains(1, 2, 3, 4); contains {
		t.Errorf("Contains() - Got %v expected %v", contains, false)
	}

	set.Clear()

	if contains := set.Contains(); !contains {
		t.Errorf("Contains() - Got %v expected %v", contains, true)
	}
	if contains := set.Contains(1); contains {
		t.Errorf("Contains() - Got %v expected %v", contains, false)
	}
}

func TestSetRemove(t *testing.T) {
	set := New()
	set.Add(3, 2, 1)
	set.Remove()
	if size := set.Size(); size != 3 {
		t.Errorf("Got %v expected %v", size, 3)
	}
	set.Remove(1)
	if size := set.Size(); size != 2 {
		t.Errorf("Got %v expected %v", size, 2)
	}
	set.Remove(3)
	set.Remove(3)
	set.Remove()
	set.Remove(2)
	if size := set.Size(); size != 0 {
		t.Errorf("Got %v expected %v", size, 0)
	}
}

func benchmarkContains(b *testing.B, set *Set, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Contains(n)
		}
	}
}

func benchmarkAdd(b *testing.B, set *Set, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Add(n)
		}
	}
}

func benchmarkRemove(b *testing.B, set *Set, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			set.Remove(n)
		}
	}
}

func BenchmarkHashSetContains100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkHashSetContains1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkHashSetContains10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkHashSetContains100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkHashSetAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := New()
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkHashSetAdd1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkHashSetAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkHashSetAdd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkHashSetRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkHashSetRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkHashSetRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkHashSetRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := New()
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func Equals(a []interface{}, b []interface{}) bool {
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
