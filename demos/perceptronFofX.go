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
    "github.com/josephdpurcell/go-neural-network/random"
    "github.com/josephdpurcell/go-neural-network/perceptronFofX"
)

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
    input []float64
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
    var x float64 = random.Random(xmin, xmax)
    var y float64 = random.Random(ymin, ymax)

    if (y < f(x)) {
        answer = -1
    } else {
        answer = 1
    }

    var input = []float64{x, y, 1}
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
    p := perceptronFofX.PerceptronFactory(3, 0.00001)

    // Train our Perceptron.
    for i := 0; i < len(trainers); i++ {
        fmt.Printf("%v: ", i)
        p.Train(trainers[i].input, trainers[i].answer)
    }
}


