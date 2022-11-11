package main

import (
	"fmt"
	"strings"

	. "github.com/dimayasha7123/computational_geometry/lab1"
)

func Task1() {
	a := Dot{X: -5, Y: -5}
	b := Dot{X: 5, Y: 5}
	fmt.Println(WhatAngleWithOxMore(a, b))
}

func Task2() {
	a := Dot{X: 4.5, Y: 4}
	s := Segment{A: Dot{X: 2, Y: 3}, B: Dot{X: 7, Y: 5}}
	onSeg := CheckDotOnSegment(a, s)
	if onSeg {
		fmt.Printf("Dot %v is on segment (%v)\n", a, s)
	} else {
		fmt.Printf("Dot %v is NOT on segment (%v)\n", a, s)
	}
}

func Task3() {
	a := Dot{X: 1, Y: 10}
	s := Segment{A: Dot{X: 0, Y: 0}, B: Dot{X: 5, Y: 4}}
	onSeg := NormOnSegment(a, s)
	if onSeg {
		fmt.Printf("Normal of dot %v on segment %v is on segment\n", a, s)
	} else {
		fmt.Printf("Normal of dot %v on segment %v is outside of segment\n", a, s)
	}
}

func Task4() {
	l := *FromSegment(Segment{A: Dot{X: 2, Y: 2}, B: Dot{X: 3, Y: 5}})
	s := Segment{A: Dot{X: 0, Y: 0}, B: Dot{X: 0, Y: 5}}
	isInter := LineSegmentIntersection(l, s)
	if isInter {
		fmt.Printf("Line %v intersects with segment %v\n", l, s)
	} else {
		fmt.Printf("Line %v NOT intersects with segment %v\n", l, s)
	}
}

func Task5() {
	s1 := Segment{A: Dot{X: 0, Y: 0}, B: Dot{X: 4, Y: 3}}
	s2 := Segment{A: Dot{X: 3, Y: -1}, B: Dot{X: 2, Y: 4}}
	isInter := SegmentsIntersection(s1, s2)
	if isInter {
		fmt.Printf("Segment %v intersects with segment %v\n", s1, s2)
	} else {
		fmt.Printf("Segment %v NOT intersects with segment %v\n", s1, s2)
	}
}

func Task6() {
	d := Dot{X: 0, Y: 5}
	t := Triangle{A: Dot{X: 0, Y: 0}, B: Dot{X: 5, Y: 4}, C: Dot{X: 1, Y: 7}}
	pos := DotAndTriangle(d, t)

	var sPos string
	switch pos {
	case Inside:
		sPos = "inside"
	case OnBorder:
		sPos = "on border of"
	case Outside:
		sPos = "outside"
	}

	fmt.Printf("Dot %v is %s triangle %v\n", d, sPos, t)
}

func Lab1() {
	Task1()
	Task2()
	Task3()
	Task4()
	Task5()
	Task6()
}

func Task7() {
	segments := []Segment{
		{A: Dot{X: 0.0, Y: 0.0}, B: Dot{X: 0.0, Y: 1.0}},
		{A: Dot{X: 1.0, Y: 0.0}, B: Dot{X: 1.0, Y: -1.0}},
		{A: Dot{X: 2.0, Y: 0.0}, B: Dot{X: 2.0, Y: 1.0}},
		{A: Dot{X: 3.0, Y: 0.0}, B: Dot{X: 3.0, Y: -1.0}},
		{A: Dot{X: 5.0, Y: 5.0}, B: Dot{X: 5.0, Y: 6.0}},
	}
	lineSegment, segment_nums := MaxIntersectionLine(segments)

	fmt.Println("Task 1")
	fmt.Println("Segments:")
	for i, s := range segments {
		fmt.Printf("%d: (%.1f, %.1f), (%.1f, %.1f)\n",
			i,
			s.A.X,
			s.A.Y,
			s.B.X,
			s.B.Y)
	}
	fmt.Printf("Line with dots (%.1f, %.1f), (%.1f, %.1f) intersects with %d segments: %s\n",
		lineSegment.A.X,
		lineSegment.A.Y,
		lineSegment.B.X,
		lineSegment.B.Y,
		len(segment_nums),
		fmt.Sprint(segment_nums)[1:len(fmt.Sprint(segment_nums))-1],
	)
}

func Task8() {

	matrix := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	cropped, symms := CheckSymmetry(matrix)

	fmt.Println("Task 2")
	fmt.Println("Input drawing:")
	fmt.Println(strings.ReplaceAll(fmt.Sprint(matrix), "] ", "]\n "))
	fmt.Println("Cropped drawing:")
	fmt.Println(strings.ReplaceAll(fmt.Sprint(cropped), "] ", "]\n "))
	fmt.Println("Axes of symmetry:", fmt.Sprint(symms)[1:len(fmt.Sprint(symms))-1])

}

