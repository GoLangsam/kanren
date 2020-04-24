package smap

// SMap represents the mapping of logic Variables to symbolic eXpressions/terms
// and intentionally mimics and extends the interface of a sync.Map.
type SMap map[V]X

func New() SMap {
	return make(map[V]X)
}

// Clone returns a shallow copy.
func (m SMap) Clone() SMap {
	clone := New()
	for key, value := range m {
		clone[key] = value // Store(key, value)
	}
	return clone
}

// Store sets the value for a key.
func (m SMap) Store(key V, value X) {
	m[key] = value
}

// Delete deletes the value for a key.
func (m SMap) Delete(key V) {
	delete(m, key)
}

// Load returns the value stored in the map for a key, or nil if no value is
// present. The ok result indicates whether value was found in the map.
func (m SMap) Load(key V) (value X, ok bool) {
	value, ok = m[key]
	return
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m SMap) LoadOrStore(key V, value X) (actual X, loaded bool) {
	actual, loaded = m[key] // m.Load(key)
	if !loaded {
		m[key] = value // m.Store(key, value)
		actual = value
	}
	return
}

// =============================================================================
