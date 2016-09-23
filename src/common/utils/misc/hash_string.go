package misc

import (
	"hash/fnv"
)

func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func GetHash(s string, count int) uint32 {
	if count == 1 {
		return 0
	}

	hValue := Hash(s)

	return hValue % uint32(count)
}
