package utils

import (
	"fmt"
	"regexp"
	"strings"
)



func CheckField(field string) (err error) {
	if len(field) <= 0 {
		return fmt.Errorf("field is null")
	}
	//验证是否为field
	pattern := `^[a-zA-Z0-9_,\.]*$`
	matched, err := regexp.MatchString(pattern, field)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("所传字段[%s]存在注入风险!", field)
	}
	return nil
}

//IsInStringArr 是否在数组中
func IsInStringArr(arr []string, id string) bool {
	v := strings.ToLower(strings.TrimSpace(id))
	for _, rsID := range arr {
		if len(rsID) == 0 {
			continue
		}
		nv := strings.ToLower(strings.TrimSpace(rsID))
		if v == nv {
			return true
		}
	}
	return false
}