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
    brain Perceptron
    location PVector
    velocity PVector
    acceleration PVector
    maxspeed float64
    maxforce float64
    mass float64
}

/**
 * Move the object based on acceleration and velocity.
 */
func (m *Mover) update () {
    m.velocity = m.velocity.add(m.acceleration)
    m.velocity = m.velocity.limit(m.maxspeed)
    m.location = m.location.add(m.velocity)
    m.acceleration = m.acceleration.mult(0)
}

/**
 * Apply the given force on the mover.
 */
func (m *Mover) applyForce (f PVector) {
    f = f.div(m.mass)
    m.acceleration = m.acceleration.add(f)
}

/**
 * Get the steering force toward a target.
 */
func (m Mover) getSteeringForce (target PVector) PVector {
    var desired PVector
    var steer PVector
    var mag float64

    // Compute the desired velocity.
    // @todo we don't want to normalize and scale by the magnitude in all
    // directions so we can handle the case where x forces > y forces
    desired = target.sub(m.location)
    mag = desired.mag()
    desired = desired.normalize();
    // Note: we are saying when the remaining distance is within twice our
    // max-speed, begin to slow down.
    if (mag < (m.maxspeed * 2)) {
        desired = desired.mult(mag / 2);
    } else {
        desired = desired.mult(m.maxspeed);
    }

    // Compute the desired steering force.
    steer = desired.sub(m.velocity)
    steer = steer.sub(m.acceleration)
    steer = steer.mult(m.mass)
    steer = steer.limit(m.maxforce)

    return steer
}

/**
 * Move the object toward the given targets.
 */
func (m *Mover) seek (targets []PVector) {
    // Gather forces.
    var forces = make([]PVector, len(targets))
    for i := 0; i < len(targets); i++ {
        forces[i] = m.getSteeringForce(targets[i])
    }

    // Compute the steering force and apply.
    output := m.brain.feedforward(forces)
    m.applyForce(output)

    // Train the brain based on the error.
    desired := PVectorFactory(400, 400)
    fmt.Printf("DESIRED: %v", desired)
    fmt.Println()
    error := desired.sub(m.location)
    fmt.Printf("ERROR: %v", error)
    fmt.Println()
    m.brain.train(forces, error)
}

/**
 * The means of creating a Mover.
 */
func MoverFactory (location, velocity, acceleration PVector) Mover {
    m := Mover{
        brain: PerceptronFactory(2, 0.01),
        location: location,
        velocity: velocity,
        acceleration: acceleration,
        maxspeed: 20,
        maxforce: 2000,
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
    weights []float64
    learning float64
}

/**
 * This function adjusts each input's weight based on the error.
 */
func (p *Perceptron) train (forces []PVector, error PVector) {
    for i := 0; i < len(p.weights); i++ {
        p.weights[i] = p.weights[i] + (forces[i].x * error.x * p.learning)
        p.weights[i] = p.weights[i] + (forces[i].y * error.y * p.learning)
    }
    fmt.Printf("%v", p.weights)
    fmt.Println()
}

/**
 * Feedforward means: here are the inputs for the Perceptron, get the
 * Perceptron to tell us the value.
 */
func (p Perceptron) feedforward (forces []PVector) PVector {
    var sum PVector

    for i := 0; i < len(p.weights); i++ {
        // Vector addition and multiplication
        forces[i] = forces[i].mult(p.weights[i]);
        sum = sum.add(forces[i]);
    }

    // No activation function
    return sum
}

/**
 * Create a Perceptron.
 */
func PerceptronFactory (n int, learning float64) Perceptron {
    weights := make([]float64, n)
    for i := 0; i < n; i++ {
        weights[i] = random(-1, 1)
    }
    p := Perceptron{
        weights: weights,
        learning: learning,
    }
    return p
}

func main() {
    // Create our mover.
    location := PVectorFactory(100, 100)
    velocity := PVectorFactory(0, 0)
    acceleration := PVectorFactory(0, 0)
    mover := MoverFactory(location, velocity, acceleration)

    // Target we are seeking.
    targets := []PVector{PVectorFactory(200, 200), PVectorFactory(800, 800)}

    // External forces.
    //wind := PVectorFactory(1000, 0)
    //gravity := PVectorFactory(0, -0.01 * mover.mass)
    var friction PVector
    var c float64 = 0.01

    // Iterate over time.
    for t := 0; t < 100; t++ {
        // Compute the friction.
        friction = mover.velocity
        friction = friction.mult(-1)
        friction = friction.normalize()
        friction = friction.mult(c)

        // Apply external forces.
        //mover.applyForce(friction)
        //mover.applyForce(wind)
        //mover.applyForce(gravity)

        // Seek the target.
        mover.seek(targets)

        // Update and display the result.
        mover.update()
        fmt.Printf("LOC: %v", mover.location)
        fmt.Println()
        fmt.Println()
    }
}

