package maps

import "iter"

// Keys returns a slice of all keys in m
func Keys[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range m {
			if !yield(k) {
				return
			}
		}
	}
}

// Values returns a slice of all values in m
func Values[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range m {
			if !yield(v) {
				return
			}
		}
	}
}

// Union returns a new map of all key/value pairs in left and right. If a key exists
// in both left and right the value in right will appear in the resultant map.
func Union[Map ~map[K]V, K comparable, V any](left, right Map) Map {
	result := make(Map)
	for k, v := range left {
		result[k] = v
	}

	for k, v := range right {
		result[k] = v
	}

	return result
}

// UnionInPlace modifies left to include key/value pairs in right.  If a key exists
// in both left and right the value in right will be used.
func UnionInPlace[Map ~map[K]V, K comparable, V any](left, right Map) Map {
	for k, v := range right {
		left[k] = v
	}

	return left
}

// Intersect returns a new map of key/value pairs where the key exists in both left and right.
// The value from right will be used in the return map.
func Intersect[Map ~map[K]V, K comparable, V any](left, right Map) Map {
	result := make(Map)
	for k := range left {
		if rv, ok := right[k]; ok {
			result[k] = rv
		}
	}

	return result
}

// Difference returns a new map of key/value pairs that only appear in left.
func Difference[Map ~map[K]V, K comparable, V any](left, right Map) Map {
	result := make(Map)
	for k, v := range left {
		if _, ok := right[k]; !ok {
			result[k] = v
		}
	}

	return result
}

// YieldAll iterates over a provided iterator and returns a Map containing all elements from the iterator.
func YieldAll[Map ~map[K]V, K comparable, V any](i iter.Seq2[K, V]) Map {
	result := make(Map)
	for k, v := range i {
		result[k] = v
	}

	return result
}
