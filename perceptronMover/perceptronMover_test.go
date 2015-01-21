package perceptronMover

import (
    "testing"
    "github.com/josephdpurcell/go-neural-network/pvector"
)

func TestPerceptronFactory(t *testing.T) {
    p := PerceptronFactory(2, 0.01)

    if len(p.weights) != 2 {
        t.Errorf("PerceptronFactory(%v, %v) == %v, want %v", 2, 0.01, len(p.weights), 2)
    }

    if p.learning != 0.01 {
        t.Errorf("PerceptronFactory(%v, %v) == %v, want %v", 2, 0.01, p.learning, 0.01)
    }
}

func TestPerceptronFeedforward(t *testing.T) {
    p := PerceptronFactory(2, 0.01)

    var input []pvector.PVector
    var got pvector.PVector
    var want pvector.PVector

    // With weights {{0,0}, {0,0}}, {{1,1}, {1,1}} should return {0,0}.
    p.weights = []pvector.PVector{pvector.PVector{0, 0}, pvector.PVector{0, 0}}
    input = []pvector.PVector{pvector.PVector{0,0}, pvector.PVector{0,0}}
    got = p.Feedforward(input)
    want = pvector.PVector{0, 0}
    if got.X != want.X || got.Y != want.Y {
        t.Errorf("p.Feedforward(%v) == %v, want %v", input, got, want)
    }

    // With weights {{1, 1}, {1, 1}}, {{1,1}, {1,1}} should return {2, 2}.
    p.weights = []pvector.PVector{pvector.PVector{1, 1}, pvector.PVector{1, 1}}
    input = []pvector.PVector{pvector.PVector{1, 1}, pvector.PVector{1, 1}}
    got = p.Feedforward(input)
    want = pvector.PVector{2, 2}
    if got.X != want.X || got.Y != want.Y {
        t.Errorf("p.Feedforward(%v) == %v, want %v", input, got, want)
    }

    // With weights {{0, 1}, {0, 1}}, {{1,1}, {1,1}} should return {0, 2}.
    p.weights = []pvector.PVector{pvector.PVector{0, 1}, pvector.PVector{0, 1}}
    input = []pvector.PVector{pvector.PVector{1, 1}, pvector.PVector{1, 1}}
    got = p.Feedforward(input)
    want = pvector.PVector{0, 2}
    if got.X != want.X || got.Y != want.Y {
        t.Errorf("p.Feedforward(%v) == %v, want %v", input, got, want)
    }

    // With weights {{1, 0}, {1, 0}}, {{1,1}, {1,1}} should return {2, 0}.
    p.weights = []pvector.PVector{pvector.PVector{1, 0}, pvector.PVector{1, 0}}
    input = []pvector.PVector{pvector.PVector{1, 1}, pvector.PVector{1, 1}}
    got = p.Feedforward(input)
    want = pvector.PVector{2, 0}
    if got.X != want.X || got.Y != want.Y {
        t.Errorf("p.Feedforward(%v) == %v, want %v", input, got, want)
    }
}

func TestPerceptronTrain(t *testing.T) {
    p := PerceptronFactory(2, 0.01)

    var input []pvector.PVector
    var error pvector.PVector
    var want []pvector.PVector

    // With weights {{0,0}, {0,0}}, {{1,1}, {1,1}} and error {0,0} should not modify weights.
    p.weights = []pvector.PVector{pvector.PVector{0, 0}, pvector.PVector{0, 0}}
    input = []pvector.PVector{pvector.PVector{1, 1}, pvector.PVector{1, 1}}
    error = pvector.PVector{0, 0}
    p.Train(input, error)
    want = []pvector.PVector{pvector.PVector{0, 0}, pvector.PVector{0, 0}}
    if p.weights[0] != want[0] || p.weights[1] != want[1] {
        t.Errorf("p.Train(%v, %v) ==> %v, want %v", input, p.weights, want)
    }

    // With weights {{1,1}, {1,1}}, {{1,1}, {1,1}} and error {1,1}, want {{1.01,1.01}, {1.01,1.01}}
    p.weights = []pvector.PVector{pvector.PVector{1, 1}, pvector.PVector{1, 1}}
    input = []pvector.PVector{pvector.PVector{1, 1}, pvector.PVector{1, 1}}
    error = pvector.PVector{1, 1}
    p.Train(input, error)
    want = []pvector.PVector{pvector.PVector{1.01, 1.01}, pvector.PVector{1.01, 1.01}}
    if p.weights[0] != want[0] || p.weights[1] != want[1] {
        t.Errorf("p.Train(%v, %v) ==> %v, want %v", input, p.weights, want)
    }

    // With weights {{1,1}, {1,1}}, {{1,1}, {1,1}} and error {0,1}, want {{1,1.01}, {1,1.01}}
    p.weights = []pvector.PVector{pvector.PVector{1, 1}, pvector.PVector{1, 1}}
    input = []pvector.PVector{pvector.PVector{1, 1}, pvector.PVector{1, 1}}
    error = pvector.PVector{0, 1}
    p.Train(input, error)
    want = []pvector.PVector{pvector.PVector{1, 1.01}, pvector.PVector{1, 1.01}}
    if p.weights[0] != want[0] || p.weights[1] != want[1] {
        t.Errorf("p.Train(%v, %v) ==> %v, want %v", input, p.weights, want)
    }
}

