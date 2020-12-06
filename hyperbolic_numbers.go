package hyperbolic

type Hypnum interface {
	Add(summand Hypnum) Hypnum
	NewHypNum(x, y float64) Hypnum
}

type hypnum struct {
	x, y float64
}

// method that defines the addition of
// hyperbolic numbers
func (h hypnum) Add(summand hypnum) hypnum {
	var result hypnum
	result.x = h.x + summand.x
	result.y = h.y + summand.y
	return result
}

// method that defines the multiplication of
// hyperbolic numbers
func (h hypnum) Multiply(factor hypnum) hypnum {
	var result hypnum
	result.x = h.x*factor.x + h.y*factor.y
	result.y = h.x*factor.y + h.y*factor.x
	return result
}

// creates an hyperbolic number entity
func NewHypNum(x, y float64) hypnum {
	return hypnum{
		x: x,
		y: y,
	}
}