func Task9() {
	dots := []float64{-1, 0, 1, 100, 1000}
	middle := DotWithMinSumToOthersDots(dots)

	fmt.Println("Task 3")
	fmt.Println("Dots:  ", fmt.Sprint(dots)[1:len(fmt.Sprint(dots))-1])
	fmt.Println("Middle:", middle)
}

func Lab2() {
	Task7()
	Task8()
	Task9()
}

func Task10() {
	p := *NewPolygon([]Dot{{0, 0}, {5, 0}, {5, 5}, {0, 5}, {0, 10}, {5, 5}})
	simple := p.IsSimple()

	fmt.Println("Task 1")
	fmt.Println(p)
	if simple {
		fmt.Println("Polygon is simple")
	} else {
		fmt.Println("Polygon is NOT simple")
	}
}

func Task11() {
	p := *NewPolygon([]Dot{{0, 0}, {10, 0}, {10, 10}, {5,5}, {0, 10}})
	area, err := p.Area()

	fmt.Println("Task 2")
	fmt.Println(p)
	if err != nil {
		fmt.Println("Oooops, error:", err)
		return
	}
	fmt.Println("Area: ", area)
}

func Task12() {
	p := *NewPolygon([]Dot{{0, 0}, {10, 0}, {10, 10}, {5,8}, {0, 2}})
	simple := p.IsConvex()

	fmt.Println("Task 3")
	fmt.Println(p)
	if simple {
		fmt.Println("Polygon is convex")
	} else {
		fmt.Println("Polygon is NOT convex")
	}
}

func Lab3() {
	Task10()
	Task11()
	Task12()
}

func Task13() {
	dots := []Dot{
		{X: 0, Y: 0},
		{X: 10, Y: 0},
		{X: 5, Y: 10},
		{X: 5, Y: 5},
	}
	hull := CCHGrahamAndrew(dots)

	dotsStr := fmt.Sprint(dots)
	dotsStr = dotsStr[1:len(dotsStr)-1]
	fmt.Println("GRAHAM ANDREW ALGORITHM")
	fmt.Println("All dots:", dotsStr)
	fmt.Println("Hull:", hull)
}

func Task14() {
	dots := []Dot{
		{X: 0, Y: 0},
		{X: 2, Y: 0},
		{X: 3, Y: 1},
		{X: 5, Y: 0},
		{X: 1, Y: -2},
		{X: 3, Y: -2},
		{X: 5, Y: -2},
		{X: 6, Y: 2},
		{X: 6, Y: 5},
	}
	hull := CCHJarvis(dots)

	dotsStr := fmt.Sprint(dots)
	dotsStr = dotsStr[1:len(dotsStr)-1] 
	fmt.Println("JARVIS ALGORITHM")
	fmt.Println("All dots:", dotsStr)
	fmt.Println("Hull:", hull)
}

func Task15() {
	dots := []Dot{
		{X: 0, Y: 0},
		{X: 2, Y: 0},
		{X: 3, Y: 1},
		{X: 5, Y: 0},
		{X: 1, Y: -2},
		{X: 3, Y: -2},
		{X: 5, Y: -2},
		{X: 6, Y: 2},
		{X: 6, Y: 5},
	}
	hull := CCHJarvis(dots)

	dotsStr := fmt.Sprint(dots)
	dotsStr = dotsStr[1:len(dotsStr)-1] 
	fmt.Println("DIVIDE AND CONQUER ALGORITHM")
	fmt.Println("All dots:", dotsStr)
	fmt.Println("Hull:", hull)
}

func Lab4() {
	Task13()
	Task14()
	Task15()
}

func main() {
	// Lab1()
	// Lab2()
	// Lab3()
	Lab4()

	// p1 := Dot{X: -1, Y: -5}
	// p2 := Dot{X: 0, Y: 0}
	// p3 := Dot{X: 6, Y: 1}
	// p4 := Dot{X: 1, Y: 5}
	// p5 := Dot{X: -5, Y: 4}

	// a1 := SignedDoubleTriangleArea(p1, p2, p3)
	// a2 := SignedDoubleTriangleArea(p1, p2, p4)
	// a3 := SignedDoubleTriangleArea(p1, p2, p5)

	// fmt.Println("a1:", a1)
	// fmt.Println("a2:", a2)
	// fmt.Println("a3:", a3)

}
