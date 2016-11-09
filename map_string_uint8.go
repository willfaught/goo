package goo

var _ Map = MapStringUint8(nil)

// MapStringUint8 is a map from string to uint8.
type MapStringUint8 map[string]uint8

// Delete implements Map.
func (m MapStringUint8) Delete(k interface{}) {
	delete(m, k.(string))
}

// Get implements Map.
func (m MapStringUint8) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringUint8) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringUint8) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringUint8) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringUint8) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringUint8) Make(c int) Map {
	return make(MapStringUint8, c)
}

// Set implements Map.
func (m MapStringUint8) Set(k, v interface{}) {
	m[k.(string)] = v.(uint8)
}