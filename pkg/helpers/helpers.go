package helpers

import (
	"github.com/google/uuid"
	"reflect"
)

func GenerateUUID() string {
	guid := uuid.New()
	return guid.String()
}

func PointerDeepEqual(x interface{}, y interface{}) bool {
	rawType := reflect.TypeOf(y)

	if rawType.Kind() == reflect.Ptr {
		rawType = rawType.Elem()
		return reflect.DeepEqual(x, reflect.New(rawType).Interface())
	}
	return false
}
