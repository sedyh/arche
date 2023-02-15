package generic

// Code generated by go generate; DO NOT EDIT.

import (
	"reflect"

	"github.com/mlange-42/arche/ecs"
)

//////////////////////////////////////////////////////////////////////////

// Filter0 builds a [Query0] query
type Filter0 struct {
	include  []reflect.Type
	optional []reflect.Type
	exclude  []reflect.Type
	compiled compiledQuery
}

// NewQuery0 creates a generic filter for two components.
//
// See also [ecs.World.Query].
func NewFilter0() *Filter0 {
	return &Filter0{}
}

// Optional makes some of the query's components optional.
//
// Create the required mask with [Mask1], [Mask2], etc.
//
// Only affects component types that were specified in the query.
func (q *Filter0) Optional(mask []reflect.Type) *Filter0 {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds more required components that are not accessible using Get... methods.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter0) With(mask []reflect.Type) *Filter0 {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter0) Without(mask []reflect.Type) *Filter0 {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query0] query for iteration.
func (q *Filter0) Query(w *ecs.World) Query0 {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query0{
		w.Query(q.compiled.Filter()),
		q.compiled.Ids,
	}
}

// Filter generates and return the [ecs.Filter] used after [Filter0.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter0.Query].
func (q *Filter0) Filter(w *ecs.World) ecs.MaskPair {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.Filter()
}

// Query0 is a generic query iterator for two components.
//
// Create one with [NewFilter0] and [Filter0.Query]
type Query0 struct {
	ecs.Query
	ids []ecs.ID
}

//////////////////////////////////////////////////////////////////////////

// Filter1 builds a [Query1] query
type Filter1[A any] struct {
	include  []reflect.Type
	optional []reflect.Type
	exclude  []reflect.Type
	compiled compiledQuery
}

// NewQuery1 creates a generic filter for two components.
//
// See also [ecs.World.Query].
func NewFilter1[A any]() *Filter1[A] {
	return &Filter1[A]{
		include: []reflect.Type{typeOf[A]()},
	}
}

// Optional makes some of the query's components optional.
//
// Create the required mask with [Mask1], [Mask2], etc.
//
// Only affects component types that were specified in the query.
func (q *Filter1[A]) Optional(mask []reflect.Type) *Filter1[A] {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds more required components that are not accessible using Get... methods.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter1[A]) With(mask []reflect.Type) *Filter1[A] {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter1[A]) Without(mask []reflect.Type) *Filter1[A] {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query1] query for iteration.
func (q *Filter1[A]) Query(w *ecs.World) Query1[A] {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query1[A]{
		w.Query(q.compiled.Filter()),
		q.compiled.Ids,
	}
}

// Filter generates and return the [ecs.Filter] used after [Filter1.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter1.Query].
func (q *Filter1[A]) Filter(w *ecs.World) ecs.MaskPair {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.Filter()
}

// Query1 is a generic query iterator for two components.
//
// Create one with [NewFilter1] and [Filter1.Query]
type Query1[A any] struct {
	ecs.Query
	ids []ecs.ID
}

// GetAll returns the [ecs.Entity] and all queried components for the current query iterator position.
func (q *Query1[A]) GetAll() (ecs.Entity, *A) {
	return q.Entity(), (*A)(q.Query.Get(q.ids[0]))
}

// Get1 returns the first queried component for the current query position
func (q *Query1[A]) Get1() *A {
	return (*A)(q.Query.Get(q.ids[0]))
}

//////////////////////////////////////////////////////////////////////////

// Filter2 builds a [Query2] query
type Filter2[A any, B any] struct {
	include  []reflect.Type
	optional []reflect.Type
	exclude  []reflect.Type
	compiled compiledQuery
}

// NewQuery2 creates a generic filter for two components.
//
// See also [ecs.World.Query].
func NewFilter2[A any, B any]() *Filter2[A, B] {
	return &Filter2[A, B]{
		include: []reflect.Type{typeOf[A](), typeOf[B]()},
	}
}

