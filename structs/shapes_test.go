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

	t.Run("check triangle parameter", func(t *testing.T) {
		rectangle := Triangle{10.0, 10.0}
		want := 30.0

		checkPerimeter(t, rectangle, want)
	})

}

func TestArea(t *testing.T) {

	// Table driven tests: build a list of test cases that
	// can be tested in the same manner
	// https://github.com/golang/go/wiki/TableDrivenTests

	// This is an anonymous struct that defines the test cases
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		//Named elements instead of positioinal, for extra
		// clarity
		{shape: Rectangle{Width: 5, Height: 4}, want: 20.0},
		{shape: Circle{Radius: 5}, want: 78.53981633974483},
		{shape: Triangle{Base: 3, Height: 2}, want: 3.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}

}
