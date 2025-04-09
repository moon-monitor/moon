package kv

func NewStringMap(ms ...map[string]string) StringMap {
	return New[string, string](ms...)
}

type StringMap = Map[string, string]
