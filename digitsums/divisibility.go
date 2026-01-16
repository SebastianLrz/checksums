package digitsums

import "math"

// IsDivisibleBy3 prüft, ob die gegebene Zahl n durch 3 teilbar ist.
//
// Randbedingung: Verwenden Sie NICHT den Modulo-Operator (%).
func IsDivisibleBy3(n int) bool {
	m := float64(n) / float64(3)

	if m > 0 {
		if m == math.Floor(m) {
			return true
		} else {
			return false
		}
	}
	if m < 0 {
		if m == math.Ceil(m) {
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}
