package elements

import "github.com/christiangelone/collgo/elements/primitives"

type IElement interface {
	Value() primitives.Any
	primitives.IConvertible
}








