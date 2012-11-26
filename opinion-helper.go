package sublogic

import (
  "math"
)

func outOfRange(val float64) bool {
	if (val < 0) || (val > 1) {
		return true
	}
	return false
}

func adjust(x float64) float64 {
    if x >= math.MaxFloat64 {
        return math.MaxFloat64
    }
    if x == 0.0 {
        return 0.0
    }
    return math.Floor(x * 100000000000.0 + 0.5) / 100000000000.0
}

func constrain(x float64) float64 {
    return math.Min(1.0, math.Max(0.0, x))
}
