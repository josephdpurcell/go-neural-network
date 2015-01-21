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
    "github.com/josephdpurcell/go-neural-network/mover"
    "github.com/josephdpurcell/go-neural-network/pvector"
)

func main() {
    // Create our mover.
    location := pvector.PVectorFactory(100, 100)
    velocity := pvector.PVectorFactory(0, 0)
    acceleration := pvector.PVectorFactory(0, 0)
    mover := mover.MoverFactory(location, velocity, acceleration)

    // Target we are seeking.
    targets := []pvector.PVector{pvector.PVectorFactory(209, 215), pvector.PVectorFactory(400, 400)}

    // External forces.
    //wind := pvector.PVectorFactory(1000, 0)
    //gravity := pvector.PVectorFactory(0, -0.01 * mover.mass)
    //var friction pvector.PVector
    //var c float64 = 0.01

    fmt.Printf("STARTING LOC: %v", location)
    fmt.Println()
    fmt.Println()

    // Iterate over time.
    for t := 0; t < 1000; t++ {
        /*
        // Compute the friction.
        friction = mover.velocity
        friction = friction.mult(-1)
        friction = friction.normalize()
        friction = friction.mult(c)
        */

        // Apply external forces.
        //mover.applyForce(friction)
        //mover.applyForce(wind)
        //mover.applyForce(gravity)

        // Seek the target.
        mover.Seek(targets)

        // Update and display the result.
        mover.Update()
    }
}

