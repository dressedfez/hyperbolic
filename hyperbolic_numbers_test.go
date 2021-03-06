package hyperbolic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestAdd tests the addition of two hyperbolic numbers
func TestAdd(t *testing.T) {
	expectedResult := NewHypNum(3, 5)
	a := NewHypNum(1, 2)
	b := NewHypNum(2, 3)
	result := a.Add(b)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
}

//TestMultiply tests the multiplication of two hyperbolic numbers
func TestMultiply(t *testing.T) {
	expectedResult := NewHypNum(8, 7)
	a := NewHypNum(1, 2)
	b := NewHypNum(2, 3)
	result := a.Multiply(b)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
}
