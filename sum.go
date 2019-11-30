package gomod_sum

import "reflect"

func Sum(numbers ...interface{}) int64 {
	res := int64(0)
	for _, number := range numbers {
		res += convertToInt64(number)
	}
	return res
}

func convertToInt64(data interface{}) int64 {
	if data == nil {
		return 0
	}
	int64Type := reflect.TypeOf(int64(0))
	v := reflect.ValueOf(data)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(int64Type) {
		return 0
	}
	res := v.Convert(int64Type)
	return res.Int()
}