// Optional makes some of the query's components optional.
//
// Create the required mask with [Mask1], [Mask2], etc.
//
// Only affects component types that were specified in the query.
func (q *Filter2[A, B]) Optional(mask []reflect.Type) *Filter2[A, B] {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds more required components that are not accessible using Get... methods.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter2[A, B]) With(mask []reflect.Type) *Filter2[A, B] {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter2[A, B]) Without(mask []reflect.Type) *Filter2[A, B] {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query2] query for iteration.
func (q *Filter2[A, B]) Query(w *ecs.World) Query2[A, B] {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query2[A, B]{
		w.Query(q.compiled.Filter()),
		q.compiled.Ids,
	}
}

// Filter generates and return the [ecs.Filter] used after [Filter2.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter2.Query].
func (q *Filter2[A, B]) Filter(w *ecs.World) ecs.MaskPair {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.Filter()
}

// Query2 is a generic query iterator for two components.
//
// Create one with [NewFilter2] and [Filter2.Query]
type Query2[A any, B any] struct {
	ecs.Query
	ids []ecs.ID
}

// GetAll returns the [ecs.Entity] and all queried components for the current query iterator position.
func (q *Query2[A, B]) GetAll() (ecs.Entity, *A, *B) {
	return q.Entity(), (*A)(q.Query.Get(q.ids[0])), (*B)(q.Query.Get(q.ids[1]))
}

// Get1 returns the first queried component for the current query position
func (q *Query2[A, B]) Get1() *A {
	return (*A)(q.Query.Get(q.ids[0]))
}

// Get2 returns the first queried component for the current query position
func (q *Query2[A, B]) Get2() *B {
	return (*B)(q.Query.Get(q.ids[1]))
}

//////////////////////////////////////////////////////////////////////////

// Filter3 builds a [Query3] query
type Filter3[A any, B any, C any] struct {
	include  []reflect.Type
	optional []reflect.Type
	exclude  []reflect.Type
	compiled compiledQuery
}

// NewQuery3 creates a generic filter for two components.
//
// See also [ecs.World.Query].
func NewFilter3[A any, B any, C any]() *Filter3[A, B, C] {
	return &Filter3[A, B, C]{
		include: []reflect.Type{typeOf[A](), typeOf[B](), typeOf[C]()},
	}
}

// Optional makes some of the query's components optional.
//
// Create the required mask with [Mask1], [Mask2], etc.
//
// Only affects component types that were specified in the query.
func (q *Filter3[A, B, C]) Optional(mask []reflect.Type) *Filter3[A, B, C] {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds more required components that are not accessible using Get... methods.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter3[A, B, C]) With(mask []reflect.Type) *Filter3[A, B, C] {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter3[A, B, C]) Without(mask []reflect.Type) *Filter3[A, B, C] {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query3] query for iteration.
func (q *Filter3[A, B, C]) Query(w *ecs.World) Query3[A, B, C] {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query3[A, B, C]{
		w.Query(q.compiled.Filter()),
		q.compiled.Ids,
	}
}

// Filter generates and return the [ecs.Filter] used after [Filter3.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter3.Query].
func (q *Filter3[A, B, C]) Filter(w *ecs.World) ecs.MaskPair {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.Filter()
}

// Query3 is a generic query iterator for two components.
//
// Create one with [NewFilter3] and [Filter3.Query]
type Query3[A any, B any, C any] struct {
	ecs.Query
	ids []ecs.ID
}

// GetAll returns the [ecs.Entity] and all queried components for the current query iterator position.
func (q *Query3[A, B, C]) GetAll() (ecs.Entity, *A, *B, *C) {
	return q.Entity(), (*A)(q.Query.Get(q.ids[0])), (*B)(q.Query.Get(q.ids[1])), (*C)(q.Query.Get(q.ids[2]))
}

