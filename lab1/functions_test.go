package lab1

import (
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
