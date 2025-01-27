package ecs

const (
	numChunks = 8
	chunkSize = 16
)

// idMap maps component IDs to values.
//
// Is is a data structure meant for fast lookup while being memory-efficient.
// Access time is around 2ns, compared to 0.5ns for array access and 20ns for map[int]T.
//
// The memory footprint is reduced by using chunks, and only allocating chunks if they contain a key.
//
// The range of keys is limited from 0 to [MaskTotalBits]-1.
type idMap[T any] struct {
	chunks    [][]T
	used      Mask
	chunkUsed []uint8
	zeroValue T
}

// newIDMap creates a new idMap
func newIDMap[T any]() idMap[T] {
	return idMap[T]{
		chunks:    make([][]T, numChunks),
		used:      Mask{},
		chunkUsed: make([]uint8, numChunks),
	}
}

// Get returns the value at the given key and whether the key is present.
func (m *idMap[T]) Get(index uint8) (T, bool) {
	if !m.used.Get(index) {
		return m.zeroValue, false
	}
	return m.chunks[index/chunkSize][index%chunkSize], true
}

// Get returns a pointer to the value at the given key and whether the key is present.
func (m *idMap[T]) GetPointer(index uint8) (*T, bool) {
	if !m.used.Get(index) {
		return nil, false
	}
	return &m.chunks[index/chunkSize][index%chunkSize], true
}

// Set sets the value at the given key.
func (m *idMap[T]) Set(index uint8, value T) {
	chunk := index / chunkSize
	if m.chunks[chunk] == nil {
		m.chunks[chunk] = make([]T, chunkSize)
	}
	m.chunks[chunk][index%chunkSize] = value
	m.used.Set(index, true)
	m.chunkUsed[chunk]++
}

// Remove removes the value at the given key.
// It de-allocates empty chunks.
func (m *idMap[T]) Remove(index uint8) {
	chunk := index / chunkSize
	m.used.Set(index, false)
	m.chunkUsed[chunk]--
	if m.chunkUsed[chunk] == 0 {
		m.chunks[chunk] = nil
	}
}
