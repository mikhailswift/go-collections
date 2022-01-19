package slices

// Where returns a new slice that passes the predicate function
func Where[T any](arr []T, predicate func(T) bool) []T {
  selected := make([]T, 0)
  for _, t := range arr {
    if predicate(t) {
      selected = append(selected, t)
    }
  }

  return selected
}

// First returns the first element that passes the predicate function.
// If no element is found the zero value of the type will be returned.
func First[T any](arr []T, predicate func(T) bool) T {
  for _, t := range arr {
    if predicate(t) {
      return t
    }
  }

  var zero T
  return zero  
}

// Last returns the last element that passes the predicate function.
// If no element is found the zero value of the type will be returned.
func Last[T any](arr []T, predicate func(T) bool) T {
  for i := len(arr)-1; i >= 0; i-- {
    if predicate(arr[i]) {
      return arr[i]
    }
  }

  var zero T
  return zero 
}

// Any returns true if any element passes the predicate function
func Any[T any](arr []T, predicate func(T) bool) bool {
  for _, t := range arr {
    if predicate(t) {
      return true
    }
  }

  return false
}

// All returns true if all elements pass the predicate function
func All[T any](arr []T, predicate func(T) bool) bool {
  for _, t := range arr {
    if !predicate(t) {
      return false
    }
  }
  
  return true
}

// Map returns a new slice where each element is the result of fn for the corresponding element in the original slice
func Map[T any, U any](arr []T, fn func(T) U) []U {
  result := make([]U, len(arr))
  for i, t := range arr {
    result[i] = fn(t)
  }

  return result
}

// Contains returns true if find appears in arr
func Contains[T comparable](arr []T, find T) bool {
  for _, t := range arr {
    if t == find {
      return true
    }
  }

  return false
}

// IndexOf returns the index of find if it appears in arr. If find is not in arr, -1 will be returned.
func IndexOf[T comparable](arr []T, find T) int {
  for i, t := range arr {
    if t == find {
      return i
    }
  }

  return -1
}


// GroupBy returns a map that is keyed by keySelector and contains an slice of elements returned by valSelector
func GroupBy[T any, K comparable, V any](arr []T, keySelector func(T) K, valSelector func(T) V) map[K][]V {
  grouping := make(map[K][]V)
  for _, t := range arr {
    key := keySelector(t)
    grouping[key] = append(grouping[key], valSelector(t))
  }

  return grouping
}

// ToSet returns a map keyed by keySelector and contains a value of an empty struct
func ToSet[T any, K comparable](arr []T, keySelector func(T) K) map[K]struct{} {
  set := make(map[K]struct{})
  for _, t := range arr {
    set[keySelector(t)] = struct{}{}
  }

  return set
}

// ToMap return a map that is keyed keySelector and has the value of valSelector for each element in arr.
// If multiple elements return the same key the element that appears later in arr will be chosen.
func ToMap[T any, K comparable, V any](arr []T, keySelector func(T) K, valSelector func(T) V) map[K]V {
  m := make(map[K]V)
  for _, t := range arr {
    m[keySelector(t)] = valSelector(t)
  }

  return m
}