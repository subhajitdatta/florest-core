package treemap

import (
	"fmt"
	"github.com/jabong/florest-core/src/common/collections/utils"
	"reflect"
	"testing"
)

func TestMapPutGet(t *testing.T) {
	m := NewWith(utils.GetIntComparator())
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(5, "e")
	m.Put(2, "b")
	m.Put(4, "d")
	m.Put(3, "f") //overwrite

	if actualSize := m.Size(); actualSize != 5 {
		t.Errorf("Size() - Got %v expected %v", actualSize, 5)
	}
	if actualKeys, expectedKeys := m.Keys(), []interface{}{1, 2, 3, 4, 5}; !Equals(actualKeys, expectedKeys) {
		t.Errorf("Keys() - Got %v expected %v", actualKeys, expectedKeys)
	}
	if actualValues, expectedValues := m.Values(), []interface{}{"a", "b", "f", "d", "e"}; !Equals(actualValues, expectedValues) {
		t.Errorf("Values() - Got %v expected %v", actualValues, expectedValues)
	}

	// Testing Get Method
	// key,expectedValue,expectedFound
	expectedMap := [][]interface{}{
		{1, "a", true},
		{2, "b", true},
		{8, nil, false},
		{3, "f", true},
		{4, "d", true},
		{5, "e", true},
	}

	for _, expectedArray := range expectedMap {
		actualValue, actualFound := m.Get(expectedArray[0])
		if actualValue != expectedArray[1] || actualFound != expectedArray[2] {
			t.Errorf("Get() - Got %v expected %v", actualValue, expectedArray[1])
		}
	}
}

func TestMapPutGetRemove(t *testing.T) {

	// Testing Put
	m := NewWith(utils.GetStringComparator())
	m.Put("e", 5)
	m.Put("f", 6)
	m.Put("g", 7)
	m.Put("a", 3)
	m.Put("d", 4)
	m.Put("x", 24)
	m.Put("b", 2)
	m.Put("a", 1) //overwrite

	// Testing Remove
	m.Remove("e")
	m.Remove("f")
	m.Remove("e")
	m.Remove("g")
	m.Remove("h")

	if actualKeys, expectedKeys := m.Keys(), []interface{}{"a", "b", "d", "x"}; !Equals(actualKeys, expectedKeys) {
		t.Errorf("Keys() - Got %v expected %v", actualKeys, expectedKeys)
	}

	if actualValues, expectedValues := m.Values(), []interface{}{1, 2, 4, 24}; !Equals(actualValues, expectedValues) {
		t.Errorf("Values() - Got %v expected %v", actualValues, expectedValues)
	}
	if actualSize := m.Size(); actualSize != 4 {
		t.Errorf("Size() - Got %v expected %v", actualSize, 4)
	}

	// Testing Get
	expectedMap := [][]interface{}{
		{"a", 1, true},
		{"b", 2, true},
		{"d", 4, true},
		{"x", 24, true},
		{"e", nil, false},
	}

	for _, expectedArray := range expectedMap {
		actualValue, actualFound := m.Get(expectedArray[0])
		if actualValue != expectedArray[1] || actualFound != expectedArray[2] {
			t.Errorf("Get() - Got %v expected %v", actualValue, expectedArray[1])
		}
	}

	// Remove Everything else
	m.Remove("a")
	m.Remove("b")
	m.Remove("d")
	m.Remove("b")
	m.Remove("x")
	m.Remove("d")

	if actualKeys, expectedKeys := fmt.Sprintf("%s", m.Keys()), "[]"; actualKeys != expectedKeys {
		t.Errorf("Keys() - Got %v expected %v", actualKeys, expectedKeys)
	}
	if actualValues, expectedValues := fmt.Sprintf("%s", m.Values()), "[]"; actualValues != expectedValues {
		t.Errorf("Values() - Got %v expected %v", actualValues, expectedValues)
	}
	if actualSize := m.Size(); actualSize != 0 {
		t.Errorf("Size() - Got %v expected %v", actualSize, 0)
	}
	if isEmpty := m.IsEmpty(); !isEmpty {
		t.Errorf("IsEmpty() - Got %v expected %v", isEmpty, true)
	}
}

