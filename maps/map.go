package maps

func Keys[K comparable, V any](m map[K]V) []K {
  keys := make([]K, 0)
  for k := range m {
    keys = append(keys, k)
  }

  return keys
}

func Values[K comparable, V any](m map[K]V) []V {
  values := make([]V, 0)
  for _, v := range m {
    values = append(values, v)
  }

  return values
}
