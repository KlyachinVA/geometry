package geometry

import (

"math"

)
type Point struct {
X float64
Y float64
Z float64
}

func Dot(a Point,b Point) float64{
	return a.X*b.X + a.Y*b.Y+a.Z*b.Z
	}
	
func Cross(a Point,b Point)	Point{
	var c = new(Point)
	c.X = a.Y*b.Z - a.Z*b.Y
	c.Y = a.X*b.Z - a.Z*b.X
	c.Z = a.X*b.Y - a.Y*b.X
	return *c
	}

func (a * Point ) Norm() float64{
	return math.Sqrt(Dot(*a,*a))
	}

