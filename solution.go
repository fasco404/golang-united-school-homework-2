package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import (
	"math"
	"golang.org/x/text/cases"
)

type TypeSideNum int

const SidesCircle TypeSideNum = 0
const SidesTriangle TypeSideNum = 3
const SidesSquare TypeSideNum = 4

func CalcSquare(sideLen float64, sidesNum TypeSideNum) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pi * math.Pow(float64(sideLen), 2)
	case SidesTriangle:
		return (math.Sqrt(3) / 4) * math.Pow(float64(sideLen), 2)
	case SidesSquare:
		return math.Pow(float64(sideLen), 2)
	}
	return 0
}
