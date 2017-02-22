package triangle

import (
	"fmt"
	"math"
)

const testVersion = 3

// Kind stores the geometric type of a triangle
type Kind int

func (k Kind) String() string {
	switch k {
	case NaT:
		return "not a triangle"
	case Equ:
		return "equilateral triangle"
	case Iso:
		return "isoceles triangle"
	case Sca:
		return "scalene triangle"
	}
	return fmt.Sprintf("Unknown triangle %d", k)
}

const (
	// NaT is not a triangle
	NaT = iota
	// Equ is an eqilateral triangle
	Equ = iota
	// Iso is an isoceles triangle
	Iso = iota
	// Sca is a scalene triangle
	Sca = iota
)

// KindFromSides takes the length of the three sides of a triangle and reports
// if they describe a valid triangle and if so the type.
func KindFromSides(a, b, c float64) Kind {
	// Make sure all sides have length
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	if a+b < c || b+c < a || a+c < b {
		return NaT
	}

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}

	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}
