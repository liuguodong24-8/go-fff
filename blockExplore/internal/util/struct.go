package util

import (
	"log"
	"reflect"
)

// GetFieldName 获取结构体中字段的名称
func GetFieldName(structName interface{}) map[string]interface{} {
	typ := reflect.TypeOf(structName)
	val := reflect.ValueOf(structName)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}

	result := make(map[string]interface{})
	for i := 0; i < typ.NumField(); i++ {

		/*fmt.Println(typ.Field(i).Name)
		fmt.Println(typ.Field(i).Type)
		fmt.Println(val.Elem().FieldByName(typ.Field(i).Name))
		*/
		result[typ.Field(i).Name] = val.Elem().FieldByName(typ.Field(i).Name)
	}
	return result
}
