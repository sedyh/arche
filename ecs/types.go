package ecs

import (
	"github.com/mlange-42/arche/internal/base"
)

// eid is the entity identifier/index type
type eid = base.Eid

// ID is the component identifier type
type ID = base.ID

// bitMask is a bitmask.
type bitMask = base.BitMask

// maskTotalBits is the size of Mask in bits.
//
// It is the maximum number of component types that may exist in any [World].
const maskTotalBits = base.MaskTotalBits

// Mask is a mask for a combination of components.
type Mask = base.Mask

// Component is a Component ID/Component pointer pair
type Component struct {
	ID
	Component interface{}
}
