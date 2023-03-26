package anyUtil

import "reflect"

// AnyToBool 将给定的值转换为bool
func AnyToBool(i interface{}) bool {
	if i == nil {
		return false
	}

	// 检查解引用后的值是否为 nil
	if reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil() {
		return false
	}

	v := reflect.ValueOf(i)
	// 处理指针类型
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		if reflect.ValueOf(i).IsNil() {
			return false
		}
		v = reflect.ValueOf(i).Elem()
	}

	switch v.Kind() {
	case reflect.Bool:
		return v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return v.Float() != 0
	case reflect.Complex64, reflect.Complex128:
		return v.Complex() != 0
	case reflect.String:
		val := v.String()
		if val == "true" {
			return true
		} else if val == "false" {
			return false
		}
		return v.String() != ""
	default:
		return false
	}
}