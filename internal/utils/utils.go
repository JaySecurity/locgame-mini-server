package utils

import "locgame-mini-server/pkg/dto/errors"

func Includes[T comparable](list []T, val T) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func ReplaceValue[T comparable](list *[]T, current T, updated T) error {
	foundIdx := -1
	for i, item := range *list {
		if item == current {
			foundIdx = i
			break
		}
	}
	if foundIdx == -1 {
		return errors.ErrMatchNotFound
	}
	if foundIdx < len(*list) {
		(*list)[foundIdx] = updated
	}
	return nil
}

func RemoveValue[T comparable](list *[]T, val T) error {
	foundIdx := -1
	for i, item := range *list {
		if item == val {
			foundIdx = i
			break
		}
	}
	if foundIdx == -1 {
		return errors.ErrMatchNotFound
	}
	if foundIdx == 0 && len(*list) == 1 {
		*list = []T{}
	} else if foundIdx < len(*list) {
		*list = append((*list)[:foundIdx], (*list)[foundIdx+1:]...)
	}
	return nil
}
