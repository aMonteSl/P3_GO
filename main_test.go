package main

import (
	"math/rand"
	"sort"
	"testing"
)

func TestAirportSimulation(t *testing.T) {
	tests := []struct {
		name       string
		numTypeA   int // Planes with >100 passengers (Category A)
		numTypeB   int // Planes with 50-100 passengers (Category B)
		numTypeC   int // Planes with <50 passengers (Category C)
		runways    int // Number of runways
		gates      int // Number of gates
	}{
		{"Equal distribution (10 each)", 10, 10, 10, 3, 5},
		{"More high-priority planes (20 A, 5 B, 5 C)", 20, 5, 5, 4, 6},
		{"More low-priority planes (5 A, 5 B, 20 C)", 5, 5, 20, 3, 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rand.Seed(42) // Fixed seed for reproducibility

			// Total number of planes
			totalPlanes := test.numTypeA + test.numTypeB + test.numTypeC

			// Initialize airport
			airport := newAirport(test.runways, test.gates, totalPlanes)

			// Generate planes for each category
			planes := make([]Plane, 0, totalPlanes)

			// Generate Category A planes
			for i := 0; i < test.numTypeA; i++ {
				planes = append(planes, Plane{id: len(planes) + 1, passengers: rand.Intn(50) + 101, category: "A", priority: 1})
			}

			// Generate Category B planes
			for i := 0; i < test.numTypeB; i++ {
				planes = append(planes, Plane{id: len(planes) + 1, passengers: rand.Intn(51) + 50, category: "B", priority: 2})
			}

			// Generate Category C planes
			for i := 0; i < test.numTypeC; i++ {
				planes = append(planes, Plane{id: len(planes) + 1, passengers: rand.Intn(50) + 1, category: "C", priority: 3})
			}

			// Shuffle planes for randomized arrival order
			rand.Shuffle(len(planes), func(i, j int) { planes[i], planes[j] = planes[j], planes[i] })

			// Sort planes by priority
			sort.Slice(planes, func(i, j int) bool {
				return planes[i].priority < planes[j].priority
			})

			// Start airport processing
			go airport.start()

			// Send planes to the queue
			for _, plane := range planes {
				airport.queue <- plane
			}
			close(airport.queue)

			// Wait for all planes to complete
			for i := 0; i < totalPlanes; i++ {
				<-airport.doneSignal
			}

			// Log completion for test case
			t.Logf("Test '%s' completed successfully with %d planes.", test.name, totalPlanes)
		})
	}
}
