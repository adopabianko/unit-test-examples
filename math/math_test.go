package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	got := Add(6, 4)
	want := 10

	assert.Equal(t, got, want, "they should be equal")
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}

func TestSubstract(t *testing.T) {
	got := Substract(10, 5)
	want := 5

	assert.Equal(t, got, want, "they should be equal")
}

func BenchmarkSubstract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}
