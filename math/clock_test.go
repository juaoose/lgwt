package clock

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {

	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondsInRadians(c.time)

			if !aproxFloat64Equal(got, c.angle) {
				t.Errorf("got %v radians, want %v radians", got, c.angle)
			}
		})
	}

	thirtySecs := time.Date(1, 0, 0, 0, 0, 30, 0, time.UTC)

	// 360 deg = 2pi radians
	want := math.Pi
	got := secondsInRadians(thirtySecs)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minutesInRadians(c.time)
			if !aproxFloat64Equal(got, c.angle) {
				t.Errorf("got %v, want %v", got, c.angle)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hoursInRadians(c.time)
			if !aproxFloat64Equal(got, c.angle) {
				t.Errorf("got %v, want %v", got, c.angle)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(*testing.T) {
			got := secondHandPoint(c.time)

			// Since this is not critical, we have an approximate equality check
			if !aproxPointEqual(got, c.point) {
				t.Errorf("got %v, want %v", got, c.point)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(*testing.T) {
			got := minuteHandPoint(c.time)

			if !aproxPointEqual(got, c.point) {
				t.Errorf("got %v, want %v", got, c.point)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(*testing.T) {
			got := hourHandPoint(c.time)

			if !aproxPointEqual(got, c.point) {
				t.Errorf("got %v, want %v", got, c.point)
			}
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(0, 0, 0, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func aproxFloat64Equal(a, b float64) bool {
	const eqThres = 1e-7
	return math.Abs(a-b) < eqThres
}

func aproxPointEqual(a, b Point) bool {
	return aproxFloat64Equal(a.X, b.X) && aproxFloat64Equal(a.Y, b.Y)
}
