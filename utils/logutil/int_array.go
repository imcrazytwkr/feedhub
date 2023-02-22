package logutil

import "github.com/rs/zerolog"

func IntArray(source []int) *zerolog.Array {
	array := zerolog.Arr()
	for _, v := range source {
		array.Int(v)
	}

	return array
}
