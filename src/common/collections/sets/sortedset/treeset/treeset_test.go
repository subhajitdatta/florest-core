package treeset

import (
	"github.com/jabong/florest-core/src/common/collections/utils"
	"reflect"
	"testing"
)

func TestSetAdd(t *testing.T) {
	set := NewWith(utils.GetIntComparator())
	set.Add()
	set.Add(3)
	set.Add(2)
	set.Add(2, 1)
	set.Add()
	if actualValue := set.IsEmpty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := set.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValues, expectedValues := set.Values(), []interface{}{1, 2, 3}; !Equals(actualValues, expectedValues) {
		t.Errorf("Values() - Got %v expected %v", actualValues, expectedValues)
	}
}

func TestSetContainsClear(t *testing.T) {
	set := NewWith(utils.GetIntComparator())
	set.Add(3, 1, 2)
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
	set := NewWith(utils.GetIntComparator())
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

func TestMapFirstLast(t *testing.T) {
	set := NewWith(utils.GetStringComparator())
	set.Add("e", "a", "b")

	// Test First
	if actualValue, expectedValue := set.First(), "a"; actualValue != expectedValue {
		t.Errorf("FirstElement - Got %v expected %v", actualValue, expectedValue)
	}

	// Test Last
	if actualValue, expectedValue := set.Last(), "e"; actualValue != expectedValue {
		t.Errorf("LastElement - Got %v expected %v", actualValue, expectedValue)
	}

}

func TestSetIterator(t *testing.T) {
	set := NewWith(utils.GetStringComparator())

	// Test Empty Set
	it := set.Iterator()
	if hasNext := it.HasNext(); hasNext {
		t.Errorf("HasNext() - Shouldn't iterate on empty set")
	}

	set = NewWith(utils.GetStringComparator())
	it = set.Iterator()
	if next := it.Next(); next != nil {
		t.Errorf("Next() - Shouldn't iterate on empty set")
	}

	// Test Set with elements
	set.Add("c", "a", "b")
	set.Add("c", "b")
	set.Add()

	it = set.Iterator()
	count := 0
	for it.HasNext() {
		count++
		entry := it.Next()
		key := entry.GetKey()
		value := entry.GetValue()
		switch key {
		case 0:
			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 1:
			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, "c"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			t.Errorf("Too many")
		}
	}
	if actualCount, expectedCount := count, 3; actualCount != expectedCount {
		t.Errorf("Count() - Got %v expected %v", actualCount, expectedCount)
	}

	// Test Reset
	it.Reset()

	entry := it.Next()

	if actualKey, expectedKey := entry.GetKey(), 0; actualKey != expectedKey {
		t.Errorf("Got %v expected %v", actualKey, expectedKey)
	}

	if actualValue, expectedValue := entry.GetValue(), "a"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
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

func BenchmarkTreeSetContains100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkTreeSetContains1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkTreeSetContains10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkTreeSetContains100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkContains(b, set, size)
}

func BenchmarkTreeSetAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := NewWith(utils.GetIntComparator())
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkTreeSetAdd1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkTreeSetAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkTreeSetAdd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, set, size)
}

func BenchmarkTreeSetRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	set := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkTreeSetRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	set := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkTreeSetRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func BenchmarkTreeSetRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	set := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		set.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, set, size)
}

func Equals(a []interface{}, b []interface{}) bool {
	return reflect.DeepEqual(a, b)
}
