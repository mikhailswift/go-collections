package array

func Where[T any](arr []T, predicate func(T) bool) []T {
  selected := make([]T, 0)
  for _, t := range arr {
    if predicate(t) {
      selected = append(selected, t)
    }
  }

  return selected
}

func First[T any](arr []T, predicate func(T) bool) T {
  for _, t := range arr {
    if predicate(t) {
      return t
    }
  }

  var zero T
  return zero  
}

func Last[T any](arr []T, predicate func(T) bool) T {
  for i := len(arr)-1; i >= 0; i-- {
    if predicate(arr[i]) {
      return arr[i]
    }
  }

  var zero T
  return zero 
}

func Any[T any](arr []T, predicate func(T) bool) bool {
  for _, t := range arr {
    if predicate(t) {
      return true
    }
  }

  return false
}

func All[T any](arr []T, predicate func(T) bool) bool {
  for _, t := range arr {
    if !predicate(t) {
      return false
    }
  }
  
  return true
}

func Map[T any, U any](arr []T, fn func(T) U) []U {
  result := make([]U, 0)
  for _, t := range arr {
    result = append(result, fn(t))
  }

  return result
}

func Contains[T comparable](arr []T, find T) bool {
  for _, t := range arr {
    if t == find {
      return true
    }
  }

  return false
}

func GroupBy[T any, K comparable, V any](arr []T, keySelector func(T) K, valSelector func(T) V) map[K][]V {
  grouping := make(map[K][]V)
  for _, t := range arr {
    grouping[keySelector(t)] = append(grouping[keySelector(t)], valSelector(t))
  }

  return grouping
}

func ToSet[T any, K comparable](arr []T, keySelector func(T) K) map[K]struct{} {
  set := make(map[K]struct{})
  for _, t := range arr {
    set[keySelector(t)] = struct{}{}
  }

  return set
}

func ToMap[T any, K comparable, V any](arr []T, keySelector func(T) K, valSelector func(T) V) map[K]V {
  m := make(map[K]V)
  for _, t := range arr {
    m[keySelector(t)] = valSelector(t)
  }

  return m
}