// Get1 returns the first queried component for the current query position
func (q *Query3[A, B, C]) Get1() *A {
	return (*A)(q.Query.Get(q.ids[0]))
}

// Get2 returns the first queried component for the current query position
func (q *Query3[A, B, C]) Get2() *B {
	return (*B)(q.Query.Get(q.ids[1]))
}

// Get3 returns the first queried component for the current query position
func (q *Query3[A, B, C]) Get3() *C {
	return (*C)(q.Query.Get(q.ids[2]))
}

//////////////////////////////////////////////////////////////////////////

// Filter4 builds a [Query4] query
type Filter4[A any, B any, C any, D any] struct {
	include  []reflect.Type
	optional []reflect.Type
	exclude  []reflect.Type
	compiled compiledQuery
}

// NewQuery4 creates a generic filter for two components.
//
// See also [ecs.World.Query].
func NewFilter4[A any, B any, C any, D any]() *Filter4[A, B, C, D] {
	return &Filter4[A, B, C, D]{
		include: []reflect.Type{typeOf[A](), typeOf[B](), typeOf[C](), typeOf[D]()},
	}
}

// Optional makes some of the query's components optional.
//
// Create the required mask with [Mask1], [Mask2], etc.
//
// Only affects component types that were specified in the query.
func (q *Filter4[A, B, C, D]) Optional(mask []reflect.Type) *Filter4[A, B, C, D] {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds more required components that are not accessible using Get... methods.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter4[A, B, C, D]) With(mask []reflect.Type) *Filter4[A, B, C, D] {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter4[A, B, C, D]) Without(mask []reflect.Type) *Filter4[A, B, C, D] {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query4] query for iteration.
func (q *Filter4[A, B, C, D]) Query(w *ecs.World) Query4[A, B, C, D] {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query4[A, B, C, D]{
		w.Query(q.compiled.Filter()),
		q.compiled.Ids,
	}
}

// Filter generates and return the [ecs.Filter] used after [Filter4.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter4.Query].
func (q *Filter4[A, B, C, D]) Filter(w *ecs.World) ecs.MaskPair {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.Filter()
}

// Query4 is a generic query iterator for two components.
//
// Create one with [NewFilter4] and [Filter4.Query]
type Query4[A any, B any, C any, D any] struct {
	ecs.Query
	ids []ecs.ID
}

// GetAll returns the [ecs.Entity] and all queried components for the current query iterator position.
func (q *Query4[A, B, C, D]) GetAll() (ecs.Entity, *A, *B, *C, *D) {
	return q.Entity(), (*A)(q.Query.Get(q.ids[0])), (*B)(q.Query.Get(q.ids[1])), (*C)(q.Query.Get(q.ids[2])), (*D)(q.Query.Get(q.ids[3]))
}

// Get1 returns the first queried component for the current query position
func (q *Query4[A, B, C, D]) Get1() *A {
	return (*A)(q.Query.Get(q.ids[0]))
}

// Get2 returns the first queried component for the current query position
func (q *Query4[A, B, C, D]) Get2() *B {
	return (*B)(q.Query.Get(q.ids[1]))
}

// Get3 returns the first queried component for the current query position
func (q *Query4[A, B, C, D]) Get3() *C {
	return (*C)(q.Query.Get(q.ids[2]))
}

// Get4 returns the first queried component for the current query position
func (q *Query4[A, B, C, D]) Get4() *D {
	return (*D)(q.Query.Get(q.ids[3]))
}

//////////////////////////////////////////////////////////////////////////

// Filter5 builds a [Query5] query
type Filter5[A any, B any, C any, D any, E any] struct {
	include  []reflect.Type
	optional []reflect.Type
	exclude  []reflect.Type
	compiled compiledQuery
}

// NewQuery5 creates a generic filter for two components.
//
// See also [ecs.World.Query].
func NewFilter5[A any, B any, C any, D any, E any]() *Filter5[A, B, C, D, E] {
	return &Filter5[A, B, C, D, E]{
		include: []reflect.Type{typeOf[A](), typeOf[B](), typeOf[C](), typeOf[D](), typeOf[E]()},
	}
}

