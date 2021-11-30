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
