package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// String convert any to string
func String(val interface{}) string {
	if val == nil {
		return ""
	}
	reValue := reflect.ValueOf(val)
	for reValue.Kind() == reflect.Ptr {
		reValue = reValue.Elem()
		if !reValue.IsValid() {
			return ""
		}
		val = reValue.Interface()
		if val == nil {
			return ""
		}
		reValue = reflect.ValueOf(val)
	}

	switch v := val.(type) {
	case string:
		return v
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case json.Number:
		return v.String()
	case []byte:
		return string(v)
	default:
		js, err := json.Marshal(v)
		if err != nil {
			return fmt.Sprint(v)
		}
		return string(js)
	}
}

// Int64 convert any to int64
func Int64(val interface{}) (int64, bool) {
	if val == nil {
		return 0, false
	}
	reValue := reflect.ValueOf(val)
	for reValue.Kind() == reflect.Ptr {
		reValue = reValue.Elem()
		if !reValue.IsValid() {
			return 0, false
		}
		val = reValue.Interface()
		if val == nil {
			return 0, false
		}
		reValue = reflect.ValueOf(val)
	}

	switch v := val.(type) {
	case bool:
		if v {
			return 1, true
		}
		return 0, true
	case uint8:
		return int64(v), true
	case uint16:
		return int64(v), true
	case uint32:
		return int64(v), true
	case uint64:
		return int64(v), true
	case int8:
		return int64(v), true
	case int16:
		return int64(v), true
	case int:
		return int64(v), true
	case int32:
		return int64(v), true
	case int64:
		return v, true
	case float32:
		return int64(v), true
	case float64:
		return int64(v), true
	case json.Number:
		i, err := v.Int64()
		if err != nil {
			return 0, false
		}
		return i, true
	case []byte:
		s := strings.SplitN(string(v), ".", 2)[0]
		t, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			return t, true
		}
		return 0, false
	case string:
		v = strings.SplitN(v, ".", 2)[0]
		t, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			return t, true
		}
		return 0, false
	default:
		return 0, false
	}
}

// Uint64 convert any to uint64
func Uint64(val interface{}) (uint64, bool) {
	if val == nil {
		return 0, false
	}
	reValue := reflect.ValueOf(val)
	for reValue.Kind() == reflect.Ptr {
		reValue = reValue.Elem()
		if !reValue.IsValid() {
			return 0, false
		}
		val = reValue.Interface()
		if val == nil {
			return 0, false
		}
		reValue = reflect.ValueOf(val)
	}

	switch v := val.(type) {
	case bool:
		if v {
			return 1, true
		}
		return 0, true
	case uint8:
		return uint64(v), true
	case uint16:
		return uint64(v), true
	case uint32:
		return uint64(v), true
	case uint64:
		return v, true
	case int8:
		return uint64(v), true
	case int16:
		return uint64(v), true
	case int:
		return uint64(v), true
	case int32:
		return uint64(v), true
	case int64:
		return uint64(v), true
	case float32:
		return uint64(v), true
	case float64:
		return uint64(v), true
	case json.Number:
		i, err := v.Int64()
		if err != nil {
			return 0, false
		}
		return uint64(i), true
	case []byte:
		s := strings.SplitN(string(v), ".", 2)[0]
		t, err := strconv.ParseUint(s, 10, 64)
		if err == nil {
			return t, true
		}
		return 0, false
	case string:
		v = strings.SplitN(v, ".", 2)[0]
		t, err := strconv.ParseUint(v, 10, 64)
		if err == nil {
			return t, true
		}
		return 0, false
	default:
		return 0, false
	}
}

// Atoi return strconv.Atoi without error
func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// Int convert any to int
func Int(val interface{}) (int, bool) {
	ival, succ := Int64(val)
	return int(ival), succ
}

// Uint convert any to uint
func Uint(val interface{}) (uint, bool) {
	uval, succ := Uint64(val)
	return uint(uval), succ
}