// Optional makes some of the query's components optional.
//
// Create the required mask with [Mask1], [Mask2], etc.
//
// Only affects component types that were specified in the query.
func (q *Filter5[A, B, C, D, E]) Optional(mask []reflect.Type) *Filter5[A, B, C, D, E] {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds more required components that are not accessible using Get... methods.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter5[A, B, C, D, E]) With(mask []reflect.Type) *Filter5[A, B, C, D, E] {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter5[A, B, C, D, E]) Without(mask []reflect.Type) *Filter5[A, B, C, D, E] {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query5] query for iteration.
func (q *Filter5[A, B, C, D, E]) Query(w *ecs.World) Query5[A, B, C, D, E] {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query5[A, B, C, D, E]{
		w.Query(q.compiled.Filter()),
		q.compiled.Ids,
	}
}

// Filter generates and return the [ecs.Filter] used after [Filter5.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter5.Query].
func (q *Filter5[A, B, C, D, E]) Filter(w *ecs.World) ecs.MaskPair {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.Filter()
}

// Query5 is a generic query iterator for two components.
//
// Create one with [NewFilter5] and [Filter5.Query]
type Query5[A any, B any, C any, D any, E any] struct {
	ecs.Query
	ids []ecs.ID
}

// GetAll returns the [ecs.Entity] and all queried components for the current query iterator position.
func (q *Query5[A, B, C, D, E]) GetAll() (ecs.Entity, *A, *B, *C, *D, *E) {
	return q.Entity(), (*A)(q.Query.Get(q.ids[0])), (*B)(q.Query.Get(q.ids[1])), (*C)(q.Query.Get(q.ids[2])), (*D)(q.Query.Get(q.ids[3])), (*E)(q.Query.Get(q.ids[4]))
}

// Get1 returns the first queried component for the current query position
func (q *Query5[A, B, C, D, E]) Get1() *A {
	return (*A)(q.Query.Get(q.ids[0]))
}

// Get2 returns the first queried component for the current query position
func (q *Query5[A, B, C, D, E]) Get2() *B {
	return (*B)(q.Query.Get(q.ids[1]))
}

// Get3 returns the first queried component for the current query position
func (q *Query5[A, B, C, D, E]) Get3() *C {
	return (*C)(q.Query.Get(q.ids[2]))
}

// Get4 returns the first queried component for the current query position
func (q *Query5[A, B, C, D, E]) Get4() *D {
	return (*D)(q.Query.Get(q.ids[3]))
}

// Get5 returns the first queried component for the current query position
func (q *Query5[A, B, C, D, E]) Get5() *E {
	return (*E)(q.Query.Get(q.ids[4]))
}

//////////////////////////////////////////////////////////////////////////

// Filter6 builds a [Query6] query
type Filter6[A any, B any, C any, D any, E any, F any] struct {
	include  []reflect.Type
	optional []reflect.Type
	exclude  []reflect.Type
	compiled compiledQuery
}

// NewQuery6 creates a generic filter for two components.
//
// See also [ecs.World.Query].
func NewFilter6[A any, B any, C any, D any, E any, F any]() *Filter6[A, B, C, D, E, F] {
	return &Filter6[A, B, C, D, E, F]{
		include: []reflect.Type{typeOf[A](), typeOf[B](), typeOf[C](), typeOf[D](), typeOf[E](), typeOf[F]()},
	}
}

