package validate

// SliceFindByValue find slice by value, return value and ok
func SliceFindByValue[T any, R comparable](s []T, val R, f func(v T) R) (r T, ok bool) {
	for _, v := range s {
		if f(v) == val {
			return v, true
		}
	}
	return
}
