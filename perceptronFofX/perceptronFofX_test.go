package perceptronFofX

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

    // 0 should return -1
    got = p.activate(0)
    if got != -1 {
        t.Errorf("p.activate(%v) == %v, want %v", 0, got, -1)
    }

    // -0.01 should return -1
    got = p.activate(-0.01)
    if got != -1 {
        t.Errorf("p.activate(%v) == %v, want %v", -0.01, got, -1)
    }

    // 0.01 should return 1
    got = p.activate(0.01)
    if got != 1 {
        t.Errorf("p.activate(%v) == %v, want %v", 0.01, got, 1)
    }
}

func TestPerceptronFeedforward(t *testing.T) {
    p := PerceptronFactory(2, 0.01)

    var input []float64
    var got float64

    // With weights {0, 0}, {1,1} should return -1.
    p.weights[0] = 0
    p.weights[1] = 0
    input = []float64{1, 1}
    got = p.feedforward(input)
    if got != -1 {
        t.Errorf("p.feedforward(%v) == %v, want %v", input, got, -1)
    }

    // With weights {0.1, 0.1}, {1,1} should return 1.
    p.weights[0] = 0.1
    p.weights[1] = 0.1
    input = []float64{1, 1}
    got = p.feedforward(input)
    if got != 1 {
        t.Errorf("p.feedforward(%v) == %v, want %v", input, got, 1)
    }

    // With weights {-0.1, -0.1}, {1,1} should return -1.
    p.weights[0] = -0.1
    p.weights[1] = -0.1
    input = []float64{1, 1}
    got = p.feedforward(input)
    if got != -1 {
        t.Errorf("p.feedforward(%v) == %v, want %v", input, got, -1)
    }

    // With weights {-0.2, 0.1}, {1,1} should return -1.
    p.weights[0] = -0.2
    p.weights[1] = 0.1
    input = []float64{1, 1}
    got = p.feedforward(input)
    if got != -1 {
        t.Errorf("p.feedforward(%v) == %v, want %v", input, got, -1)
    }
}

func TestPerceptronTrain(t *testing.T) {
    p := PerceptronFactory(2, 0.01)

    var desired float64
    var input []float64

    // Assume f(x) = 2 * x + 1

    // With weights {1, 1}, {3, 0} = 1 should not modify weights.
    p.weights[0] = 1
    p.weights[1] = 1
    input = []float64{3, 0}
    desired = 1
    p.Train(input, desired)
    if p.weights[0] != 1 || p.weights[1] != 1 {
        t.Errorf("Weights should have remained the same, but are: %v", p.weights)
    }

    // With weights {1, 1}, {-3, 0} = -1 should not modify weights.
    p.weights[0] = 1
    p.weights[1] = 1
    input = []float64{-3, 0}
    desired = -1
    p.Train(input, desired)
    if p.weights[0] != 1 || p.weights[1] != 1 {
        t.Errorf("Weights should have remained the same, but are: %v", p.weights)
    }

    // With weights {0, 0}, {3, 0} = 1 should set the weights to {-0.06, 0}
    p.weights[0] = 0
    p.weights[1] = 0
    input = []float64{-3, 0}
    desired = 1
    p.Train(input, desired)
    if p.weights[0] != -0.06 || p.weights[1] != 0 {
        t.Errorf("Weights are wrong. They are: %v", p.weights)
    }

    // With weights {1, 1}, {-3, 0} = -1 should set the weights to {0.94, 1}
    p.weights[0] = 1
    p.weights[1] = 1
    input = []float64{3, 0}
    desired = -1
    p.Train(input, desired)
    if p.weights[0] != 0.94 || p.weights[1] != 1 {
        t.Errorf("Weights are wrong. They are: %v", p.weights)
    }
}

