package util

import (
	"log"
	"reflect"
)

// GetFieldName 获取结构体中字段的名称
func GetFieldName(structName interface{}) map[string]interface{} {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}

	fieldNum := t.NumField()
	result := make(map[string]interface{})
	for i := 0; i < fieldNum; i++ {
		val, _ := t.FieldByName(t.Field(i).Name)
		result[t.Field(i).Name] = val
	}
	return result
}
