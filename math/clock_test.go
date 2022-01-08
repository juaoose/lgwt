package clock_test

import (
	"testing"
	"time"

	"github.com/juaoose/lgwt/math/clock"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	// Draw a line up with length 90
	want := clock.Point{X: 150, Y: 150 - 90}
	got := clock.SecondHand(tm)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
