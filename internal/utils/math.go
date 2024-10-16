package utils

// RoundToInt32 rounds float32 to int.
func RoundToInt32(a float32) int32 {
	if a < 0 {
		return int32(a - 0.5)
	}
	return int32(a + 0.5)
}

// RoundUpToInt32 rounds a float32 up to a larger int value.
func RoundUpToInt32(a float32) int32 {
	if a < 0 {
		return int32(a - 1)
	}
	return int32(a + 1)
}
