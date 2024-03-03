package helper

import "strconv"

func ConvertToInt(input string) (res int64, err error) {
	cvDate, er := strconv.Atoi(input)
	cv64 := int64(cvDate)
	if er != nil {
		return res, er
	}
	return cv64, err
}

func ConvertToString(input int64) (res string, err error) {
	cv64 := strconv.Itoa(int(input))
	return cv64, err
}
