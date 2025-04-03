package slices

// FindByValue find slice by value, return value and ok
func FindByValue[T any, R comparable](s []T, val R, f func(v T) R) (r T, ok bool) {
	for _, v := range s {
		if f(v) == val {
			return v, true
		}
	}
	return
}

// Map map slice
func Map[T any, R any](s []T, f func(v T) R) []R {
	r := make([]R, 0, len(s))
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}
