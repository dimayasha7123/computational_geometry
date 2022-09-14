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

func LengthBetweenDots(A, B Dot) float64 {
	return math.Sqrt(math.Pow(A.X-B.X, 2) + math.Pow(A.Y-B.Y, 2))
}

func (s *Segment) Length() float64 {
	return LengthBetweenDots(s.A, s.B)
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
func WhatAngleWithOxMore(A, B Dot) string {
	aAngle := OXAngle(A)
	bAngle := OXAngle(B)
	switch {
	case aAngle > bAngle:
		return fmt.Sprintf("A angle (%f) is bigger than B angle (%f)", aAngle, bAngle)
	case aAngle < bAngle:
		return fmt.Sprintf("B angle (%f) is bigger than A angle (%f)", bAngle, aAngle)
	default:
		return fmt.Sprintf("Angles are equal (%f)", aAngle)
	}
}

// for task2
func CheckDotOnSegment(A Dot, Seg Segment) bool {
	s1 := LengthBetweenDots(A, Seg.A)
	s2 := LengthBetweenDots(A, Seg.B)
	return math.Abs(s1+s2-Seg.Length()) < EPS
}

func GetAngleABC(A, B, C Dot) float64 {
	as := Segment{B, C}
	bs := Segment{A, C}
	cs := Segment{A, B}
	a := as.Length()
	b := bs.Length()
	c := cs.Length()
	cos := (a*a + c*c - b*b) / (2 * a * c)
	acos := math.Acos(cos)
	deg := RadToDegree(acos)
	return deg
}

// task3
func NormOnSegment(A Dot, Seg Segment) bool {
	a1 := GetAngleABC(A, Seg.A, Seg.B)
	a2 := GetAngleABC(A, Seg.B, Seg.A)

	return (a1 < 90 && a2 < 90) ||
		(a1 < 90 && math.Abs(a2-90) < EPS) ||
		(a2 < 90 && math.Abs(a1-90) < EPS)
}