// Optional makes some of the query's components optional.
//
// Create the required mask with [Mask1], [Mask2], etc.
//
// Only affects component types that were specified in the query.
func (q *Filter6[A, B, C, D, E, F]) Optional(mask []reflect.Type) *Filter6[A, B, C, D, E, F] {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds more required components that are not accessible using Get... methods.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter6[A, B, C, D, E, F]) With(mask []reflect.Type) *Filter6[A, B, C, D, E, F] {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter6[A, B, C, D, E, F]) Without(mask []reflect.Type) *Filter6[A, B, C, D, E, F] {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query6] query for iteration.
func (q *Filter6[A, B, C, D, E, F]) Query(w *ecs.World) Query6[A, B, C, D, E, F] {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query6[A, B, C, D, E, F]{
		w.Query(q.compiled.Filter()),
		q.compiled.Ids,
	}
}

// Filter generates and return the [ecs.Filter] used after [Filter6.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter6.Query].
func (q *Filter6[A, B, C, D, E, F]) Filter(w *ecs.World) ecs.MaskPair {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.Filter()
}

// Query6 is a generic query iterator for two components.
//
// Create one with [NewFilter6] and [Filter6.Query]
type Query6[A any, B any, C any, D any, E any, F any] struct {
	ecs.Query
	ids []ecs.ID
}

// GetAll returns the [ecs.Entity] and all queried components for the current query iterator position.
func (q *Query6[A, B, C, D, E, F]) GetAll() (ecs.Entity, *A, *B, *C, *D, *E, *F) {
	return q.Entity(), (*A)(q.Query.Get(q.ids[0])), (*B)(q.Query.Get(q.ids[1])), (*C)(q.Query.Get(q.ids[2])), (*D)(q.Query.Get(q.ids[3])), (*E)(q.Query.Get(q.ids[4])), (*F)(q.Query.Get(q.ids[5]))
}

// Get1 returns the first queried component for the current query position
func (q *Query6[A, B, C, D, E, F]) Get1() *A {
	return (*A)(q.Query.Get(q.ids[0]))
}

// Get2 returns the first queried component for the current query position
func (q *Query6[A, B, C, D, E, F]) Get2() *B {
	return (*B)(q.Query.Get(q.ids[1]))
}

// Get3 returns the first queried component for the current query position
func (q *Query6[A, B, C, D, E, F]) Get3() *C {
	return (*C)(q.Query.Get(q.ids[2]))
}

// Get4 returns the first queried component for the current query position
func (q *Query6[A, B, C, D, E, F]) Get4() *D {
	return (*D)(q.Query.Get(q.ids[3]))
}

// Get5 returns the first queried component for the current query position
func (q *Query6[A, B, C, D, E, F]) Get5() *E {
	return (*E)(q.Query.Get(q.ids[4]))
}

// Get6 returns the first queried component for the current query position
func (q *Query6[A, B, C, D, E, F]) Get6() *F {
	return (*F)(q.Query.Get(q.ids[5]))
}

//////////////////////////////////////////////////////////////////////////

// Filter7 builds a [Query7] query
type Filter7[A any, B any, C any, D any, E any, F any, G any] struct {
	include  []reflect.Type
	optional []reflect.Type
	exclude  []reflect.Type
	compiled compiledQuery
}

// NewQuery7 creates a generic filter for two components.
//
// See also [ecs.World.Query].
func NewFilter7[A any, B any, C any, D any, E any, F any, G any]() *Filter7[A, B, C, D, E, F, G] {
	return &Filter7[A, B, C, D, E, F, G]{
		include: []reflect.Type{typeOf[A](), typeOf[B](), typeOf[C](), typeOf[D](), typeOf[E](), typeOf[F](), typeOf[G]()},
	}
}

// Optional makes some of the query's components optional.
//
// Create the required mask with [Mask1], [Mask2], etc.
//
// Only affects component types that were specified in the query.
func (q *Filter7[A, B, C, D, E, F, G]) Optional(mask []reflect.Type) *Filter7[A, B, C, D, E, F, G] {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds more required components that are not accessible using Get... methods.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter7[A, B, C, D, E, F, G]) With(mask []reflect.Type) *Filter7[A, B, C, D, E, F, G] {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter7[A, B, C, D, E, F, G]) Without(mask []reflect.Type) *Filter7[A, B, C, D, E, F, G] {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query7] query for iteration.
func (q *Filter7[A, B, C, D, E, F, G]) Query(w *ecs.World) Query7[A, B, C, D, E, F, G] {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query7[A, B, C, D, E, F, G]{
		w.Query(q.compiled.Filter()),
		q.compiled.Ids,
	}
}

