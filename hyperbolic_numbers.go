package hyperbolic

import "fmt"

// Hypnum is the structure for hyperbolic numbers
type Hypnum struct {
	X, Y float64
}

// Add method defines the addition of
// hyperbolic numbers
func (h Hypnum) Add(summand Hypnum) Hypnum {
	var result Hypnum
	result.X = h.X + summand.X
	result.Y = h.Y + summand.Y
	return result
}

// Multiply method defines the multiplication of
// hyperbolic numbers
func (h Hypnum) Multiply(factor Hypnum) Hypnum {
	var result Hypnum
	result.X = h.X*factor.X + h.Y*factor.Y
	result.Y = h.X*factor.Y + h.Y*factor.X
	return result
}

//Conjugate method allows to conjugate the hyperbolic
//number
func (h Hypnum) Conjugate() Hypnum {
	var result Hypnum
	result.Y = -h.Y
	return result
}

//Modulus method allows to calculate the modulus of
//the hyperbolic number
func (h Hypnum) Modulus() float64 {
	return h.X*h.X - h.Y*h.Y
}

// ToString allows to print a string representation of the hyperbolic
// numbers
func (h Hypnum) ToString() string {
	return fmt.Sprintf("%f, %f", h.X, h.Y)
}

// NewHypNum creates a hyperbolic number
func NewHypNum(x, y float64) Hypnum {
	return Hypnum{
		X: x,
		Y: y,
	}
}
