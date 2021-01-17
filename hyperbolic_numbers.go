package hyperbolic

import "fmt"

// Hypnum is the structure for hyperbolic numbers
type Hypnum struct {
	x, y float64
}

// Add method defines the addition of
// hyperbolic numbers
func (h Hypnum) Add(summand Hypnum) Hypnum {
	var result Hypnum
	result.x = h.x + summand.x
	result.y = h.y + summand.y
	return result
}

// Multiply method defines the multiplication of
// hyperbolic numbers
func (h Hypnum) Multiply(factor Hypnum) Hypnum {
	var result Hypnum
	result.x = h.x*factor.x + h.y*factor.y
	result.y = h.x*factor.y + h.y*factor.x
	return result
}

// ToString allows to print a string representation of the hyperbolic
// numbers
func (h Hypnum) ToString() string {
	return fmt.Sprintf("(%f, %f)", h.x, h.y)
}

// NewHypNum creates a hyperbolic number
func NewHypNum(x, y float64) Hypnum {
	return Hypnum{
		x: x,
		y: y,
	}
}
