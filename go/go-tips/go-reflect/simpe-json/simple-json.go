// reference: 《Go 语言从入门到进阶实战》

package simpe_json

import (
	"bytes"
	"errors"
	"reflect"
	"strconv"
)

func MarshalJson(v interface{}) (string, error) {
	var b bytes.Buffer

	if err := writeAny(&b, reflect.ValueOf(v)); err == nil {
		return b.String(), nil
	} else {
		return "", err
	}
}

func writeAny(buff *bytes.Buffer, value reflect.Value) error {
	switch value.Kind() {
	case reflect.String:
		// 添加双引号
		buff.WriteString(strconv.Quote(value.String()))

	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
		// int -> string
		buff.WriteString(strconv.FormatInt(value.Int(), 10))
	case reflect.Slice:
		return writeSlice(buff, value)
	case reflect.Struct:
		return writeStruct(buff, value)
	case reflect.Ptr:
		if value.IsNil() {
			buff.WriteString("nil")
		} else {
			writeAny(buff, value.Elem())
		}
	default:
		return errors.New("unsupport kind: " + value.Kind().String())
	}
	return nil
}

func writeSlice(buff *bytes.Buffer, value reflect.Value) error {
	buff.WriteString("[")
	for i := 0; i < value.Len(); i ++ {
		// 去除每一个值
		sliceValue := value.Index(i)
		// 递归操作
		writeAny(buff, sliceValue)
		// 最后一个元素不添加 comma
		if i < value.Len() - 1 {
			buff.WriteString(",")
		}
	}
	buff.WriteString("]")
	return nil
}

func writeStruct(buff *bytes.Buffer, value reflect.Value) error {
	valueType := value.Type()
	buff.WriteString("{")
	for i := 0; i < valueType.NumField(); i ++ {
		// 取出对应的值
		fieldValue := value.Field(i)
		// 取出对应的类型
		fieldType := valueType.Field(i)

		// 以 Key-Value 写入字段
		buff.WriteString("\"")
		buff.WriteString(fieldType.Name)
		buff.WriteString("\":")
		// 字段的值
		writeAny(buff, fieldValue)
		// 最后一个元素不添加 comma
		if i < value.NumField() - 1 {
			buff.WriteString(",")
		}
	}
	buff.WriteString("}")
	return nil
}