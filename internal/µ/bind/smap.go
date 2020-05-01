package bind

// sMap represents a type-safe drop-in replacement for sync.Map.
type sMap map[key]value

// newMap returns a new sMap
func newMap() sMap {
	return make(map[key]value)
}

// Clone returns a shallow copy.
func (m sMap) Clone() sMap {
	clone := newMap()
	for key, value := range m {
		clone[key] = value // Store(key, value)
	}
	return clone
}

// Store sets the value for a key.
func (m sMap) Store(key key, value value) {
	m[key] = value
}

// Load returns the value stored in the map for a key, or nil if no value is
// present. The ok result indicates whether value was found in the map.
func (m sMap) Load(key key) (value value, ok bool) {
	value, ok = m[key]
	return
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m sMap) LoadOrStore(key key, value value) (actual value, loaded bool) {
	actual, loaded = m[key] // m.Load(key)
	if !loaded {
		m[key] = value // m.Store(key, value)
		actual = value
	}
	return
}

// Delete deletes the value for a key.
func (m sMap) Delete(key key) {
	delete(m, key)
}

// =============================================================================
