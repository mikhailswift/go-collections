package maps

// Keys returns an array of all keys in m
func Keys[K comparable, V any](m map[K]V) []K {
  keys := make([]K, 0)
  for k := range m {
    keys = append(keys, k)
  }

  return keys
}


// Values returns an array of all values in m
func Values[K comparable, V any](m map[K]V) []V {
  values := make([]V, 0)
  for _, v := range m {
    values = append(values, v)
  }

  return values
}

// Union returns a new map of all key/value pairs in left and right. If a key exists
// in both left and right the value in right will appear in the resultant map.
func Union[K comparable, V any](left, right map[K]V) map[K]V {
  result := make(map[K]V)
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
func UnionInPlace[K comparable, V any](left, right map[K]V) map[K]V {
  for k, v := range right {
    left[k] = v
  }

  return left
}


// Intersect returns a new map of key/value pairs where the key exists in both left and right.
// The value from right will be used in the return map.
func Intersect[K comparable, V any](left, right map[K]V) map[K]V {
  result := make(map[K]V)
  for k := range left {
    if rv, ok := right[k]; ok {
      result[k] = rv
    }
  }

  return result
}

// Difference returns a new map of key/value pairs that only appear in left.
func Difference[K comparable, V any](left, right map[K]V) map[K]V {
  result := make(map[K]V)
  for k, v := range left {
    if _, ok := right[k]; !ok {
      result[k] = v
    }
  }

  return result
}
