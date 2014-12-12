/**
 * A Neural Network Perceptron Example Learning NAND.
 *
 * This is an example of creating a rudimentary neural network to determine if
 * given the boolean values x_1 and x_2 if it passes NAND.
 *
 * Source(s):
 *   - http://natureofcode.com/book/chapter-10-neural-networks/
 *   - http://en.wikipedia.org/wiki/Perceptron
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
 * Generate a random float64 "boolean" 0 or 1.
 */
func randomBool() float64 {
    var random float64 = rand.Float64()
    if (random >= 0.5) {
        return 1
    } else {
        return 0
    }
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
        fmt.Println("Correct!")
    } else {
        fmt.Println("Incorrect.")
    }
}

/**
 * Feedforward means: here are the inputs for the Perceptron, get the
 * Perceptron to tell us the value.
 */
func (p Perceptron) feedforward (input [3]float64) float64 {
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
    var weights [3]float64
    for i := 0; i < n; i++ {
        weights[i] = 0
    }
    p := Perceptron{
        weights: weights,
        learning: learning,
    }
    return p
}

/**
 * This is the definition of nand.
 */
func nand(x1, x2 float64) bool {
    return (x1 == 1 || x2 == 1) && (x1 != 1 || x2 != 1)
}

/**
 * A trainer has a point (the input) and the answer (nand).
 */
type Trainer struct {
    input [3]float64
    answer float64
}

/**
 * Create a Trainer based on x1 and x2
 */
func TrainerFactory () Trainer {
    var answer float64
    var x1 float64 = randomBool()
    var x2 float64 = randomBool()

    if (nand(x1, x2)) {
        answer = -1
    } else {
        answer = 1
    }

    var input = [3]float64{x1, x2, 1}
    t := Trainer{
        input: input,
        answer: answer,
    }

    return t
}

func main() {
    // Setup the trainers.
    const count int = 180
    var trainers [count]Trainer
    var key int
    for i := 0; i < (count / 4); i++ {
        key = (i * 4) + 0
        trainers[key] = Trainer{
            input: [3]float64{1, 0, 0},
            answer: 1,
        }


        key = (i * 4) + 1
        trainers[key] = Trainer{
            input: [3]float64{1, 0, 1},
            answer: 1,
        }

        key = (i * 4) + 2
        trainers[key] = Trainer{
            input: [3]float64{1, 1, 0},
            answer: 1,
        }

        key = (i * 4) + 3
        trainers[key] = Trainer{
            input: [3]float64{1, 1, 1},
            answer: 0,
        }
    }

    // Learning Constant is low just b/c it's fun to watch, this is not necessarily optimal
    p := PerceptronFactory(3, 0.1)

    // Train our Perceptron.
    for i := 0; i < len(trainers); i++ {
        p.train(trainers[i].input, trainers[i].answer)
    }
}

