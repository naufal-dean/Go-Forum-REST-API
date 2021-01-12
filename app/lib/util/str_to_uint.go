package util

import "strconv"

func StrToUint(s string) (uint, error) {
	u, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(u), nil
}
