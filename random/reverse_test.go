package random

import "testing"

func TestRandom(t *testing.T) {
    got := Random(10, 20)
    if got < 10 {
        t.Errorf("Random(10, 20) == %v, want >= 10", got)
    }
    if got > 20 {
        t.Errorf("Random(10, 20) == %v, want <= 20", got)
    }
}
