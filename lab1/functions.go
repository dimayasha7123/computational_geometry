package lab1

import (
	"fmt"
	"math"
	"strings"
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

func CheckDotOnSegmentInteger(a Dot, s Segment) bool {
	if (s.A.X-a.X)*(s.B.Y-a.Y)+(s.B.X-a.X)*(s.A.Y-a.Y) != 0 {
		return false
	}
	return (s.A.X-a.X)/(s.B.X-a.X) < 0
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

func (l *Line) Distanse(d Dot) float64 {
	return (math.Abs(l.A*d.X + l.B*d.Y + l.C)) / (math.Sqrt(d.X*d.X + d.Y*d.Y))
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
func LineSegmentIntersection(l Line, s Segment) bool {
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

// check intersection of two segments
func SegmentsIntersection(s1, s2 Segment) bool {
	l1 := *FromSegment(s1)
	l2 := *FromSegment(s2)

	// check if parallel
	if IsParallel(l1, l2) {
		// check, if equal or laying on one line and intersects
		return !(math.Abs(l1.Distanse(s2.A)) > EPS || math.Abs(l2.Distanse(s1.A)) > EPS)
	}
	// if not parallel

	// find intersection dot
	u := det(l1.A, l1.B, l2.A, l2.B)
	d := Dot{
		X: -det(l1.C, l1.B, l2.C, l2.B) / u,
		Y: -det(l1.A, l1.C, l2.A, l2.C) / u,
	}

	// dot must be on segments
	return CheckDotOnSegment(d, s1) && CheckDotOnSegment(d, s2)
}

// signed area of triangle
func SignedDoubleTriangleArea(a, b, c Dot) float64 {
	return (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)
}

// it will work right if with integers...

// func IsClockwise(a, b, c Dot) bool {
// 	return TriangleArea(a, b, c) < 0
// }

//          B
//        |
//       |
//      |
//     |
//    |
//   |____________ A
//
// if triangle area > 0, than angle from A to B (counterclockwise), else from B to A (clockwise)
//

type Position int

const (
	Inside Position = iota
	OnBorder
	Outside
)

type Triangle struct {
	A Dot
	B Dot
	C Dot
}

func DotAndTriangle(d Dot, t Triangle) Position {
	ss := []Segment{{t.A, t.B}, {t.B, t.C}, {t.C, t.A}}
	wasOnBorder := false

	for _, s := range ss {
		sdta := SignedDoubleTriangleArea(d, s.A, s.B)
		switch {
		case math.Abs(sdta) < EPS:
			wasOnBorder = true
			continue
		case sdta < 0:
			return Outside
		}
	}

	if wasOnBorder {
		return OnBorder
	}
	return Inside
}

// not works if len(segments) < 2
func MaxIntersectionLine(segments []Segment) (Segment, []int) {
	dots := make([]Dot, len(segments)*2)
	for i, s := range segments {
		dots[2*i] = s.A
		dots[2*i+1] = s.B
	}

	var bestA, bestB Dot
	var bestSegments []int
	first := true

	for i := 0; i < len(dots); i++ {
		j := i + 1
		if j%2 == 1 {
			j++
		}
		for ; j < len(dots); j++ {
			if dots[i] == dots[j] {
				continue
			}
			line := FromSegment(Segment{dots[i], dots[j]})
			maybeSegments := make([]int, 0)
			for k, s := range segments {
				if LineSegmentIntersection(*line, s) {
					maybeSegments = append(maybeSegments, k)
				}
			}
			if first || len(maybeSegments) > len(bestSegments) {
				first = false
				bestA = dots[i]
				bestB = dots[j]
				bestSegments = maybeSegments
			}
		}
	}

	return Segment{bestA, bestB}, bestSegments
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CheckSymmetry(matrix [][]int) []string {
	n := len(matrix)
	m := len(matrix[0])

	left := m - 1
	right := 0
	up := n - 1
	down := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] != 0 {
				left = min(left, j)
				break
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := m - 1; j >= 0; j-- {
			if matrix[i][j] != 0 {
				right = max(right, j)
				break
			}
		}
	}
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			if matrix[i][j] != 0 {
				up = min(up, i)
			}
		}
	}
	for j := 0; j < m; j++ {
		for i := n - 1; i >= 0; i-- {
			if matrix[i][j] != 0 {
				down = max(down, i)
			}
		}
	}

	wide := right - left + 1
	height := down - up + 1

	cropped := make([][]int, height)
	for i := 0; i < height; i++ {
		cropped[i] = make([]int, wide)
		for j := 0; j < wide; j++ {
			cropped[i][j] = matrix[i+up][j+left]
		}
	}

	fmt.Println(strings.ReplaceAll(fmt.Sprint(cropped), "] ", "]\n "))

	symms := make([]string, 0)

	check := true
	for i := 0; i < height; i++ {
		for j := 0; j <= wide/2; j++ {
			if cropped[i][j] != cropped[i][wide-1-j] {
				check = false
				break
			}
		}
		if !check {
			break
		}
	}
	if check {
		symms = append(symms, "vertical")
	}

	check = true
	for j := 0; j < wide; j++ {
		for i := 0; i <= height/2; i++ {
			if cropped[i][j] != cropped[height-1-i][j] {
				check = false
				break
			}
		}
		if !check {
			break
		}
	}
	if check {
		symms = append(symms, "horizontal")
	}

	if wide == height {

		check = true
		for i := 0; i < wide-1; i++ {
			for j := i + 1; j < wide; j++ {
				if cropped[i][j] != cropped[j][i] {
					check = false
					break
				}
			}
			if !check {
				break
			}
		}
		if check {
			symms = append(symms, "main_diagonal")
		}

		check = true
		for i := 0; i < wide-1; i++ {
			for j := 0; j < wide-1-i; j++ {
				if cropped[i][j] != cropped[wide-1-j][wide-1-i] {
					check = false
					break
				}
			}
			if !check {
				break
			}
		}
		if check {
			symms = append(symms, "extra_diagonal")
		}

	}

	return symms
}

func DotWithMinSumToOthersDots(dots []float64) float64 {
	n := len(dots)
	if n%2 == 0 {
		return (dots[n/2-1] + dots[n/2]) / 2
	}
	return dots[n/2]
}
