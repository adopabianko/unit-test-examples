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

func TestSubstract(t *testing.T) {
	got := Substract(10, 5)
	want := 5

	assert.Equal(t, got, want, "they should be equal")
}
