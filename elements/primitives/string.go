package primitives

import (
	"strconv"
)

type String struct {
	value string
}

func NewString(aString string) *String {
	return &String{
		value: aString,
	}
}

func (s *String) Value() Any {
	return s.value
}

func (s *String) ToString() *String {
	return s
}

func (s *String) ToInt() *Int {
	anInt, err := strconv.Atoi(s.value)
	if err != nil { panic(err) }
	return NewInt(anInt)
}

func (s *String) ToDouble() *Double {
	aDouble, err := strconv.ParseFloat(s.value, 64)
	if err != nil { panic(err) }
	return NewDouble(aDouble)
}

func (s *String) ToFloat() *Float {
	aFloat, err := strconv.ParseFloat(s.value, 32)
	if err != nil { panic(err) }
	return NewFloat(float32(aFloat))
}

func (s *String) ToBool() *Bool {
	aBool, err := strconv.ParseBool(s.value)
	if err != nil { panic(err) }
	return NewBool(aBool)
}
