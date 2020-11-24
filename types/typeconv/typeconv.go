package typeconv

import (
	"fmt"
	"strconv"
)

/**
 * Copyright (C) 2007-2020 56X.NET,All rights reserved.
 *
 * name : typeconv.go
 * author : jarrysix (jarrysix#gmail.com)
 * date : 2020-11-21 20:19
 * description :
 * history :
 */

// 将类型转为string
func MustInt(d interface{}) int {
	switch d.(type) {
	case string:
		i, err := strconv.Atoi(d.(string))
		if err != nil {
			panic("parse string to int fail:" + err.Error())
		}
		return i
	case float32:
		return int(d.(float32))
	case float64:
		return int(d.(float64))
	case int:
		return d.(int)
	case int8:
		return d.(int)
	case int16:
		return d.(int)
	case int32:
		return d.(int)
	case int64:
		return d.(int)
	}
	panic("not support type:" + fmt.Sprintf("%+v", d))
}


// Get string of interface, if can't converted,
// will return false
func String(d interface{}) (string, bool) {
	switch d.(type) {
	case string:
		return d.(string), true
	case []byte:
		return string(d.([]byte)), true
	case float32:
		return strconv.FormatFloat(float64(d.(float32)), 'g', 2, 32), true
	case float64:
		return strconv.FormatFloat(d.(float64), 'g', 2, 64), true
	case int:
		return strconv.FormatInt(int64(d.(int)), 10), true
	case int8:
		return strconv.FormatInt(int64(d.(int8)), 10), true
	case int16:
		return strconv.FormatInt(int64(d.(int16)), 10), true
	case int32:
		return strconv.FormatInt(int64(d.(int32)), 10), true
	case int64:
		return strconv.FormatInt(d.(int64), 10), true
	case uint:
		return strconv.FormatUint(uint64(d.(uint)), 10), true
	case uint8:
		return strconv.FormatUint(uint64(d.(uint8)), 10), true
	case uint16:
		return strconv.FormatUint(uint64(d.(uint16)), 10), true
	case uint32:
		return strconv.FormatUint(uint64(d.(uint32)), 10), true
	case uint64:
		return strconv.FormatUint(d.(uint64), 10), true
	case bool:
		return strconv.FormatBool(d.(bool)), true
	}
	return "", false
}