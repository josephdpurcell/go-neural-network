package pvector

import (
    "math"
)

/**
 * A vector with X and Y values.
 */
type PVector struct {
    X float64
    Y float64
}

/**
 * The means of creating a PVector.
 */
func PVectorFactory (X, Y float64) PVector {
    p := PVector{
        X: X,
        Y: Y,
    }
    return p
}

/**
 * Add a vector.
 */
func (v1 PVector) Add (v2 PVector) PVector {
    v1.X = v1.X + v2.X
    v1.Y = v1.Y + v2.Y
    return v1
}

/**
 * Subtract a vector.
 */
func (v1 PVector) Sub (v2 PVector) PVector {
    v1.X = v1.X - v2.X
    v1.Y = v1.Y - v2.Y
    return v1
}

/**
 * Scale a vector with multiplication.
 */
func (v1 PVector) Mult (n float64) PVector {
    v1.X = v1.X * n
    v1.Y = v1.Y * n
    return v1
}

/**
 * Scale a vector with division.
 */
func (v1 PVector) Div (n float64) PVector {
    v1.X = v1.X / n
    v1.Y = v1.Y / n
    return v1
}

/**
 * Calculate the magnitude of a vector.
 */
func (v1 PVector) Mag () float64 {
    return math.Sqrt((v1.X * v1.X) + (v1.Y * v1.Y))
}

/**
 * Set the magnitude of a vector.
 */
func (v1 PVector) SetMag(mag float64) PVector {
    return v1.Normalize().Mult(mag)
}

/**
 * Normalize the vector to a unit length of 1.
 */
func (v1 PVector) Normalize () PVector {
    var mag float64 = v1.Mag()
    if (mag != 0) {
        return v1.Div(mag)
    } else {
        return v1
    }
}

/**
 * Limit the magnitude of a vector.
 */
func (v1 PVector) Limit (mag float64) PVector {
    if (v1.Mag() > mag) {
        return v1.Normalize().SetMag(mag)
    } else {
        return v1
    }
}

/*
heading() — the 2D heading of a vector eXpressed as an angle

rotate() — rotate a 2D vector by an angle

lerp() — linear interpolate to another vector

dist() — the Euclidean distance between two vectors (considered as points)

angleBetween() — find the angle between two vectors

dot() — the dot product of two vectors

cross() — the cross product of two vectors (only relevant in three dimensions)

random2D() - make a random 2D vector

random3D() - make a random 3D vector
*/

