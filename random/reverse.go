package random

import (
    "math/rand"
)

/**
 * Generate a random float64 between max and min.
 */
func Random(min, max float64) float64 {
    return (rand.Float64() * (max - min)) + min
}