// Filter generates and return the [ecs.Filter] used after [Filter7.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter7.Query].
func (q *Filter7[A, B, C, D, E, F, G]) Filter(w *ecs.World) ecs.MaskPair {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.Filter()
}

// Query7 is a generic query iterator for two components.
//
// Create one with [NewFilter7] and [Filter7.Query]
type Query7[A any, B any, C any, D any, E any, F any, G any] struct {
	ecs.Query
	ids []ecs.ID
}

// GetAll returns the [ecs.Entity] and all queried components for the current query iterator position.
func (q *Query7[A, B, C, D, E, F, G]) GetAll() (ecs.Entity, *A, *B, *C, *D, *E, *F, *G) {
	return q.Entity(), (*A)(q.Query.Get(q.ids[0])), (*B)(q.Query.Get(q.ids[1])), (*C)(q.Query.Get(q.ids[2])), (*D)(q.Query.Get(q.ids[3])), (*E)(q.Query.Get(q.ids[4])), (*F)(q.Query.Get(q.ids[5])), (*G)(q.Query.Get(q.ids[6]))
}

// Get1 returns the first queried component for the current query position
func (q *Query7[A, B, C, D, E, F, G]) Get1() *A {
	return (*A)(q.Query.Get(q.ids[0]))
}

// Get2 returns the first queried component for the current query position
func (q *Query7[A, B, C, D, E, F, G]) Get2() *B {
	return (*B)(q.Query.Get(q.ids[1]))
}

// Get3 returns the first queried component for the current query position
func (q *Query7[A, B, C, D, E, F, G]) Get3() *C {
	return (*C)(q.Query.Get(q.ids[2]))
}

// Get4 returns the first queried component for the current query position
func (q *Query7[A, B, C, D, E, F, G]) Get4() *D {
	return (*D)(q.Query.Get(q.ids[3]))
}

// Get5 returns the first queried component for the current query position
func (q *Query7[A, B, C, D, E, F, G]) Get5() *E {
	return (*E)(q.Query.Get(q.ids[4]))
}

// Get6 returns the first queried component for the current query position
func (q *Query7[A, B, C, D, E, F, G]) Get6() *F {
	return (*F)(q.Query.Get(q.ids[5]))
}

// Get7 returns the first queried component for the current query position
func (q *Query7[A, B, C, D, E, F, G]) Get7() *G {
	return (*G)(q.Query.Get(q.ids[6]))
}

//////////////////////////////////////////////////////////////////////////

// Filter8 builds a [Query8] query
type Filter8[A any, B any, C any, D any, E any, F any, G any, H any] struct {
	include  []reflect.Type
	optional []reflect.Type
	exclude  []reflect.Type
	compiled compiledQuery
}

// NewQuery8 creates a generic filter for two components.
//
// See also [ecs.World.Query].
func NewFilter8[A any, B any, C any, D any, E any, F any, G any, H any]() *Filter8[A, B, C, D, E, F, G, H] {
	return &Filter8[A, B, C, D, E, F, G, H]{
		include: []reflect.Type{typeOf[A](), typeOf[B](), typeOf[C](), typeOf[D](), typeOf[E](), typeOf[F](), typeOf[G](), typeOf[H]()},
	}
}

// Optional makes some of the query's components optional.
//
// Create the required mask with [Mask1], [Mask2], etc.
//
// Only affects component types that were specified in the query.
func (q *Filter8[A, B, C, D, E, F, G, H]) Optional(mask []reflect.Type) *Filter8[A, B, C, D, E, F, G, H] {
	q.optional = append(q.optional, mask...)
	q.compiled.Reset()
	return q
}

