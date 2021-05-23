package main

import "testing"

func BenchmarkUuidAsNonce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuidAsNonce()
	}
}

func BenchmarkLarrysNonce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		larrysNonce()
	}
}

func BenchmarkUnixNanoAsNonce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unixNanoAsNonce()
	}
}

func BenchmarkIncrementalNonce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		incrementalNonce()
	}
}
