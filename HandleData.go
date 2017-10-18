package utils

import (
	"bytes"
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func StringInArray(key string, array []string) bool {
	if key == "" || len(array) == 0 {
		return false
	}
	for _, v := range array {
		if strings.ToLower(key) == strings.ToLower(v) {
			return true
		}
	}
	return false
}

func StructToMap(data interface{}, lowerKey, bExcept bool, fields ...string) (m map[string]interface{}, err error) {
	m = make(map[string]interface{}, 0)
	if data == nil {
		err = errors.New("struct to map error data is nil")
		return
	}
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		err = errors.New("data must be struct")
		return
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if bExcept {
			if StringInArray(t.Field(i).Name, fields) {
				continue
			}
		} else {
			if !StringInArray(t.Field(i).Name, fields) {
				continue
			}
		}
		if lowerKey {
			m[strings.ToLower(t.Field(i).Name)] = f.Interface()
		} else {
			m[t.Field(i).Name] = f.Interface()
		}
	}
	return
}

func IpToUint32(ipStr string) uint32 {
	parts := strings.Split(ipStr, ".")
	if len(parts) != 4 {
		return 0
	}
	var (
		dst uint64
	)
	for k, v := range parts {
		tmp, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return 0
		}
		dst |= tmp << uint((3-k)*8)
	}
	return uint32(dst)
}

func Uint32ToIp(ipUint uint32) string {
	ipSegs := make([]string, 4)
	length := len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < length; i++ {
		tempUint := ipUint & 0xFF
		ipSegs[length-i-1] = strconv.FormatUint(uint64(tempUint), 10)
		ipUint = ipUint >> 8
	}
	for i := 0; i < length; i++ {
		buffer.WriteString(ipSegs[i])
		if i < length-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}
