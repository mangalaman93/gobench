package gobench

import (
	"fmt"
	"testing"
)

func BenchmarkIncr(b *testing.B) {
	counter := &Counter{}
	for n := 0; n < b.N; n++ {
		counter.Incr()
	}

	fmt.Println("final value of counter", counter.Value)
}

func BenchmarkIncrAtomic(b *testing.B) {
	counter := &Counter{}
	for n := 0; n < b.N; n++ {
		counter.IncrAtomic()
	}

	fmt.Println("final value of counter", counter.Value)
}

func BenchmarkIncrLock(b *testing.B) {
	counter := &Counter{}
	for n := 0; n < b.N; n++ {
		counter.IncrLock()
	}

	fmt.Println("final value of counter", counter.Value)
}
