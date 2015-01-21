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
    "github.com/josephdpurcell/go-neural-network/perceptronNAND"
)

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
    input []float64
    answer float64
}

func main() {
    // Setup the trainers.
    const count int = 180
    var trainers [count]Trainer
    var key int
    for i := 0; i < (count / 4); i++ {
        key = (i * 4) + 0
        trainers[key] = Trainer{
            input: []float64{1, 0, 0},
            answer: 1,
        }

        key = (i * 4) + 1
        trainers[key] = Trainer{
            input: []float64{1, 0, 1},
            answer: 1,
        }

        key = (i * 4) + 2
        trainers[key] = Trainer{
            input: []float64{1, 1, 0},
            answer: 1,
        }

        key = (i * 4) + 3
        trainers[key] = Trainer{
            input: []float64{1, 1, 1},
            answer: 0,
        }
    }

    // Learning Constant is low just b/c it's fun to watch, this is not necessarily optimal
    p := perceptronNAND.PerceptronFactory(3, 0.1)

    // Train our Perceptron.
    for i := 0; i < len(trainers); i++ {
        fmt.Printf("%v: ", i)
        p.Train(trainers[i].input, trainers[i].answer)
    }
}

