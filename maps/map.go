package maps

// Keys returns a slice of all keys in m
func Keys[Map ~map[K]V, K comparable, V any](m Map) []K {
  keys := make([]K, len(m))
  i := 0
  for k := range m {
    keys[i] = k
    i++
  }

  return keys
}


// Values returns a slice of all values in m
func Values[Map ~map[K]V, K comparable, V any](m Map) []V {
  values := make([]V, len(m))
  i := 0
  for _, v := range m {
    values[i] = v
    i++
  }

  return values
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
  result := make(map[K]V)
  for k, v := range left {
    if _, ok := right[k]; !ok {
      result[k] = v
    }
  }

  return result
}
