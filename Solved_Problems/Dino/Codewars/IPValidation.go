package kata

import (
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {

	test := strings.Split(ip, ".")
	if len(test) != 4 {
		return false
	}
	for _, i := range test {
		if _, err := strconv.ParseInt(i, 10, 64); err != nil {
			return false
		}
		if i[0] == '0' && len(i) > 1 {
			return false
		}
		integer, _ := strconv.ParseInt(i, 10, 64)
		if integer > 255 || integer < 0 {
			return false
		}
	}
	return true
}
