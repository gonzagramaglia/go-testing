package perimeter

import (
	"math"
	"testing"
)

func TestGetPerimeter(t *testing.T) {

	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.GetPerimeter()
		if got != want {
			t.Errorf("got: %.2f want %.2f", got, want)
		}

	}

	t.Run("rectangles", func(t *testing.T) {

		rectangle := Rectangle{10.0, 10.0}
		checkPerimeter(t, rectangle, 40.0)

	})

	t.Run("circles", func(t *testing.T) {

		circle := Circle{10.0}
		want := 2 * math.Pi * circle.Radius
		checkPerimeter(t, circle, want)

	})

}

func TestGetArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {

		t.Helper()
		got := shape.GetArea()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {

		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)

	})

	t.Run("circles", func(t *testing.T) {

		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)

	})
}