// Float64 convert any to float64
func Float64(val interface{}) (float64, bool) {
	if val == nil {
		return 0, false
	}
	reValue := reflect.ValueOf(val)
	for reValue.Kind() == reflect.Ptr {
		reValue = reValue.Elem()
		if !reValue.IsValid() {
			return 0, false
		}
		val = reValue.Interface()
		if val == nil {
			return 0, false
		}
		reValue = reflect.ValueOf(val)
	}

	switch v := val.(type) {
	case bool:
		if v {
			return 1, true
		}
		return 0, true
	case uint8:
		return float64(v), true
	case uint16:
		return float64(v), true
	case uint32:
		return float64(v), true
	case uint64:
		return float64(v), true
	case int8:
		return float64(v), true
	case int16:
		return float64(v), true
	case int32:
		return float64(v), true
	case int:
		return float64(v), true
	case int64:
		return float64(v), true
	case float32:
		return float64(v), true
	case float64:
		return v, true
	case []byte:
		t, err := strconv.ParseFloat(string(v), 64)
		if err == nil {
			return t, true
		}
		return 0, false
	case json.Number:
		i, err := v.Float64()
		if err != nil {
			return 0, false
		}
		return i, true
	case string:
		if len(v) > 15 {
			return 0, false
		}
		t, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return t, true
		}
		return 0, false
	default:
		return 0, false
	}
}

// Bool convert any to bool
func Bool(val interface{}) (bool, bool) {
	ival, succ := Int64(val)
	return ival != 0, succ
}

// IsNil reports whether val is nil
func IsNil(val interface{}) bool {
	if val == nil {
		return true
	}
	reValue := reflect.ValueOf(val)
	for reValue.Kind() == reflect.Ptr {
		reValue = reValue.Elem()
		if !reValue.IsValid() || reValue.IsNil() {
			return true
		}
		reValue = reflect.ValueOf(reValue.Interface())
	}
	return false
}

// Time convert any to time.Time
func Time(val interface{}) (time.Time, bool) {
	if val == nil {
		return time.Time{}, false
	}
	reValue := reflect.ValueOf(val)
	for reValue.Kind() == reflect.Ptr {
		reValue = reValue.Elem()
		if !reValue.IsValid() {
			return time.Time{}, false
		}
		val = reValue.Interface()
		if val == nil {
			return time.Time{}, false
		}
		reValue = reflect.ValueOf(val)
	}

	switch v := val.(type) {
	case time.Time:
		return v, true
	case string:
		t, _, ok := TimeAndFormat(v)
		return t, ok
	default:
		return time.Time{}, false
	}
}

// TimeAndFormat convert string to time.Time and return format
func TimeAndFormat(v string) (time.Time, string, bool) {
	tlen := len(v)
	var t time.Time
	var format string
	var err error
	switch tlen {
	case 6:
		t, err = time.ParseInLocation("060102", v, time.Local)
		format = "060102"
	case 8:
		t, err = time.ParseInLocation("20060102", v, time.Local)
		format = "20060102"
	case 10:
		t, err = time.ParseInLocation("2006-01-02", v, time.Local)
		format = "2006-01-02"
	case 19:
		t, err = time.ParseInLocation("2006-01-02 15:04:05", v, time.Local)
		format = "2006-01-02 15:04:05"
		if err != nil {
			t, err = time.Parse(time.RFC822, v)
			format = time.RFC822
		}
	case 20:
		t, err = time.ParseInLocation("2006-01-02 15:04:05Z", v, time.Local)
		if err != nil {
			t, err = time.Parse("2006-01-02T15:04:05Z", v)
			format = "2006-01-02T15:04:05Z"
		}
		format = "2006-01-02 15:04:05Z"
	case len(time.ANSIC):
		t, err = time.Parse(time.ANSIC, v)
		format = time.ANSIC
	case len(time.UnixDate):
		t, err = time.Parse(time.UnixDate, v)
		format = time.UnixDate
	case len(time.RubyDate):
		t, err = time.Parse(time.RFC850, v)
		format = time.RFC850
		if err != nil {
			t, err = time.Parse(time.RubyDate, v)
			format = time.RubyDate
		}
	case len(time.RFC822Z):
		t, err = time.Parse(time.RFC822Z, v)
		format = time.RFC822Z
	case len(time.RFC1123):
		t, err = time.Parse(time.RFC1123, v)
		format = time.RFC1123
	case len(time.RFC1123Z):
		t, err = time.Parse(time.RFC1123Z, v)
		format = time.RFC1123Z
	case len(time.RFC3339):
		t, err = time.Parse(time.RFC3339, v)
		format = time.RFC3339
	default:
		t, err = time.Parse(time.RFC3339Nano, v)
		format = time.RFC3339Nano
	}
	if err != nil {
		return t, "", false
	}
	return t, format, true
}

// TimePtr convert any to *time.Time
func TimePtr(val interface{}) *time.Time {
	t, ok := Time(val)
	if ok {
		return &t
	}

	return nil
}
