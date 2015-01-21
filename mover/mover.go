package mover

import (
    "fmt"
    "github.com/josephdpurcell/go-neural-network/pvector"
    "github.com/josephdpurcell/go-neural-network/perceptronMover"
)

/**
 * An object that can move.
 */
type Mover struct {
    brain perceptronMover.Perceptron
    location pvector.PVector
    velocity pvector.PVector
    acceleration pvector.PVector
    maxspeed float64
    maxforce float64
    mass float64
}

/**
 * Move the object based on acceleration and velocity.
 *
 * With every iteration, we add the acceleration to the velocity, limit the
 * velocity to its max speed, then move the Mover the distance based on
 * velocity. Finally, we set acceleration to 0 to allow re-computation of the
 * steering force.
 */
func (m *Mover) Update () {
    m.velocity = m.velocity.Add(m.acceleration)
    m.velocity = m.velocity.Limit(m.maxspeed)
    m.location = m.location.Add(m.velocity)
    m.acceleration = m.acceleration.Mult(0)
}

/**
 * Apply the given force on the mover.
 */
func (m *Mover) applyForce (f pvector.PVector) {
    f = f.Div(m.mass)
    m.acceleration = m.acceleration.Add(f)
}

/**
 * Get the steering force toward a target.
 */
func (m Mover) getSteeringForce (target pvector.PVector) pvector.PVector {
    var desired pvector.PVector
    var steer pvector.PVector
    var mag float64

    // Compute the desired velocity.
    // @todo we don't want to normalize and scale by the magnitude in all
    // directions so we can handle the case where x forces > y forces
    desired = target.Sub(m.location)
    mag = desired.Mag()
    desired = desired.Normalize();
    // Note: we are saying when the remaining distance is within twice our
    // max-speed, begin to slow down.
    if (mag < (m.maxspeed * 2)) {
        desired = desired.Mult(mag / 2);
    } else {
        desired = desired.Mult(m.maxspeed);
    }

    // Compute the desired steering force.
    steer = desired.Sub(m.velocity)
    //steer = steer.Sub(m.acceleration)
    steer = steer.Mult(m.mass)
    fmt.Printf("STEER: %v", steer)
    fmt.Println()
    steer = steer.Limit(m.maxforce)
    fmt.Printf("STEER: %v", steer.Div(m.mass))
    fmt.Println()

    return steer
}

/**
 * Move the object toward the given targets.
 */
func (m *Mover) Seek (targets []pvector.PVector) {
    // Gather forces.
    var forces = make([]pvector.PVector, len(targets))
    for i := 0; i < len(targets); i++ {
        forces[i] = m.getSteeringForce(targets[i])
    }

    // Compute the steering force and apply.
    output := m.brain.Feedforward(forces)
    m.applyForce(output)

    // Train the brain to go towards a specific one.
    desired := pvector.PVectorFactory(209, 215)
    error := desired.Sub(m.location)
    fmt.Printf("LOC: %v", m.location)
    fmt.Println()
    fmt.Printf("DES: %v", desired)
    fmt.Println()
    fmt.Printf("ERROR: %v", error)
    fmt.Println()
    m.brain.Train(forces, error)
}

/**
 * The means of creating a Mover.
 */
func MoverFactory (location, velocity, acceleration pvector.PVector) Mover {
    m := Mover{
        brain: perceptronMover.PerceptronFactory(2, 0.00001),
        location: location,
        velocity: velocity,
        acceleration: acceleration,
        maxspeed: 20,
        maxforce: 2000,
        mass: 100,
    }
    return m
}

