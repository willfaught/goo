package goo

var _ Map = MapInt8Struct(nil)

// MapInt8Struct is a map from int8 to struct{}.
type MapInt8Struct map[int8]struct{}

// Delete implements Map.
func (m MapInt8Struct) Delete(k interface{}) {
	delete(m, k.(int8))
}

// Get implements Map.
func (m MapInt8Struct) Get(k interface{}) interface{} {
	return m[k.(int8)]
}

// GetCheck implements Map.
func (m MapInt8Struct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int8)]

	return v, ok
}

// KeyValues implements Map.
func (m MapInt8Struct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapInt8Struct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapInt8Struct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapInt8Struct) Make(c int) Map {
	return make(MapInt8Struct, c)
}

// Set implements Map.
func (m MapInt8Struct) Set(k, v interface{}) {
	m[k.(int8)] = v.(struct{})
}