package utils

import "strconv"

func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return int(num)
}

func StringToInt8(str string) int8 {
	num64, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0
	}
	return int8(num64)
}

func StringToInt16(str string) int16 {
	num64, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0
	}
	return int16(num64)
}

func StringToInt32(str string) int32 {
	num64, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0
	}
	return int32(num64)
}

func StringToInt64(str string) int64 {
	num64, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0
	}
	return num64
}
