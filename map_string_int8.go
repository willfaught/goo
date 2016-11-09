package goo

var _ Map = MapStringInt8(nil)

// MapStringInt8 is a map from string to int8.
type MapStringInt8 map[string]int8

// Delete implements Map.
func (m MapStringInt8) Delete(k interface{}) {
	delete(m, k.(string))
}

// Get implements Map.
func (m MapStringInt8) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringInt8) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringInt8) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringInt8) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringInt8) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringInt8) Make(c int) Map {
	return make(MapStringInt8, c)
}

// Set implements Map.
func (m MapStringInt8) Set(k, v interface{}) {
	m[k.(string)] = v.(int8)
}