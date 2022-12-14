package lab1

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
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

func MaxIntersectionLine(segments []Segment) (Segment, []int) {
	if len(segments) == 1 {
		return Segment{segments[0].A, Dot{segments[0].B.X + 10, segments[0].B.Y + 10}}, []int{0}
	}

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

// no tests... but i'm tired... just believe, it works
func CheckSymmetry(matrix [][]int) ([][]int, []string) {
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

	if wide <= 0 || height <= 0 {
		return [][]int{}, []string{}
	}

	cropped := make([][]int, height)
	for i := 0; i < height; i++ {
		cropped[i] = make([]int, wide)
		for j := 0; j < wide; j++ {
			cropped[i][j] = matrix[i+up][j+left]
		}
	}

	// fmt.Println(strings.ReplaceAll(fmt.Sprint(cropped), "] ", "]\n "))

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

	return cropped, symms
}

func DotWithMinSumToOthersDots(dots []float64) float64 {
	n := len(dots)
	if n%2 == 0 {
		return (dots[n/2-1] + dots[n/2]) / 2
	}
	return dots[n/2]
}

type Polygon []Dot

// in polygon on last place we have first element to simplify calculations
func NewPolygon(dots []Dot) *Polygon {
	if len(dots) == 0 {
		return &Polygon{}
	}
	dots = append(dots, dots[0])
	p := make(Polygon, len(dots))
	copy(p, dots)
	return &p
}

func (p Polygon) String() string {
	sb := strings.Builder{}
	sb.WriteString("Polygon dots:\n")
	sb.WriteString("   X     Y \n")
	sb.WriteString("-----------\n")
	for i := 0; i < len(p); i++ {
		sb.WriteString(fmt.Sprintf("%5.1f %5.1f\n", p[i].X, p[i].Y))
	}
	out := sb.String()
	out = out[:len(out)-1]
	return out
}

func (p Polygon) IsSimple() bool {
	for i := 0; i < len(p)-3; i++ {
		for j := i + 2; j < len(p)-1; j++ {
			if i == 0 && j == len(p)-2 {
				continue
			}
			s1 := Segment{p[i], p[i+1]}
			s2 := Segment{p[j], p[j+1]}
			if SegmentsIntersection(s1, s2) {
				return false
			}
		}
	}
	return true
}

const (
	NotSimplePolygonError string = "polygon isn't simple (have edges intersection)"
)

func (p Polygon) Area() (float64, error) {
	if !p.IsSimple() {
		return 0, fmt.Errorf(NotSimplePolygonError)
	}
	summ := 0.0
	for i := 0; i < len(p)-1; i++ {
		summ += (p[i+1].X - p[i].X) * (p[i+1].Y + p[i].Y)
	}
	summ /= 2
	return summ, nil
}

func (p Polygon) IsConvex() bool {
	sdtas := make([]float64, len(p)-1)
	for i := 0; i < len(p)-2; i++ {
		sdtas[i] = SignedDoubleTriangleArea(p[i], p[i+1], p[i+2])
	}
	sdtas[len(sdtas)-1] = SignedDoubleTriangleArea(p[len(p)-1], p[0], p[1])
	right := 0
	left := 0
	for _, v := range sdtas {
		if math.Abs(v) < EPS {
			continue
		}
		if v > 0 {
			right++
		} else {
			left--
		}
	}
	if left == 0 || right == 0 {
		return true
	}
	return false
}

// CCH is Construction of a Convex Hull
// check this for understanding http://www.e-maxx-ru.1gb.ru/algo/convex_hull_graham
func CCHGrahamAndrew(dots []Dot) Polygon {
	if len(dots) <= 2 {
		return *NewPolygon(dots)
	}

	dotsCopy := make([]Dot, len(dots))
	copy(dotsCopy, dots)

	// sort dots by X ascending, if equal, then by Y ascending
	sort.Slice(dotsCopy, func(i, j int) bool {
		switch {
		case dotsCopy[i].X < dotsCopy[j].X:
			return true
		case dotsCopy[i].X > dotsCopy[j].X:
			return false
		case dotsCopy[i].Y < dotsCopy[j].Y:
			return true
		case dotsCopy[i].Y > dotsCopy[j].Y:
			return false
		default:
			return true // shit in your own pants, if u call this function this with two equal dots
		}
	})

	leftDownDot := dotsCopy[0]
	rightUpDot := dotsCopy[len(dotsCopy)-1]
	up := make([]Dot, 0, len(dotsCopy)/2)
	up = append(up, rightUpDot, leftDownDot) // its stupid to add this two dots (because then i remove it), but i've done it, because i can
	downPrep := make([]Dot, 0, len(dotsCopy)/2)

	for i := 1; i < len(dotsCopy)-1; i++ {
		sign := SignedDoubleTriangleArea(leftDownDot, rightUpDot, dotsCopy[i])
		switch {
		case math.Abs(sign) < EPS:
			continue
		case sign > 0:
			up = append(up, dotsCopy[i])
		default:
			downPrep = append(downPrep, dotsCopy[i])
		}
	}
	up = append(up, rightUpDot)

	// add leftDownDot, rightUpDot and reversed downPrep to down
	down := make([]Dot, len(downPrep)+2)
	down[0] = leftDownDot
	down[1] = rightUpDot
	for i := 2; i < len(down); i++ {
		down[i] = downPrep[len(downPrep)+1-i]
	}
	down = append(down, leftDownDot)

	output := make([]Dot, 1)
	output[0] = leftDownDot
	output = makeOneHalfConvexHull(up, output)
	output = makeOneHalfConvexHull(down, output)
	output = output[:len(output)-1]

	// fmt.Println("up:    ", up)
	// fmt.Println("down:  ", down)
	// fmt.Println("output:", output)

	return *NewPolygon(output)
}

func makeOneHalfConvexHull(dots []Dot, output []Dot) []Dot {
	if len(dots) != 2 {
		output = append(output, dots[2])
		for i := 3; i < len(dots); i++ {
			sign := SignedDoubleTriangleArea(output[len(output)-2], output[len(output)-1], dots[i])
			for math.Abs(sign) < EPS || math.Abs(sign) > EPS && sign > 0 {
				output = output[:len(output)-1]
				if len(output) <= 2 {
					break
				}
				sign = SignedDoubleTriangleArea(output[len(output)-2], output[len(output)-1], dots[i])
			}
			output = append(output, dots[i])
		}
	}
	return output
}

func CheckAllDotsAreRightFromLineAB(a, b Dot, dots []Dot) bool {
	for _, d := range dots {
		if d == a || d == b {
			continue
		}
		sign := SignedDoubleTriangleArea(a, b, d)
		if math.Abs(sign) < EPS && !CheckDotOnSegment(d, Segment{a, b}) || sign > 0 {
			return false
		}
	}
	return true
}

func CCHJarvis(dots []Dot) Polygon {
	n := len(dots)
	met := make([]bool, n)

	var a, b Dot
	found := false
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if CheckAllDotsAreRightFromLineAB(dots[i], dots[j], dots) {
				a = dots[i]
				b = dots[j]

				met[i] = true
				met[j] = true

				found = true
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		panic("don't found line ab")
	}
	output := make([]Dot, 2)
	output[0] = a
	output[1] = b

	for {
		found := false
		for i, d := range dots {
			if met[i] {
				continue
			}
			if CheckAllDotsAreRightFromLineAB(output[len(output)-1], d, dots) {
				output = append(output, d)
				met[i] = true
				found = true
				break
			}
		}
		if !found {
			break
		}
	}
	// fmt.Println(output)
	return *NewPolygon(output)
}

func CCHDivideAndConquer(dots []Dot) Polygon {
	ROUTINES := 16
	FUNC := CCHGrahamAndrew

	dotSets := make([][]Dot, ROUTINES)
	for i := 0; i < ROUTINES; i++ {
		dotSets[i] = make([]Dot, 0)
	}
	for i, d := range dots {
		dotSets[i%ROUTINES] = append(dotSets[i%ROUTINES], d)
	}

	hulls := make([]Polygon, ROUTINES)
	wg := &sync.WaitGroup{}
	wg.Add(ROUTINES)
	for i := range hulls {
		go func(i int, wg *sync.WaitGroup) {
			hulls[i] = FUNC(dotSets[i])
			if len(hulls[i]) != 0 {
				hulls[i] = hulls[i][:len(hulls[i])-1]
			}
			wg.Done()
		}(i, wg)
	}
	wg.Wait()

	unity := make([]Dot, 0)
	for i := range hulls {
		for j := range hulls[i] {
			d := hulls[i][j]
			check := false
			for k := range unity {
				if unity[k] == d {
					check = true
					break
				}
			}
			if !check {
				unity = append(unity, d)
			}
		}
	}

	output := FUNC(unity)
	output = output[:len(output)-1]

	return *NewPolygon(output)
}
