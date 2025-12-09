package utils

import "sync"

type SettledResult[T any] struct {
	Value T
	Error error
}

func PromiseAllSettled[T any](funcs []func() (T, error)) []SettledResult[T] {
	var wg sync.WaitGroup
	results := make([]SettledResult[T], len(funcs))

	for i, fn := range funcs {
		wg.Add(1)
		go func(idx int, f func() (T, error)) {
			defer wg.Done()
			res, err := f()
			results[idx] = SettledResult[T]{
				Value: res,
				Error: err,
			}
		}(i, fn)
	}

	wg.Wait()
	return results
}

func ForEach[T any](items []T, fn func(T)) {
	for _, item := range items {
		fn(item)
	}
}

func Map[T any, R any](items []T, fn func(T) R) []R {
	result := make([]R, len(items))
	for i, item := range items {
		result[i] = fn(item)
	}
	return result
}

func Filter[T any](items []T, fn func(T) bool) []T {
	var result []T
	for _, item := range items {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

func Find[T any](items []T, fn func(T) bool) *T {
	for _, item := range items {
		if fn(item) {
			return &item
		}
	}
	return nil
}
