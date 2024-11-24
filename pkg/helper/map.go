package helper

func GetMapFromSlice[T comparable](slice []T) map[T]struct{} {
	result := make(map[T]struct{}, len(slice))
	for _, item := range slice {
		result[item] = struct{}{}
	}
	return result
}
