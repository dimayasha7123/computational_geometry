package main

import (
	"fmt"

	. "github.com/dimayasha7123/computational_geometry/lab1"
)

func Task1() {
	a := Dot{X: 5, Y: 2}
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
	a := Dot{X: 2, Y: 5}
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
	s := Segment{A: Dot{X: 0, Y: 0}, B: Dot{X: 4, Y: 3}}
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

func main() {
	Task1()
	Task2()
	Task3()
	Task4()
}
