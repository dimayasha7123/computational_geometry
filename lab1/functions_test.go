package lab1

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOXAngle(t *testing.T) {
	tests := []struct {
		name     string
		dot      Dot
		expected float64
	}{
		{
			name:     "0 degree (on OX right)",
			dot:      Dot{5, 0},
			expected: 0.0,
		},
		{
			name:     "45 degree",
			dot:      Dot{5, 5},
			expected: 45.0,
		},
		{
			name:     "90 degree (on OY up)",
			dot:      Dot{0, 5},
			expected: 90.0,
		},
		{
			name:     "135 degree",
			dot:      Dot{-5, 5},
			expected: 135.0,
		},
		{
			name:     "180 degree (on OX left)",
			dot:      Dot{-5, 0},
			expected: 180.0,
		},
		{
			name:     "225 degree",
			dot:      Dot{-5, -5},
			expected: 225.0,
		},
		{
			name:     "270 degree",
			dot:      Dot{0, -5},
			expected: 270.0,
		},
		{
			name:     "315 degree",
			dot:      Dot{5, -5},
			expected: 315.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OXAngle(tt.dot)

			assert.InDelta(t, got, tt.expected, EPS)
		})
	}
}

func TestCheckDotOnSegment(t *testing.T) {
	tests := []struct {
		name string
		a    Dot
		seg  Segment
		want bool
	}{
		{
			name: "on horizontal line",
			a:    Dot{X: 5, Y: 3},
			seg:  Segment{A: Dot{X: 2, Y: 3}, B: Dot{X: 7, Y: 3}},
			want: true,
		},
		{
			name: "left from horizontal line",
			a:    Dot{X: -123, Y: 3},
			seg:  Segment{A: Dot{X: 2, Y: 3}, B: Dot{X: 7, Y: 3}},
			want: false,
		},
		{
			name: "up from horizontal line",
			a:    Dot{X: 5, Y: 4},
			seg:  Segment{A: Dot{X: 2, Y: 3}, B: Dot{X: 7, Y: 3}},
			want: false,
		},
		{
			name: "on diagonal line",
			a:    Dot{X: 4, Y: 4},
			seg:  Segment{A: Dot{X: 2, Y: 3}, B: Dot{X: 6, Y: 5}},
			want: true,
		},
		{
			name: "on diagonal line (float)",
			a:    Dot{X: 4.5, Y: 4},
			seg:  Segment{A: Dot{X: 2, Y: 3}, B: Dot{X: 7, Y: 5}},
			want: true,
		},
		{
			name: "on border",
			a:    Dot{X: 7, Y: 5},
			seg:  Segment{Dot{2, 3}, Dot{7, 5}},
			want: true,
		},
		{
			name: "on border 2",
			a:    Dot{X: 4, Y: 3},
			seg:  Segment{Dot{0, 0}, Dot{4, 3}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckDotOnSegment(tt.a, tt.seg)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetAngleBAC(t *testing.T) {
	type args struct {
		A Dot
		B Dot
		C Dot
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "45 degree",
			args: args{Dot{5, 5}, Dot{0, 0}, Dot{5, 0}},
			want: 45.0,
		},
		{
			name: "rev 45 degree",
			args: args{Dot{5, 0}, Dot{0, 0}, Dot{5, 5}},
			want: 45.0,
		},
		{
			name: "90 degree",
			args: args{Dot{0, 0}, Dot{5, 0}, Dot{5, 5}},
			want: 90.0,
		},
		{
			name: "180 degree",
			args: args{Dot{0, 0}, Dot{5, 0}, Dot{10, 0}},
			want: 180.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAngleABC(tt.args.A, tt.args.B, tt.args.C)

			assert.InDelta(t, tt.want, got, EPS)
		})
	}
}

func TestNormOnSegment(t *testing.T) {
	type args struct {
		A   Dot
		Seg Segment
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "on",
			args: args{Dot{2, 5}, Segment{Dot{0, 0}, Dot{5, 4}}},
			want: true,
		},
		{
			name: "outside",
			args: args{Dot{2, 9}, Segment{Dot{0, 0}, Dot{5, 4}}},
			want: false,
		},
		{
			name: "on border",
			args: args{Dot{2, 7}, Segment{Dot{0, 0}, Dot{5, 4}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormOnSegment(tt.args.A, tt.args.Seg)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLineSegmentItersection(t *testing.T) {
	type args struct {
		l Line
		s Segment
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple intersection",
			args: args{*FromSegment(Segment{Dot{2, 2}, Dot{3, 5}}), Segment{Dot{0, 0}, Dot{4, 3}}},
			want: true,
		},
		{
			name: "intersection dot outside",
			args: args{*FromSegment(Segment{Dot{3, 5}, Dot{8, 4}}), Segment{Dot{0, 0}, Dot{4, 3}}},
			want: false,
		},
		{
			name: "intersection dot on border",
			args: args{*FromSegment(Segment{Dot{5, 1}, Dot{3, 5}}), Segment{Dot{0, 0}, Dot{4, 3}}},
			want: true,
		},
		{
			name: "parallel",
			args: args{*FromSegment(Segment{Dot{-1, 2}, Dot{3, 5}}), Segment{Dot{0, 0}, Dot{4, 3}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LineSegmentIntersection(tt.args.l, tt.args.s)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSegmentsIntersection(t *testing.T) {
	type args struct {
		s1 Segment
		s2 Segment
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple intersection",
			args: args{Segment{Dot{0, 0}, Dot{4, 3}}, Segment{Dot{3, -1}, Dot{2, 4}}},
			want: true,
		},
		{
			name: "not intersects",
			args: args{Segment{Dot{0, 0}, Dot{5, 0}}, Segment{Dot{0, 4}, Dot{10, 20}}},
			want: false,
		},
		{
			name: "intersection on border of one segments",
			args: args{Segment{Dot{0, 0}, Dot{4, 3}}, Segment{Dot{3, 5}, Dot{5, 1}}},
			want: true,
		},
		{
			name: "intersection on both borders",
			args: args{Segment{Dot{0, 0}, Dot{4, 3}}, Segment{Dot{4, 3}, Dot{3, -2}}},
			want: true,
		},
		{
			name: "parallel segments",
			args: args{Segment{Dot{0, 0}, Dot{4, 3}}, Segment{Dot{-1, 2}, Dot{3, 5}}},
			want: false,
		},
		{
			name: "equal segments",
			args: args{Segment{Dot{0, 0}, Dot{4, 3}}, Segment{Dot{0, 0}, Dot{4, 3}}},
			want: true,
		},
		{
			name: "on segment on the end of another",
			args: args{Segment{Dot{0, 0}, Dot{4, 3}}, Segment{Dot{2, 1.5}, Dot{6, 4.5}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SegmentsIntersection(tt.args.s1, tt.args.s2)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDotAndTriangle(t *testing.T) {
	type args struct {
		d Dot
		t Triangle
	}
	tests := []struct {
		name string
		args args
		want Position
	}{
		{
			name: "simple inside",
			args: args{Dot{2, 4}, Triangle{Dot{0, 0}, Dot{5, 4}, Dot{1, 7}}},
			want: Inside,
		},
		{
			name: "on border",
			args: args{Dot{3, 5.5}, Triangle{Dot{0, 0}, Dot{5, 4}, Dot{1, 7}}},
			want: OnBorder,
		},
		{
			name: "on corner",
			args: args{Dot{5, 4}, Triangle{Dot{0, 0}, Dot{5, 4}, Dot{1, 7}}},
			want: OnBorder,
		},
		{
			name: "outside",
			args: args{Dot{4, 0}, Triangle{Dot{0, 0}, Dot{5, 4}, Dot{1, 7}}},
			want: Outside,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DotAndTriangle(tt.args.d, tt.args.t)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCheckDotOnSegmentInteger(t *testing.T) {
	tests := []struct {
		name string
		a    Dot
		seg  Segment
		want bool
	}{
		{
			name: "on horizontal line",
			a:    Dot{X: 5, Y: 3},
			seg:  Segment{A: Dot{X: 2, Y: 3}, B: Dot{X: 7, Y: 3}},
			want: true,
		},
		{
			name: "left from horizontal line",
			a:    Dot{X: -123, Y: 3},
			seg:  Segment{A: Dot{X: 2, Y: 3}, B: Dot{X: 7, Y: 3}},
			want: false,
		},
		{
			name: "up from horizontal line",
			a:    Dot{X: 5, Y: 4},
			seg:  Segment{A: Dot{X: 2, Y: 3}, B: Dot{X: 7, Y: 3}},
			want: false,
		},
		{
			name: "on diagonal line",
			a:    Dot{X: 4, Y: 4},
			seg:  Segment{A: Dot{X: 2, Y: 3}, B: Dot{X: 6, Y: 5}},
			want: true,
		},
		{
			name: "on border",
			a:    Dot{X: 7, Y: 5},
			seg:  Segment{Dot{2, 3}, Dot{7, 5}},
			want: true,
		},
		{
			name: "on border 2",
			a:    Dot{X: 4, Y: 3},
			seg:  Segment{Dot{0, 0}, Dot{4, 3}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckDotOnSegment(tt.a, tt.seg)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMaxIntersectionLine(t *testing.T) {
	type args struct {
		segments []Segment
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				segments: []Segment{
					{Dot{3.0, 1.0}, Dot{2.0, 5.0}},
					{Dot{3.0, 8.0}, Dot{5.0, 10.0}},
					{Dot{5.0, 4.0}, Dot{6.0, 7.0}},
					{Dot{9.0, 3.0}, Dot{8.0, 6.0}},
				}},
			want: []int{0, 2, 3},
		},
		{
			name: "test2",
			args: args{
				segments: []Segment{
					{Dot{3.0, 1.0}, Dot{2.0, 5.0}},
					{Dot{3.0, 8.0}, Dot{5.0, 10.0}},
				}},
			want: []int{0, 1},
		},
		{
			name: "one_segment",
			args: args{
				segments: []Segment{
					{Dot{3.0, 1.0}, Dot{2.0, 5.0}},
				}},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := MaxIntersectionLine(tt.args.segments)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxIntersectionLine() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDotWithMinSumToOthersDots(t *testing.T) {
	type args struct {
		dots []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "even",
			args: args{
				dots: []float64{2, 4, 7, 9},
			},
			want: 5.5,
		},
		{
			name: "odd",
			args: args{
				dots: []float64{2, 4, 6},
			},
			want: 4,
		},
		{
			name: "negative",
			args: args{
				dots: []float64{-5, -10, -20},
			},
			want: -10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DotWithMinSumToOthersDots(tt.args.dots); got != tt.want {
				t.Errorf("DotWithMinSumToOthersDots() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolygon_IsSimple(t *testing.T) {
	tests := []struct {
		name string
		p    Polygon
		want bool
	}{
		{
			name: "intersects",
			p:    *NewPolygon([]Dot{{0, 0}, {5, 3}, {9, 2}, {12, -4}, {5, -7}, {7, -2}, {9, -9}, {1, -11}, {-2, -5}}),
			want: false,
		},
		{
			name: "not_intersects",
			p:    *NewPolygon([]Dot{{0, 0}, {4, 3}, {5, -2}, {3, -1}, {1, -4}}),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsSimple(); got != tt.want {
				t.Errorf("Polygon.IsSimple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolygon_Square(t *testing.T) {
	tests := []struct {
		name    string
		p       Polygon
		want    float64
		wantErr error
	}{
		{
			name:    "convex_polygon",
			p:       *NewPolygon([]Dot{{0, 0}, {4, 3}, {7, -1}, {5, -5}, {2, -3}}),
			want:    30,
			wantErr: nil,
		},
		{
			name:    "not_convex_polygon",
			p:       *NewPolygon([]Dot{{0, 0}, {-2, 4}, {0, 7}, {4, 7}, {2, 4}, {6, 1}}),
			want:    31,
			wantErr: nil,
		},
		{
			name:    "not_simple_polygon",
			p:       *NewPolygon([]Dot{{0, 0}, {5, 3}, {9, 2}, {12, -4}, {5, -7}, {7, -2}, {9, -9}, {1, -11}, {-2, -5}}),
			want:    0,
			wantErr: fmt.Errorf(NotSimplePolygonError),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Square()
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
			if got != tt.want {
				t.Errorf("Polygon.Square() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolygon_IsConvex(t *testing.T) {
	tests := []struct {
		name string
		p    Polygon
		want bool
	}{
		{
			name: "convex1",
			p:    *NewPolygon([]Dot{{0, 0}, {5, 3}, {9, 2}, {12, -4}, {9, -9}, {1, -11}, {-2, -5}}),
			want: true,
		},
		{
			name: "convex2",
			p:    *NewPolygon([]Dot{{0, 0}, {4, 3}, {7, -1}, {5, -5}, {2, -3}}),
			want: true,
		},
		{
			name: "not_convex1",
			p:    *NewPolygon([]Dot{{0, 0}, {-2, 4}, {0, 7}, {4, 7}, {2, 4}, {6, 1}}),
			want: false,
		},
		{
			name: "not_convex2",
			p:    *NewPolygon([]Dot{{0, 0}, {4, 3}, {5, -2}, {3, -1}, {1, -4}}),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsConvex(); got != tt.want {
				t.Errorf("Polygon.IsConvex() = %v, want %v", got, tt.want)
			}
		})
	}
}
