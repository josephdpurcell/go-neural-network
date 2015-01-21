package perceptronNAND

import "testing"

func TestPerceptronFactory(t *testing.T) {
    p := PerceptronFactory(2, 0.01)

    if len(p.weights) != 2 {
        t.Errorf("PerceptronFactory(%v, %v) == %v, want %v", 2, 0.01, len(p.weights), 2)
    }

    if p.learning != 0.01 {
        t.Errorf("PerceptronFactory(%v, %v) == %v, want %v", 2, 0.01, p.learning, 0.01)
    }
}

func TestPerceptronActivate(t *testing.T) {
    p := PerceptronFactory(2, 0.01)

    var got float64

    got = p.activate(0.51)

    if got != 1 {
        t.Errorf("p.activate(%v) == %v, want %v", 0.51, got, 1)
    }

    got = p.activate(0.49)

    if got != 0 {
        t.Errorf("p.activate(%v) == %v, want %v", 0.49, got, 0)
    }
}

func TestPerceptronFeedforward(t *testing.T) {
    p := PerceptronFactory(2, 0.01)

    p.weights[0] = 0.26
    p.weights[1] = 0.25

    input := []float64{1, 1}

    got := p.feedforward(input)

    if got != 1 {
        t.Errorf("p.feedforward(%v) == %v, want %v", input, got, 1)
    }

    p.weights[0] = 0.25
    p.weights[1] = 0.25

    input = []float64{1, 1}

    got = p.feedforward(input)

    if got != 0 {
        t.Errorf("p.feedforward(%v) == %v, want %v", input, got, 0)
    }
}

func TestPerceptronTrain(t *testing.T) {
    p := PerceptronFactory(2, 0.01)

    var desired float64
    var input []float64

    // With weights {0, 0}, {1, 1} should not modify weights.
    p.weights[0] = 0
    p.weights[1] = 0
    input = []float64{1, 1}
    desired = 0

    p.Train(input, desired)

    if p.weights[0] != 0 || p.weights[1] != 0 {
        t.Errorf("Weights should have remained the same, but are: %v", p.weights)
    }

    // With weights {1, 1}, {1, 1} should modify weights to {0.99, 0.99}
    p.weights[0] = 1
    p.weights[1] = 1
    input = []float64{1, 1}
    desired = 0

    p.Train(input, desired)

    if p.weights[0] != 0.99 || p.weights[1] != 0.99 {
        t.Errorf("Weights should have remained the same, but are: %v", p.weights)
    }

    // With weights {0, 0}, {0, 0} should not modify weights.
    p.weights[0] = 0
    p.weights[1] = 0
    input = []float64{0, 0}
    desired = 1

    p.Train(input, desired)

    if p.weights[0] != 0 || p.weights[1] != 0 {
        t.Errorf("Weights should have remained the same, but are: %v", p.weights)
    }

    // With weights {1, 1}, {0, 0} should modify weights to {1, 1}.
    p.weights[0] = 1
    p.weights[1] = 1
    input = []float64{0, 0}
    desired = 1

    p.Train(input, desired)

    if p.weights[0] != 1 || p.weights[1] != 1 {
        t.Errorf("Weights should have remained the same, but are: %v", p.weights)
    }

    // With weights {0, 0}, {0, 1} should modify weights to {0, 0.01}.
    p.weights[0] = 0
    p.weights[1] = 0
    input = []float64{0, 1}
    desired = 1

    p.Train(input, desired)

    if p.weights[0] != 0 || p.weights[1] != 0.01 {
        t.Errorf("Weights should have remained the same, but are: %v", p.weights)
    }

    // With weights {0, 0}, {1, 0} should modify weights to {0.01, 0}.
    p.weights[0] = 0
    p.weights[1] = 0
    input = []float64{1, 0}
    desired = 1

    p.Train(input, desired)

    if p.weights[0] != 0.01 || p.weights[1] != 0 {
        t.Errorf("Weights should have remained the same, but are: %v", p.weights)
    }
}
