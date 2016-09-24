package profiler_test

import "fmt"

func Factorial(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	return fact
}

func Example() {
	prof := NewProfiler()
	prof.StartProfiler("fact")
	_ = Factorial(100)
	t := prof.EndProfiler()
	fmt.Printf("\nFactorial took %d microseconds", t)
}
