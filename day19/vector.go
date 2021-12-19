package day19

import (
    "fmt"
    "math"
)

type Vector struct {
    X int
    Y int
    Z int
}

func (v Vector) Sub(s Vector) Vector {
    return Vector{
        X: v.X - s.X,
        Y: v.Y - s.Y,
        Z: v.Z - s.Z,
    }
}

func (v Vector) Add(s Vector) Vector {
    return Vector{
        X: v.X + s.X,
        Y: v.Y + s.Y,
        Z: v.Z + s.Z,
    }
}

func (v Vector) Mul(f int) Vector {
    return Vector{
        X: v.X * f,
        Y: v.Y * f,
        Z: v.Z * f,
    }
}

func (v Vector) Len() float64 {
    x := float64(v.X)
    y := float64(v.Y)
    z := float64(v.Z)
    return math.Sqrt(x*x + y*y + z*z)
}

func (v Vector) String() string {
    return fmt.Sprintf("%d,%d,%d", v.X, v.Y, v.Z)
}

func (v Vector) Compare(s Vector) bool {
    return v.Equal(s) || v.Mul(-1).Equal(s)
}

func (v Vector) Equal(s Vector) bool {
    return s.X == v.X && s.Y == v.Y && s.Z == v.Z
}

func (v Vector) Rotations() []Vector {
    x := v.X
    y := v.Y
    z := v.Z
    return []Vector{
        {x, y, z},
        {x, z, -y},
        {x, -y, -z},
        {x, -z, y},
        {y, -z, -x},
        {y, -x, z},
        {y, z, x},
        {y, x, -z},
        {z, x, y},
        {z, -y, x},
        {z, -x, -y},
        {z, y, -x},
        {-x, y, -z},
        {-x, z, y},
        {-x, -y, z},
        {-x, -z, -y},
        {-y, z, -x},
        {-y, x, z},
        {-y, -z, x},
        {-y, -x, -z},
        {-z, -x, y},
        {-z, -y, -x},
        {-z, x, -y},
        {-z, y, x},
    }
}
