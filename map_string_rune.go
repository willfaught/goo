package goo

var _ Map = MapStringRune(nil)

// MapStringRune is a map from string to rune.
type MapStringRune map[string]rune

// Delete implements Map.
func (m MapStringRune) Delete(k interface{}) {
	delete(m, k.(string))
}

// Get implements Map.
func (m MapStringRune) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringRune) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringRune) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringRune) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringRune) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringRune) Make(c int) Map {
	return make(MapStringRune, c)
}

// Set implements Map.
func (m MapStringRune) Set(k, v interface{}) {
	m[k.(string)] = v.(rune)
}