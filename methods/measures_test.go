package methods

import (
	"math"
	"testing"
)

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

func (t Triangle) Perimeter() float64 {
	return t.Base + t.Base + t.Height
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func TestPerimeter(t *testing.T) {
	verifyPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}

		verifyPerimeter(t, rectangle, 40.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}

		verifyPerimeter(t, circle, 62.83185307179586)
	})
}

func TestArea(t *testing.T) {
	testsArea := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12.0, Length: 6.0}, want: 72.0},
		{shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range testsArea {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f, want %.2f", got, tt.want)
		}
	}
}
