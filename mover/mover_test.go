package mover

import (
    "testing"
    "github.com/josephdpurcell/go-neural-network/pvector"
)

func TestMoverFactory(t *testing.T) {
    location := pvector.PVectorFactory(100, 100)
    velocity := pvector.PVectorFactory(0, 0)
    acceleration := pvector.PVectorFactory(0, 0)

    m := MoverFactory(location, velocity, acceleration)

    if m.location != location {
        t.Errorf("Location should have been %v, but was %v", location, m.location)
    }

    if m.velocity != velocity {
        t.Errorf("Velocity should have been %v, but was %v", velocity, m.velocity)
    }

    if m.acceleration != acceleration {
        t.Errorf("Acceleration should have been %v, but was %v", acceleration, m.acceleration)
    }
}

