package main

import "testing"

type Benchmark struct {
	Name       string
	Parameters []int
}

// Sub Benchmark
func BenchmarkTest(b *testing.B) {
	b.Run("Sum Number", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sumNumber(20, 30)
		}
	})

	b.Run("Say Goodbye", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sayGoodbye("Hutama Trirahmanto")
		}
	})
}

// Table Benchmark
func BenchmarkTable(b *testing.B) {
	benchData := []Benchmark{
		{
			Name:       "Sum Number1",
			Parameters: []int{21, 30},
		},
		{
			Name:       "Sum Number2",
			Parameters: []int{890, 12576},
		},
	}

	for _, bench := range benchData {
		b.Run(bench.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sumNumber(bench.Parameters[0], bench.Parameters[1])
			}
		})
	}
}
