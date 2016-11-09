package goo

var _ Map = MapUint8String(nil)

// MapUint8String is a map from uint8 to string.
type MapUint8String map[uint8]string

// Delete implements Map.
func (m MapUint8String) Delete(k interface{}) {
	delete(m, k.(uint8))
}

// Get implements Map.
func (m MapUint8String) Get(k interface{}) interface{} {
	return m[k.(uint8)]
}

// GetCheck implements Map.
func (m MapUint8String) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uint8)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUint8String) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUint8String) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUint8String) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUint8String) Make(c int) Map {
	return make(MapUint8String, c)
}

// Set implements Map.
func (m MapUint8String) Set(k, v interface{}) {
	m[k.(uint8)] = v.(string)
}