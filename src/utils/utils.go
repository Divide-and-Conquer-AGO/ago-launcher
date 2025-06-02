package utils

import (
	"errors"
	"math/rand"
)

func RandomElement[T any](slice []T) (T, error) {
	var zero T
	if len(slice) == 0 {
		return zero, errors.New("no elements available")
	}
	return slice[rand.Intn(len(slice))], nil
}
