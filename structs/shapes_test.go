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

	// Table driven tests: build a list of test cases that
	// can be tested in the same manner

	// "Table based tests can be a great item in your toolbox
	// but be sure that you have a need for the extra noise in
	// the tests. If you wish to test various implementations
	// of an interface, or if the data being passed in to a
	// function has lots of different requirements that need
	// testing then they are a great fit."

	// This is an anonymous struct that defines the test cases
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{5, 4}, 20.0},
		{Circle{5}, 78.53981633974483},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}
