package utils

import (
	"github.com/jabong/florest-core/src/common/collections"
	"github.com/jabong/florest-core/src/common/collections/utils/comparators"
)

// GetStringComparator returns the string comparator
func GetStringComparator() collections.Comparator {
	return comparators.NewStringComparator()
}

// GetIntComparator returns the int comparator
func GetIntComparator() collections.Comparator {
	return comparators.NewIntComparator()
}

// GetInt8Comparator returns the int8 comparator
func GetInt8Comparator() collections.Comparator {
	return comparators.NewInt8Comparator()
}

// GetInt16Comparator returns the int16 comparator
func GetInt16Comparator() collections.Comparator {
	return comparators.NewInt16Comparator()
}

// GetInt32Comparator returns the int32 comparator
func GetInt32Comparator() collections.Comparator {
	return comparators.NewInt32Comparator()
}

// GetInt64Comparator returns the int64 comparator
func GetInt64Comparator() collections.Comparator {
	return comparators.NewInt64Comparator()
}

// GetUIntComparator returns the uint comparator
func GetUIntComparator() collections.Comparator {
	return comparators.NewUIntComparator()
}

// GetUInt8Comparator returns the uint8 comparator
func GetUInt8Comparator() collections.Comparator {
	return comparators.NewUInt8Comparator()
}

// GetUInt16Comparator returns the uint16 comparator
func GetUInt16Comparator() collections.Comparator {
	return comparators.NewUInt16Comparator()
}

// GetUInt32Comparator returns the uint32 comparator
func GetUInt32Comparator() collections.Comparator {
	return comparators.NewUInt32Comparator()
}

// GetUInt64Comparator returns the uint64 comparator
func GetUInt64Comparator() collections.Comparator {
	return comparators.NewUInt64Comparator()
}

// GetByteComparator returns the byte comparator
func GetByteComparator() collections.Comparator {
	return comparators.NewByteComparator()
}

// GetRuneComparator returns the rune comparator
func GetRuneComparator() collections.Comparator {
	return comparators.NewRuneComparator()
}
