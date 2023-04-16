package ships

import (
	"fmt"
	"testing"
)

func ExamplePoint_Add_a() {
	a := Point{X: 10, Y: 10}
	b := Point{X: 10, Y: 10}
	r := a.Add(b)
	fmt.Println(r.X, r.Y)
	// Output: 20 20
}

func ExamplePoint_Add_b() {
	a := Point{X: 20, Y: 10}
	b := Point{X: 10, Y: 20}
	r := a.Add(b)
	fmt.Println(r.X, r.Y)
	// Output: 30 30
}

func ExamplePoint_Add_c() {
	a := Point{X: 15, Y: 25}
	b := Point{X: 25, Y: 35}
	r := a.Add(b)
	fmt.Println(r.X, r.Y)
	// Output: 40 60
}

func ExamplePoint_Add_d() {
	a := Point{X: 10, Y: 10}
	b := Point{X: -10, Y: -10}
	r := a.Add(b)
	fmt.Println(r.X, r.Y)
	// Output: 0 0
}

func TestPoint_Add(t *testing.T) {
	tests := []struct {
		name       string
		ship       Ship
		moveTo     Point
		shouldShip Ship
	}{
		{"test A",
			Ship([]Point{{10, 20}}),
			Point{10, 20},
			Ship([]Point{{10, 20}}),
		},
		{"test B",
			Ship([]Point{{10, 10}, {10, 11}, {11, 10}, {11, 11}}),
			Point{25, 35},
			Ship([]Point{{25, 35}, {25, 36}, {26, 35}, {26, 36}}),
		},
		{"test C",
			Ship([]Point{{10, 20}, {11, 20}, {12, 20}, {13, 20}}),
			Point{-10, -20},
			Ship([]Point{{-10, -20}, {-9, -20}, {-8, -20}, {-7, -20}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nShip := tt.ship.MoveTo(tt.moveTo)
			if len(tt.ship) != len(nShip) {
				t.Errorf("Expected %v but got %v", len(tt.shouldShip), len(nShip))
			}
			for i, j := range nShip {
				if j.X != tt.shouldShip[i].X {
					t.Errorf("Expected X in %vth point %v but got %v", i, j.X, tt.shouldShip[i].X)
				}
				if j.Y != tt.shouldShip[i].Y {
					t.Errorf("Expected Y in %vth point %v but got %v", i, j.Y, tt.shouldShip[i].Y)
				}
			}
		})
	}
}
