package main

import "testing"

//func BenchmarkAitbek(b *testing.B)  {
//	for i := 0; i<b.N; i++ {
//		Aitbek()
//	}
//}

func BenchmarkMine(b *testing.B)  {
	for i := 0; i<b.N; i++ {
		MySolution()
	}
}