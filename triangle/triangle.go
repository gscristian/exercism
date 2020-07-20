// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Kind is the type of triangle returned by KindFromSides().
type Kind string

const ()
	NaT Kind = iota // NaT not a triangle
	Equ             // Equ equilateral
	Iso             // Iso isosceles
	Sca             // Sca scalene
)

// KindFromSides expects three sides as float64 and return the kind (Kind) of the triangle
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	
	if (a == b) && (b == c){
		k = Equ 
	}
	return k
}
