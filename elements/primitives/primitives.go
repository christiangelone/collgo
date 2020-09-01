package primitives

type Any = interface{}

type IConvertible interface {
	ToString() *String
	ToInt() *Int
	ToDouble() *Double
	ToFloat() *Float
	ToBool() *Bool
}
