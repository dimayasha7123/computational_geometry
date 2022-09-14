package lab1

import (
	"fmt"
	"math"
)

const (
	EPS = 1e-9
)

type Dot struct {
	X float64
	Y float64
}

type Segment struct {
	A Dot
	B Dot
}

func LengthbetweenDots(a, b Dot) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2))
}

func (s *Segment) Length() float64 {
	return LengthbetweenDots(s.A, s.B)
}

func RadToDegree(angle float64) float64 {
	return angle * 180 / math.Pi
}

func DegreeToRad(angle float64) float64 {
	return angle * math.Pi / 180
}

func OXAngle(dot Dot) float64 {
	seg := Segment{A: Dot{0, 0}, B: dot}
	cos := dot.X / seg.Length()
	acos := math.Acos(cos)
	deg := RadToDegree(acos)
	if dot.Y < 0 {
		return 360 - deg
	}
	return deg
}

// for task 1
func WhatAngleWithOxMore(a, b Dot) string {
	aAngle := OXAngle(a)
	bAngle := OXAngle(b)
	switch {
	case aAngle > bAngle:
		return fmt.Sprintf("a angle (%f) is bigger than b angle (%f)", aAngle, bAngle)
	case aAngle < bAngle:
		return fmt.Sprintf("b angle (%f) is bigger than a angle (%f)", bAngle, aAngle)
	default:
		return fmt.Sprintf("Angles are equal (%f)", aAngle)
	}
}

// for task2
func CheckDotOnSegment(a Dot, s Segment) bool {
	s1 := LengthbetweenDots(a, s.A)
	s2 := LengthbetweenDots(a, s.B)
	return math.Abs(s1+s2-s.Length()) < EPS
}

func GetAngleABC(a, b, c Dot) float64 {
	as := Segment{b, c}
	bs := Segment{a, c}
	cs := Segment{a, b}
	al := as.Length()
	bl := bs.Length()
	cl := cs.Length()
	cos := (al*al + cl*cl - bl*bl) / (2 * al * cl)
	acos := math.Acos(cos)
	deg := RadToDegree(acos)
	return deg
}

// task3
func NormOnSegment(a Dot, s Segment) bool {
	a1 := GetAngleABC(a, s.A, s.B)
	a2 := GetAngleABC(a, s.B, s.A)

	return (a1 < 90 && a2 < 90) ||
		(a1 < 90 && math.Abs(a2-90) < EPS) ||
		(a2 < 90 && math.Abs(a1-90) < EPS)
}

// line Ax + by + c = 0
type Line struct {
	A float64
	B float64
	C float64
}

// get Line from y = ax + b
func FromAngularForm(a, b float64) *Line {
	return &Line{A: a, B: -1, C: b}
}

// get line from segment (line, that parallel this segment)
func FromSegment(s Segment) *Line {
	return &Line{
		A: s.A.Y - s.B.Y,
		B: s.B.X - s.A.X,
		C: s.A.X*s.B.Y - s.B.X*s.A.Y,
	}
}

// determinant of this
// | a b |
// | c d |
func det(a, b, c, d float64) float64 {
	return a*d - b*c
}

// check if two lines are parallel
func IsParallel(l1, l2 Line) bool {
	return math.Abs(det(l1.A, l1.B, l2.A, l2.B)) < EPS
}

// check intersection of line with segment
func LineSegmentItersection(l Line, s Segment) bool {
	// l and s must be not parallel
	l1 := l
	l2 := *FromSegment(s)
	if IsParallel(l1, l2) {
		return false
	}

	// find intersection dot
	u := det(l1.A, l1.B, l2.A, l2.B)
	d := Dot{
		X: -det(l1.C, l1.B, l2.C, l2.B) / u,
		Y: -det(l1.A, l1.C, l2.A, l2.C) / u,
	}

	// dot must be on segment
	return CheckDotOnSegment(d, s)
}
