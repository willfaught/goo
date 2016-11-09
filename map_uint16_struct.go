package goo

var _ Map = MapUint16Struct(nil)

// MapUint16Struct is a map from uint16 to struct{}.
type MapUint16Struct map[uint16]struct{}

// Delete implements Map.
func (m MapUint16Struct) Delete(k interface{}) {
	delete(m, k.(uint16))
}

// Get implements Map.
func (m MapUint16Struct) Get(k interface{}) interface{} {
	return m[k.(uint16)]
}

// GetCheck implements Map.
func (m MapUint16Struct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uint16)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUint16Struct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUint16Struct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUint16Struct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUint16Struct) Make(c int) Map {
	return make(MapUint16Struct, c)
}

// Set implements Map.
func (m MapUint16Struct) Set(k, v interface{}) {
	m[k.(uint16)] = v.(struct{})
}