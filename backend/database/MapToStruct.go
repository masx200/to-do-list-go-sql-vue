package database

import "encoding/json"

func MapToStruct[T any](m map[string]any) *T {

	str, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}
	var obj T
	err = json.Unmarshal(str, &obj)
	if err != nil {
		panic(err)
	}
	return &obj
}
