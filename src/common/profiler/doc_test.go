package profiler_test

import (
	"fmt"

	"github.com/jabong/florest-core/src/common/profiler"
)

func Factorial(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	return fact
}

func Example() {
	prof := profiler.NewProfiler()
	prof.StartProfile("fact")
	_ = Factorial(100)
	t := prof.EndProfile()
	fmt.Printf("\nFactorial took %d microseconds", t)
}
