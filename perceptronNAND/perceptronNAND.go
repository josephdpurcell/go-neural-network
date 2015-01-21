package perceptronNAND

import "fmt"

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
func (p *Perceptron) Train (input []float64, desired float64) {
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
func (p Perceptron) feedforward (input []float64) float64 {
    var sum float64 = 0

    for i := 0; i < len(input); i++ {
        sum = sum + (input[i] * p.weights[i])
    }

    return p.activate(sum)
}

/**
 * This method determines if the "neruon" should fire (1) or not fire (0).
 */
func (p Perceptron) activate (sum float64) float64 {
    if (sum > 0.5) {
        return 1;
    } else {
        return 0;
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
    weights := make([]float64, n)
    for i := 0; i < n; i++ {
        weights[i] = 0
    }
    p := Perceptron{
        weights: weights,
        learning: learning,
    }
    return p
}

