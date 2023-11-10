package kata

import (
	"math"
)

func FindNextSquare(sq int64) int64 {
	sqrt := math.Sqrt(float64(sq))
	if sqrt != math.Floor(sqrt) {
		return -1
	}
	return int64((sqrt + 1) * (sqrt + 1))
}
