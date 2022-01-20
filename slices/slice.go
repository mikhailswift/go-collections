package slices

// Filter returns a new slice consisting of all elements that pass the predicate function
func Filter[TSlice ~[]T, T any](slice TSlice, predicate func(T) bool) TSlice {
  selected := make(TSlice, 0)
  for _, t := range slice {
    if predicate(t) {
      selected = append(selected, t)
    }
  }

  return selected
}

// FirstMatch returns the first element that passes the predicate function.
// If no element is found the zero value of the type will be returned.
func FirstMatch[TSlice ~[]T, T any](slice TSlice, predicate func(T) bool) T {
  for _, t := range slice {
    if predicate(t) {
      return t
    }
  }

  var zero T
  return zero  
}

// LastMatch returns the last element that passes the predicate function.
// If no element is found the zero value of the type will be returned.
func LastMatch[TSlice ~[]T, T any](slice TSlice, predicate func(T) bool) T {
  for i := len(slice)-1; i >= 0; i-- {
    if predicate(slice[i]) {
      return slice[i]
    }
  }

  var zero T
  return zero 
}

// AnyMatch returns true if any element passes the predicate function
func AnyMatch[TSlice ~[]T, T any](slice TSlice, predicate func(T) bool) bool {
  for _, t := range slice {
    if predicate(t) {
      return true
    }
  }

  return false
}

// AllMatch returns true if all elements pass the predicate function
func AllMatch[TSlice ~[]T, T any](slice TSlice, predicate func(T) bool) bool {
  for _, t := range slice {
    if !predicate(t) {
      return false
    }
  }
  
  return true
}

// Map returns a new slice where each element is the result of fn for the corresponding element in the original slice
func Map[TSlice ~[]T, T any, U any](slice []T, fn func(T) U) []U {
  result := make([]U, len(slice))
  for i, t := range slice {
    result[i] = fn(t)
  }

  return result
}

// Contains returns true if find appears in slice
func Contains[TSlice ~[]T, T comparable](slice TSlice, find T) bool {
  for _, t := range slice {
    if t == find {
      return true
    }
  }

  return false
}

// IndexOf returns the index of find if it appears in slice. If find is not in slice, -1 will be returned.
func IndexOf[TSlice ~[]T, T comparable](slice TSlice, find T) int {
  for i, t := range slice {
    if t == find {
      return i
    }
  }

  return -1
}


// GroupBy returns a map that is keyed by keySelector and contains a slice of elements returned by valSelector
func GroupBy[TSlice ~[]T, T any, K comparable, V any](slice TSlice, keySelector func(T) K, valSelector func(T) V) map[K][]V {
  grouping := make(map[K][]V)
  for _, t := range slice {
    key := keySelector(t)
    grouping[key] = append(grouping[key], valSelector(t))
  }

  return grouping
}

// ToSet returns a map keyed by keySelector and contains a value of an empty struct
func ToSet[TSlice ~[]T, T any, K comparable](slice TSlice, keySelector func(T) K) map[K]struct{} {
  set := make(map[K]struct{})
  for _, t := range slice {
    set[keySelector(t)] = struct{}{}
  }

  return set
}

// ToMap return a map that is keyed keySelector and has the value of valSelector for each element in slice.
// If multiple elements return the same key the element that appears later in slice will be chosen.
func ToMap[TSlice ~[]T, T any, K comparable, V any](slice TSlice, keySelector func(T) K, valSelector func(T) V) map[K]V {
  m := make(map[K]V)
  for _, t := range slice {
    m[keySelector(t)] = valSelector(t)
  }

  return m
}
