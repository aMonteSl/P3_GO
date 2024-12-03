package main

import (
	"math/rand"
	"sort"
	"testing"
)

func TestAirportSimulation(t *testing.T) {
	tests := []struct {
		name      string
		numPlanes int
		runways   int
		gates     int
	}{
		{"Equal distribution", 30, 3, 5},
		{"More high-priority planes", 50, 4, 6},
		{"Many low-priority planes", 40, 2, 4},
		{"Small airport", 15, 1, 2},
		{"Large airport", 100, 10, 15},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rand.Seed(42) // Fixed seed for reproducibility
			airport := newAirport(test.runways, test.gates)

			// Generate random planes
			planes := make([]Plane, 0, test.numPlanes)
			for i := 0; i < test.numPlanes; i++ {
				planes = append(planes, newPlane(i+1))
			}

			// Sort planes by priority
			sort.Slice(planes, func(i, j int) bool {
				return planes[i].priority < planes[j].priority
			})

			// Run simulation
			go airport.start()
			for _, plane := range planes {
				airport.queue <- plane
			}
			close(airport.queue)
			airport.wg.Wait()

			t.Logf("Test '%s' completed successfully.", test.name)
		})
	}
}
