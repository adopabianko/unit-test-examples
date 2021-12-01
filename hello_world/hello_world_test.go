package hello_world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	got := HelloWorld("Ado")
	want := "Hello Ado"

	assert.Equal(t, got, want, "they should be equal")
}

func BenchmarkHelloworld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ado")
	}
}
