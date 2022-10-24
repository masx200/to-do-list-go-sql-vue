package database

import "encoding/json"

func StructToMap[T any](obj *T) map[string]any {

	str, err := json.Marshal(obj)

	if err != nil {
		panic(err)
	}
	var m = map[string]any{}
	err = json.Unmarshal(str, &m)
	if err != nil {
		panic(err)
	}
	return m
}
