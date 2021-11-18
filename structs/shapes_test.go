package shapes

import "testing"

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("check rectangle parameter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 40.0

		checkPerimeter(t, rectangle, want)
	})

	t.Run("check circle parameter", func(t *testing.T) {
		rectangle := Circle{10.0}
		want := 62.83185307179586

		checkPerimeter(t, rectangle, want)
	})

}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("calculate rectangle area", func(t *testing.T) {
		rectangle := Rectangle{5.0, 4.0}
		want := 20.0

		checkArea(t, rectangle, want)

	})

	t.Run("calculate circle area", func(t *testing.T) {
		circle := Circle{5}
		want := 78.53981633974483

		checkArea(t, circle, want)
	})
}
