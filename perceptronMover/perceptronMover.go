package perceptronMover

import (
    "github.com/josephdpurcell/go-neural-network/pvector"
)

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
    weights []pvector.PVector
    learning float64
}

/**
 * This function adjusts each input's weight based on the error.
 */
func (p *Perceptron) Train (forces []pvector.PVector, error pvector.PVector) {
    for i := 0; i < len(forces); i++ {
        p.weights[i].X = p.weights[i].X + (forces[i].X * error.X * p.learning)
        p.weights[i].Y = p.weights[i].Y + (forces[i].Y * error.Y * p.learning)
    }
}

/**
 * Feedforward means: here are the inputs for the Perceptron, get the
 * Perceptron to tell us the value.
 */
func (p Perceptron) Feedforward (forces []pvector.PVector) pvector.PVector {
    var sum pvector.PVector

    for i := 0; i < len(forces); i++ {
        // Vector addition and multiplication
        forces[i].X = forces[i].X * p.weights[i].X
        forces[i].Y = forces[i].Y * p.weights[i].Y
        sum = sum.Add(forces[i]);
    }

    // No activation function
    return sum
}

/**
 * Create a Perceptron.
 *
 * n = the number of elements in the vector
 * learning = the speed at which learning will happen
 */
func PerceptronFactory (n int, learning float64) Perceptron {
    weights := make([]pvector.PVector, n)
    for i := 0; i < n; i++ {
        weights[i] = pvector.PVectorFactory(1, 1)
    }
    p := Perceptron{
        weights: weights,
        learning: learning,
    }
    return p
}

