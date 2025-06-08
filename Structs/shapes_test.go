package shapes

import (
	"math"
	"testing"
)

func AreEqual(x, y float64) bool {
	return math.Abs(x-y) <= 1e-3
}

func AssertPerimeter(want float64, shape Shape, t testing.TB) {
	t.Helper()

	got := shape.Perimeter()

	if !AreEqual(got, want) {
		t.Errorf("%#v Perimeter is wrong got %.10f want %.10f", shape, got, want)
	}
}

func AssertArea(want float64, shape Shape, t testing.TB) {
	t.Helper()

	got := shape.Area()

	if !AreEqual(got, want) {
		t.Errorf("%#v Area is wrong got %.10f want %.10f", shape, got, want)
	}
}

func TestPerimeter(t *testing.T) {
	testCases := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "NormRectangle", shape: Rectangle{2, 5}, want: 14},
		{name: "ZeroRectangle", shape: Rectangle{0, 0}, want: 0},
		{name: "NormCircle", shape: Circle{10}, want: 62.832},
		{name: "ZeroCircle", shape: Circle{0}, want: 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			AssertPerimeter(testCase.want, testCase.shape, t)
		})
	}
}

func TestArea(t *testing.T) {
	testCases := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "NormRectangle", shape: Rectangle{2, 5}, want: 10},
		{name: "ZeroRectangle", shape: Rectangle{0, 0}, want: 0},
		{name: "NormCircle", shape: Circle{10}, want: 314.159},
		{name: "ZeroCircle", shape: Circle{0}, want: 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			AssertArea(testCase.want, testCase.shape, t)
		})
	}
}
