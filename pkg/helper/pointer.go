package helper

func GetValueFromPointer[T any](val *T, defaultVal T) T {
	if val != nil {
		return *val
	}
	return defaultVal
}

func GetPointer[T any](val T) *T {
	return &val
}

func CopyPointer[T any](valPtr *T) *T {
	if valPtr == nil {
		return nil
	}
	val := *valPtr
	return &val
}
