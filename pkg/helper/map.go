package helper

func GetMapFromSliceByField[Obj any, Field comparable](slice []Obj, getField func(Obj) Field) map[Field]Obj {
	result := make(map[Field]Obj, len(slice))
	for _, item := range slice {
		result[getField(item)] = item
	}
	return result
}
