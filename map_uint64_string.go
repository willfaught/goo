package goo

var _ Map = MapUint64String(nil)

// MapUint64String is a map from uint64 to string.
type MapUint64String map[uint64]string

// Delete implements Map.
func (m MapUint64String) Delete(k interface{}) {
	delete(m, k.(uint64))
}

// Get implements Map.
func (m MapUint64String) Get(k interface{}) interface{} {
	return m[k.(uint64)]
}

// GetCheck implements Map.
func (m MapUint64String) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uint64)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUint64String) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUint64String) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUint64String) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUint64String) Make(c int) Map {
	return make(MapUint64String, c)
}

// Set implements Map.
func (m MapUint64String) Set(k, v interface{}) {
	m[k.(uint64)] = v.(string)
}