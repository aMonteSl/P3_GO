package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	numRunways     = 3 // Number of runways
	numGates       = 5 // Number of gates
	maxQueueLength = 10
)

type Plane struct {
	id        int
	passengers int
	category  string
	priority  int
}

func newPlane(id int) Plane {
	passengers := rand.Intn(150) + 1 // Random passengers between 1 and 150
	var category string
	var priority int
	switch {
	case passengers > 100:
		category = "A"
		priority = 1
	case passengers >= 50:
		category = "B"
		priority = 2
	default:
		category = "C"
		priority = 3
	}
	return Plane{id, passengers, category, priority}
}

type Airport struct {
	runways    chan struct{}
	gates      chan struct{}
	queue      chan Plane
	doneSignal chan struct{}
}

func newAirport(numRunways, numGates int, numPlanes int) *Airport {
	return &Airport{
		runways:    make(chan struct{}, numRunways),
		gates:      make(chan struct{}, numGates),
		queue:      make(chan Plane, maxQueueLength),
		doneSignal: make(chan struct{}, numPlanes), // Track all planes
	}
}

func (a *Airport) handleLanding(plane Plane) {
	fmt.Printf("Plane %d (Category %s) is landing...\n", plane.id, plane.category)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+500)) // Simulate landing
}

func (a *Airport) handleGate(plane Plane) {
	fmt.Printf("Plane %d (Category %s) is unloading passengers...\n", plane.id, plane.category)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)+500)) // Simulate gate usage
	fmt.Printf("Plane %d disembarked.\n", plane.id)
}

func (a *Airport) processPlane(plane Plane) {
	// Landing phase
	a.runways <- struct{}{}
	a.handleLanding(plane)
	<-a.runways

	// Gate phase
	a.gates <- struct{}{}
	a.handleGate(plane)
	<-a.gates

	// Notify completion
	a.doneSignal <- struct{}{}
}

func (a *Airport) start() {
	for plane := range a.queue {
		go a.processPlane(plane)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numPlanes := 30
	airport := newAirport(numRunways, numGates, numPlanes)

	// Randomized plane creation
	planes := make([]Plane, 0, numPlanes)
	for i := 0; i < numPlanes; i++ {
		planes = append(planes, newPlane(i+1))
	}

	// Sorting planes by priority
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
	for i := 0; i < numPlanes; i++ {
		<-airport.doneSignal
	}

	fmt.Println("Simulation completed.")
}
