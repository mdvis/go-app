package main

import "fmt"

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[string]float64{
		"first":  34.99,
		"second": 12.45,
	}

	fmt.Printf("sum: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
