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

func BenchmarkRead(b *testing.B) {
	counter := &Counter{}
	var value int64
	for n := 0; n < b.N; n++ {
		value += counter.Read()
	}

	fmt.Println("final value of temp counter", value)
}

func BenchmarkReadAtomic(b *testing.B) {
	counter := &Counter{}
	var value int64
	for n := 0; n < b.N; n++ {
		value += counter.ReadAtomic()
	}

	fmt.Println("final value of temp counter", value)
}

func BenchmarkReadLock(b *testing.B) {
	counter := &Counter{}
	var value int64
	for n := 0; n < b.N; n++ {
		value += counter.ReadLock()
	}

	fmt.Println("final value of temp counter", value)
}
