// Demonstrates the generic API, which provides type-safety and convenience over the ID-based core API.
package main

import (
	"math/rand"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// Position component
type Position struct {
	X float64
	Y float64
}

// Velocity component
type Velocity struct {
	X float64
	Y float64
}

// Rotation component
type Rotation struct {
	A float64
}

// Elevation component
type Elevation struct {
	E float64
}

func main() {
	// Create a World.
	world := ecs.NewWorld()

	// Create a generic entity mutation helper.
	mutator :=
		generic.NewMutate(&world).
			WithAdd(generic.T4[Position, Velocity, Rotation, Elevation]()...)
	// Create a component mapper.
	mapper := generic.NewMap2[Position, Velocity](&world)

	// Create entities.
	for i := 0; i < 1000; i++ {
		// Create a new Entity with components.
		e := mutator.NewEntity()
		pos, vel := mapper.Get(e)

		// Initialize component fields.
		pos.X = rand.Float64() * 100
		pos.Y = rand.Float64() * 100

		vel.X = rand.NormFloat64()
		vel.Y = rand.NormFloat64()
	}

	// Generic queries support up to 8 components.
	// For more components, use World.Query() directly.
	filter := generic.NewFilter2[Position, Velocity]()

	// Time loop.
	for t := 0; t < 1000; t++ {
		// Get a fresh query iterator.
		query := filter.Query(&world)
		// Iterate it.
		for query.Next() {
			// Component access through the Query.
			_, pos, vel := query.GetAll()
			// Update component fields.
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}

	// A more complex generic query using optional and excluded components:
	filter =
		generic.
			NewFilter2[Position, Velocity](). // Components provided through Get... methods
			Optional(generic.T[Velocity]()).  // but those may be nil
			With(generic.T[Elevation]()).     // additional required components
			Without(generic.T[Rotation]())    // entities with any of these are excluded.

	q := filter.Query(&world)

	for q.Next() {
		_, pos, vel := q.GetAll()
		pos.X += vel.X
		pos.Y += vel.Y
	}
}
