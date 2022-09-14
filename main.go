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
	seg := Segment{A: Dot{X: 2, Y: 3}, B: Dot{X: 7, Y: 5}}
	onSeg := CheckDotOnSegment(a, seg)
	if onSeg {
		fmt.Printf("Dot %v is on segment (%v)\n", a, seg)
	} else {
		fmt.Printf("Dot %v is NOT on segment (%v)\n", a, seg)
	}
}

func Task3() {
	a := Dot{X: 2, Y: 5}
	seg := Segment{A: Dot{X: 0, Y: 0}, B: Dot{X: 5, Y: 4}}
	onSeg := NormOnSegment(a, seg)
	if onSeg {
		fmt.Printf("Normal of dot %v on segment %v is on segment\n", a, seg)
	} else {
		fmt.Printf("Normal of dot %v on segment %v is outside of segment\n", a, seg)
	}
}

func main() {
	Task1()
	Task2()
	Task3()
}
