package goo

var _ Map = MapFloat64Int(nil)

// MapFloat64Int is a map from float64 to int.
type MapFloat64Int map[float64]int

// Delete implements Map.
func (m MapFloat64Int) Delete(k interface{}) {
	delete(m, k.(float64))
}

// Get implements Map.
func (m MapFloat64Int) Get(k interface{}) interface{} {
	return m[k.(float64)]
}

// GetCheck implements Map.
func (m MapFloat64Int) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(float64)]

	return v, ok
}

// KeyValues implements Map.
func (m MapFloat64Int) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapFloat64Int) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapFloat64Int) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapFloat64Int) Make(c int) Map {
	return make(MapFloat64Int, c)
}

// Set implements Map.
func (m MapFloat64Int) Set(k, v interface{}) {
	m[k.(float64)] = v.(int)
}