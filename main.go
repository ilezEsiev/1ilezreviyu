package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Примеры вызова функций
	int8MaxValues, int16MaxValues, int32MaxValues, int64MaxValues := getIntMaxValue(0, 0, 0, 0)
	fmt.Println("Max Int Values:", int8MaxValues, int16MaxValues, int32MaxValues, int64MaxValues)

	uint8MaxValues, uint16MaxValues, uint32MaxValues, uint64MaxValues := getUintMaxValue(0, 0, 0, 0)
	fmt.Println("Max Uint Values:", uint8MaxValues, uint16MaxValues, uint32MaxValues, uint64MaxValues)

	fmt.Println(EqualValues(5, 5))
	a := 5
	Increment(&a)
	fmt.Println(a)
}

func getIntMaxValue(in8 int8, in16 int16, in32 int32, in64 int64) (int8, int16, int32, int64) {
	maxInt8 := int8(1<<getBits(in8) - 1)
	maxInt16 := int16(1<<getBits(in16) - 1)
	maxInt32 := int32(1<<getBits(in32) - 1)
	maxInt64 := int64(1<<getBits(in64) - 1)

	if in8 != 0 {
		maxInt8 = in8
	}
	if in16 != 0 {
		maxInt16 = in16
	}
	if in32 != 0 {
		maxInt32 = in32
	}
	if in64 != 0 {
		maxInt64 = in64
	}

	return maxInt8, maxInt16, maxInt32, maxInt64
}

func getUintMaxValue(uin8 uint8, uin16 uint16, uin32 uint32, uin64 uint64) (uint8, uint16, uint32, uint64) {
	maxUint8 := uint8(1<<getBits(uin8) - 1)
	maxUint16 := uint16(1<<getBits(uin16) - 1)
	maxUint32 := uint32(1<<getBits(uin32) - 1)
	maxUint64 := uint64(1<<getBits(uin64) - 1)

	if uin8 != 0 {
		maxUint8 = uin8
	}
	if uin16 != 0 {
		maxUint16 = uin16
	}
	if uin32 != 0 {
		maxUint32 = uin32
	}
	if uin64 != 0 {
		maxUint64 = uin64
	}

	return maxUint8, maxUint16, maxUint32, maxUint64
}

func getBits(v interface{}) int {
	rawType := fmt.Sprintf("%T", v)
	typeBits := strings.Split(rawType, "t")[1]
	bits, _ := strconv.Atoi(typeBits)

	return bits
}

func EqualValues(a, b int) bool {
	return (a ^ b) == 0
}

func Increment(a *int) {
	*a++
}
