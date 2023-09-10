package iter

import "reflect"

// MapEntry represents a key-value pair from a map.
type MapEntry[K comparable, V any] struct {
	Key   K
	Value V
}

// FromMap creates an iterator from a map.
func FromMap[K comparable, V any](m map[K]V) Iter[MapEntry[K, V]] {
	r := reflect.ValueOf(m).MapRange()
	return IterFunc[MapEntry[K, V]](func() (MapEntry[K, V], bool) {
		if !r.Next() {
			return zero[MapEntry[K, V]](), false
		}
		return MapEntry[K, V]{
			Key:   r.Key().Interface().(K),
			Value: r.Value().Interface().(V),
		}, true
	})
}

// Keys returns an iterator for the keys of the map.
func Keys[K comparable, V any](m map[K]V) Iter[K] {
	return Map(FromMap(m), func(e MapEntry[K, V]) K { return e.Key })
}

// Values returns an iterator for the values of the map.
func Values[K comparable, V any](m map[K]V) Iter[V] {
	return Map(FromMap(m), func(e MapEntry[K, V]) V { return e.Value })
}
