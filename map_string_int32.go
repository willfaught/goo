package goo

var _ Map = MapStringInt32(nil)

// MapStringInt32 is a map from string to int32.
type MapStringInt32 map[string]int32

// Delete implements Map.
func (m MapStringInt32) Delete(k interface{}) {
	delete(m, k.(string))
}

// Get implements Map.
func (m MapStringInt32) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringInt32) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringInt32) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringInt32) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringInt32) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringInt32) Make(c int) Map {
	return make(MapStringInt32, c)
}

// Set implements Map.
func (m MapStringInt32) Set(k, v interface{}) {
	m[k.(string)] = v.(int32)
}