// With adds more required components that are not accessible using Get... methods.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter8[A, B, C, D, E, F, G, H]) With(mask []reflect.Type) *Filter8[A, B, C, D, E, F, G, H] {
	q.include = append(q.include, mask...)
	q.compiled.Reset()
	return q
}

// Without excludes entities with any of the given components from the query.
//
// Create the required mask with [Mask1], [Mask2], etc.
func (q *Filter8[A, B, C, D, E, F, G, H]) Without(mask []reflect.Type) *Filter8[A, B, C, D, E, F, G, H] {
	q.exclude = append(q.exclude, mask...)
	q.compiled.Reset()
	return q
}

// Query builds a [Query8] query for iteration.
func (q *Filter8[A, B, C, D, E, F, G, H]) Query(w *ecs.World) Query8[A, B, C, D, E, F, G, H] {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return Query8[A, B, C, D, E, F, G, H]{
		w.Query(q.compiled.Filter()),
		q.compiled.Ids,
	}
}

// Filter generates and return the [ecs.Filter] used after [Filter8.Query].
//
// Can be passed to [ecs.World.Query].
// For the intended generic use, however, generate a generic query with [Filter8.Query].
func (q *Filter8[A, B, C, D, E, F, G, H]) Filter(w *ecs.World) ecs.MaskPair {
	q.compiled.Compile(w, q.include, q.optional, q.exclude)
	return q.compiled.Filter()
}

// Query8 is a generic query iterator for two components.
//
// Create one with [NewFilter8] and [Filter8.Query]
type Query8[A any, B any, C any, D any, E any, F any, G any, H any] struct {
	ecs.Query
	ids []ecs.ID
}

// GetAll returns the [ecs.Entity] and all queried components for the current query iterator position.
func (q *Query8[A, B, C, D, E, F, G, H]) GetAll() (ecs.Entity, *A, *B, *C, *D, *E, *F, *G, *H) {
	return q.Entity(), (*A)(q.Query.Get(q.ids[0])), (*B)(q.Query.Get(q.ids[1])), (*C)(q.Query.Get(q.ids[2])), (*D)(q.Query.Get(q.ids[3])), (*E)(q.Query.Get(q.ids[4])), (*F)(q.Query.Get(q.ids[5])), (*G)(q.Query.Get(q.ids[6])), (*H)(q.Query.Get(q.ids[7]))
}

// Get1 returns the first queried component for the current query position
func (q *Query8[A, B, C, D, E, F, G, H]) Get1() *A {
	return (*A)(q.Query.Get(q.ids[0]))
}

// Get2 returns the first queried component for the current query position
func (q *Query8[A, B, C, D, E, F, G, H]) Get2() *B {
	return (*B)(q.Query.Get(q.ids[1]))
}

// Get3 returns the first queried component for the current query position
func (q *Query8[A, B, C, D, E, F, G, H]) Get3() *C {
	return (*C)(q.Query.Get(q.ids[2]))
}

// Get4 returns the first queried component for the current query position
func (q *Query8[A, B, C, D, E, F, G, H]) Get4() *D {
	return (*D)(q.Query.Get(q.ids[3]))
}

// Get5 returns the first queried component for the current query position
func (q *Query8[A, B, C, D, E, F, G, H]) Get5() *E {
	return (*E)(q.Query.Get(q.ids[4]))
}

// Get6 returns the first queried component for the current query position
func (q *Query8[A, B, C, D, E, F, G, H]) Get6() *F {
	return (*F)(q.Query.Get(q.ids[5]))
}

// Get7 returns the first queried component for the current query position
func (q *Query8[A, B, C, D, E, F, G, H]) Get7() *G {
	return (*G)(q.Query.Get(q.ids[6]))
}

// Get8 returns the first queried component for the current query position
func (q *Query8[A, B, C, D, E, F, G, H]) Get8() *H {
	return (*H)(q.Query.Get(q.ids[7]))
}
