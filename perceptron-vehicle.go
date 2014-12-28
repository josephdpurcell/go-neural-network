/**
 * A Neural Network Perceptron Example Learning How to Drive.
 *
 * This is an example of creating a rudimentary neural network that learns how
 * to drive.
 *
 * Source: http://natureofcode.com/book/chapter-10-neural-networks/
 */
package main

import (
    "fmt"
    "math/rand"
    "math"
)

/**
 * Generate a random float64 between max and min.
 */
func random(min, max float64) float64 {
    return (rand.Float64() * (max - min)) + min
}

/**
 * A vector with x and y values.
 */
type PVector struct {
    x float64
    y float64
}

/**
 * The means of creating a PVector.
 */
func PVectorFactory (x, y float64) PVector {
    p := PVector{
        x: x,
        y: y,
    }
    return p
}

/**
 * Add a vector.
 */
func (v1 PVector) add (v2 PVector) PVector {
    v1.x = v1.x + v2.x
    v1.y = v1.y + v2.y
    return v1
}

/**
 * Subtract a vector.
 */
func (v1 PVector) sub (v2 PVector) PVector {
    v1.x = v1.x - v2.x
    v1.y = v1.y - v2.y
    return v1
}

/**
 * Scale a vector with multiplication.
 */
func (v1 PVector) mult (n float64) PVector {
    v1.x = v1.x * n
    v1.y = v1.y * n
    return v1
}

/**
 * Scale a vector with division.
 */
func (v1 PVector) div (n float64) PVector {
    v1.x = v1.x / n
    v1.y = v1.y / n
    return v1
}

/**
 * Calculate the magnitude of a vector.
 */
func (v1 PVector) mag () float64 {
    return math.Sqrt((v1.x * v1.x) + (v1.y * v1.y))
}

/**
 * Set the magnitude of a vector.
 */
func (v1 PVector) setMag(mag float64) PVector {
    return v1.normalize().mult(mag)
}

/**
 * Normalize the vector to a unit length of 1.
 */
func (v1 PVector) normalize () PVector {
    var mag float64 = v1.mag()
    if (mag != 0) {
        return v1.div(mag)
    } else {
        return v1
    }
}

/**
 * Limit the magnitude of a vector.
 */
func (v1 PVector) limit (mag float64) PVector {
    if (v1.mag() > mag) {
        fmt.Println("LIMIT")
        return v1.normalize().setMag(mag)
    } else {
        return v1
    }
}

/*
heading() — the 2D heading of a vector expressed as an angle

rotate() — rotate a 2D vector by an angle

lerp() — linear interpolate to another vector

dist() — the Euclidean distance between two vectors (considered as points)

angleBetween() — find the angle between two vectors

dot() — the dot product of two vectors

cross() — the cross product of two vectors (only relevant in three dimensions)

random2D() - make a random 2D vector

random3D() - make a random 3D vector
*/


/**
 * An object that can move.
 */
type Mover struct {
    location PVector
    velocity PVector
    acceleration PVector
    topspeed float64
    mass float64
}

/**
 * Move the object based on acceleration and velocity.
 */
func (m *Mover) move () {
    m.velocity = m.velocity.add(m.acceleration)
    m.velocity = m.velocity.limit(m.topspeed)
    m.location = m.location.add(m.velocity)
    m.acceleration = m.acceleration.mult(0)
}

/**
 * Apply the given force on the mover.
 */
func (m *Mover) applyForce (f PVector) {
    f.div(m.mass)
    m.acceleration = m.acceleration.add(f)
}

/**
 * Move the object toward the given location.
func (m *Mover) moveToward (loc PVector) {
    m.velocity = m.velocity.add(m.acceleration).limit(m.topspeed)
    m.location = m.location.add(m.velocity)
}
 */

/**
 * The means of creating a Mover.
 */
func MoverFactory (location, velocity, acceleration PVector) Mover {
    m := Mover{
        location: location,
        velocity: velocity,
        acceleration: acceleration,
        topspeed: 2000,
        mass: 100,
    }
    return m
}

/**
 * A Perceptron.
 *
 * The weights are used to apply to each input. Since we are using a point with
 * x and y values, the third value is the bias. I haven't figured out how to
 * make weights an unbounded array.
 *
 * The learning constant determines how large the changes in guesses are, i.e.
 * learning velocity.
 */
type Perceptron struct {
    weights [3]float64
    learning float64
}

/**
 * This function adjusts each input's weight based on the error.
 */
func (p *Perceptron) train (input [3]float64, desired float64) {
    var guess float64 = p.feedforward(input)
    var error float64 = desired - guess
    var d float64 = error * p.learning
    for i := 0; i < len(p.weights); i++ {
        p.weights[i] = p.weights[i] + (input[i] * d)
    }

    if (guess == desired) {
        fmt.Printf("Correct! Weights are now: %v", p.weights)
    } else {
        fmt.Printf("Incorrect. Weights are now: %v", p.weights)
    }
    fmt.Println()
}

/**
 * Feedforward means: here are the inputs for the Perceptron, get the
 * Perceptron to tell us the value.
 */
func (p Perceptron) feedforward (input [3]float64) float64 {
    var sum float64 = 0

    for i := 0; i < len(input); i++ {
        sum += input[i] * p.weights[i]
    }

    return p.activate(sum)
}

/**
 * This method determines if the "neruon" should fire (1) or not fire (0).
 */
func (p Perceptron) activate (sum float64) float64 {
    if (sum > 0) {
        return 1;
    } else {
        return -1;
    }
}

/**
 * Create a Perceptron.
 *
 * For now, I'm using 3 inputs: x, y, and bias. When I figure out how to have
 * "n" number of arguments for the Perceptron I can actually start using the
 * "n" parameter.
 */
func PerceptronFactory (n int, learning float64) Perceptron {
    var weights [3]float64
    for i := 0; i < n; i++ {
        weights[i] = random(-1, 1)
    }
    p := Perceptron{
        weights: weights,
        learning: learning,
    }
    return p
}

/**
 * This is our line's definition.
 */
func f(x float64) float64 {
    return 2*x + 1
}

/**
 * A trainer has a point (the input) and the answer (whether or not it is above
 * the line).
 */
type Trainer struct {
    input [3]float64
    answer float64
}

/**
 * Create a Trainer based on a point.
 */
func TrainerFactory () Trainer {
    var xmin float64 = -400
    var xmax float64 = 400
    var ymin float64 = -100
    var ymax float64 = 100
    var answer float64
    var x float64 = random(xmin, xmax)
    var y float64 = random(ymin, ymax)

    if (y < f(x)) {
        answer = -1
    } else {
        answer = 1
    }

    var input = [3]float64{x, y, 1}
    t := Trainer{
        input: input,
        answer: answer,
    }

    return t
}

func main() {
    location := PVectorFactory(100, 100)
    velocity := PVectorFactory(0, 0)
    acceleration := PVectorFactory(0, 0)

    mover := MoverFactory(location, velocity, acceleration)

    wind := PVectorFactory(0.01, 0)
    gravity := PVectorFactory(0, -0.01)

    for i := 0; i < 10; i++ {
        mover.applyForce(wind)
        mover.applyForce(gravity.mult(mover.mass))

        fmt.Printf("%v", mover.location)
        fmt.Println()
        fmt.Println()
        mover.move()
    }
}

