package main

import "testing"

func BenchmarkAitbek(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Aitbek()
	}
}

func BenchmarkFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstSolution()
	}
}

func BenchmarkModdedFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ModdedFirst()
	}
}

func BenchmarkSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SecondSolution()
	}
}

func BenchmarkModdedSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ModdedSecond()
	}
}

//go test -bench . -benchmem -memprofile=mem.out -memprofilerate=1 -cpuprofile=cpu.out
//go tool pprof mem.out
//go tool pprof cpu.out
