/**
 * A Neural Network Perceptron Example Learning f(x).
 *
 * This is an example of creating a rudimentary neural network to determine if
 * a given point is above or below a line described by f(x).
 *
 * Source: http://natureofcode.com/book/chapter-10-neural-networks/
 */
package main

import (
    "fmt"
    "math/rand"
)

/**
 * Generate a random float64 between max and min.
 */
func random(min, max float64) float64 {
    return (rand.Float64() * (max - min)) + min
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
    // Setup the trainers.
    const count int = 100000
    var trainers [count]Trainer
    for i := 0; i < len(trainers); i++ {
        trainers[i] = TrainerFactory()
    }

    // Learning Constant is low b/c it's fun to watch, not necessarily for performance.
    p := PerceptronFactory(3, 0.00001)

    // Train our Perceptron.
    for i := 0; i < len(trainers); i++ {
        fmt.Printf("%v: ", i)
        p.train(trainers[i].input, trainers[i].answer)
    }
}

