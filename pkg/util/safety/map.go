package safety

import (
	"sync"
)

// NewMap Create a thread-safe map.
func NewMap[K comparable, T any]() *Map[K, T] {
	return &Map[K, T]{
		m: new(sync.Map),
	}
}

// Map a thread-safe map.
type Map[K comparable, T any] struct {
	m *sync.Map
}

// Get Retrieve the value from the map.
func (m *Map[K, T]) Get(key K) (T, bool) {
	v, ok := m.m.Load(key)
	if !ok {
		var zero T
		return zero, false
	}
	return v.(T), true
}

// Set the value in the map.
func (m *Map[K, T]) Set(key K, value T) {
	m.m.Store(key, value)
}

// Delete the value from the map.
func (m *Map[K, T]) Delete(key K) {
	m.m.Delete(key)
}

// List Retrieve all values from the map.
func (m *Map[K, T]) List() map[K]T {
	values := make(map[K]T)
	m.m.Range(func(key, value any) bool {
		values[key.(K)] = value.(T)
		return true
	})
	return values
}

// Clear the map.
func (m *Map[K, T]) Clear() {
	m.m.Clear()
}
