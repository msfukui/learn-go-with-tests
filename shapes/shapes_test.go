package shapes

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10}, math.Pi * 10 * 10},
		{Triangle{12, 6}, 36},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}
}
