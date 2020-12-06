package hyperbolic

type Hypnum interface {
	Add(summand Hypnum) Hypnum
	NewHypNum(x, y float64) Hypnum
}

type hypnum struct {
	X, Y float64
}

// method that defines the addition of
// hyperbolic numbers
func (h hypnum) Add(summand hypnum) hypnum {
	var result hypnum
	result.X = h.X + summand.X
	result.Y = h.Y + summand.Y
	return result
}

// method that defines the multiplication of
// hyperbolic numbers
func (h hypnum) Multiply(factor hypnum) hypnum {
	var result hypnum
	result.X = h.X*factor.X + h.Y*factor.Y
	result.Y = h.X*factor.Y + h.Y*factor.X
	return result
}

// creates an hyperbolic number entity
func NewHypNum(x, y float64) hypnum {
	return hypnum{
		X: x,
		Y: y,
	}
}
