package coll

import (
	"slices"
)

func Reversed[T any](items []T) []T {
	result := append([]T(nil), items...)
	slices.Reverse(result)
	return result
}

func ToSet[T comparable](items []T) map[T]bool {
	result := make(map[T]bool, len(items))
	for _, item := range items {
		result[item] = true
	}
	return result
}

func ToSetF[T any, R comparable](items []T, keyFn func(T) R) map[R]bool {
	result := make(map[R]bool, len(items))
	for _, item := range items {
		result[keyFn(item)] = true
	}
	return result
}
