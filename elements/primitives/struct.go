package primitives

import (
	"fmt"
	"strings"
)

func NewStruct(aStruct Any) *Struct {
	return &Struct{
		value: aStruct,
	}
}

type Struct struct {
	value Any
}

func (s *Struct) ToInt() *Int {
	if i, ok := s.value.(int); ok {
		return NewInt(i)
	}
	return NewInt(0)
}

func (s *Struct) ToDouble() *Double {
	if d, ok := s.value.(float64); ok {
		return NewDouble(d)
	}
	return NewDouble(0.0)
}

func (s *Struct) ToFloat() *Float {
	if f, ok := s.value.(float32); ok {
		return NewFloat(f)
	}
	return NewFloat(0.0)
}

func (s *Struct) ToBool() *Bool {
	if b, ok := s.value.(bool); ok {
		return NewBool(b)
	}
	return NewBool(false)
}

func (s *Struct) ToString() *String {
	theType := fmt.Sprintf("%T", s.value)
	theValue := strings.ReplaceAll(fmt.Sprintf("%+v", s.value), ":", ": ")
	theValue = strings.ReplaceAll(theValue, "}", " }")
	theValue = strings.ReplaceAll(theValue, "{", "{ ")
	return NewString(fmt.Sprintf("(%s)(%s)", theType, theValue))
}

func (s *Struct) Value() Any {
	return s.value
}