func TestMapContainsClear(t *testing.T) {
	m := NewWith(utils.GetIntComparator())
	m.Put(1, "a")
	m.Put(2, "b")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(5, "e")

	// Testing contains method
	if contains := m.Contains(1, 2, 3, 4); !contains {
		t.Errorf("Contains() - Got %v expected %v", contains, true)
	}

	if contains := m.Contains(1, 2, 7, 4); contains {
		t.Errorf("!Contains() - Got %v expected %v", contains, false)
	}

	// Testing Clear method
	m.Clear()

	if actualKeys, expectedKeys := fmt.Sprintf("%s", m.Keys()), "[]"; actualKeys != expectedKeys {
		t.Errorf("Keys() - Got %v expected %v", actualKeys, expectedKeys)
	}
	if actualValues, expectedValues := fmt.Sprintf("%s", m.Values()), "[]"; actualValues != expectedValues {
		t.Errorf("Values() - Got %v expected %v", actualValues, expectedValues)
	}
	if actualSize := m.Size(); actualSize != 0 {
		t.Errorf("Size() - Got %v expected %v", actualSize, 0)
	}
	if isEmpty := m.IsEmpty(); !isEmpty {
		t.Errorf("IsEmpty() - Got %v expected %v", isEmpty, true)
	}

}

func TestMapFirstLast(t *testing.T) {
	m := NewWith(utils.GetIntComparator())
	m.Put(3, "c")
	m.Put(1, "a")
	m.Put(4, "d")
	m.Put(2, "b")
	m.Put(5, "e")

	// Test First
	firstEntry := m.First()
	if actualKey, expectedKey := firstEntry.GetKey(), 1; actualKey != expectedKey {
		t.Errorf("FirstEntry - Key() - Got %v expected %v", actualKey, expectedKey)
	}
	if actualValue, expectedValue := firstEntry.GetValue(), "a"; actualValue != expectedValue {
		t.Errorf("FirstEntry - Key() - Got %v expected %v", actualValue, expectedValue)
	}

	// Test Last
	lastEntry := m.Last()
	if actualKey, expectedKey := lastEntry.GetKey(), 5; actualKey != expectedKey {
		t.Errorf("FirstEntry - Key() - Got %v expected %v", actualKey, expectedKey)
	}
	if actualValue, expectedValue := lastEntry.GetValue(), "e"; actualValue != expectedValue {
		t.Errorf("FirstEntry - Key() - Got %v expected %v", actualValue, expectedValue)
	}

}

func TestMapIterator(t *testing.T) {
	m := NewWith(utils.GetRuneComparator())
	// Test Empty Map
	it := m.Iterator()
	if hasNext := it.HasNext(); hasNext {
		t.Errorf("HasNext() - Shouldn't iterate on empty map")
	}

	m = NewWith(utils.GetIntComparator())
	it = m.Iterator()
	if next := it.Next(); next != nil {
		t.Errorf("Next() - Shouldn't iterate on empty map")
	}

	// Test Map with entries
	m.Put(2, "b")
	m.Put(1, "a")
	m.Put(3, "c")

	it = m.Iterator()
	count := 0
	for it.HasNext() {
		count++
		entry := it.Next()
		key := entry.GetKey()
		value := entry.GetValue()
		switch key {
		case 1:
			if actualValue, expectedValue := value, "a"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 2:
			if actualValue, expectedValue := value, "b"; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		case 3:
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

	if actualKey, expectedKey := entry.GetKey(), 1; actualKey != expectedKey {
		t.Errorf("Got %v expected %v", actualKey, expectedKey)
	}

	if actualValue, expectedValue := entry.GetValue(), "a"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

}

func Equals(a []interface{}, b []interface{}) bool {
	return reflect.DeepEqual(a, b)
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
	m := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkGet(b, m, size)
}

func BenchmarkTreeMapPut100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := NewWith(utils.GetIntComparator())
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapPut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapPut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapPut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkPut(b, m, size)
}

func BenchmarkTreeMapRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	m := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeMapRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	m := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeMapRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	m := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}

func BenchmarkTreeMapRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	m := NewWith(utils.GetIntComparator())
	for n := 0; n < size; n++ {
		m.Put(n, struct{}{})
	}
	b.StartTimer()
	benchmarkRemove(b, m, size)
}
