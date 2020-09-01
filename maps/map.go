package maps

import (
	. "github.com/christiangelone/collgo"
	"github.com/christiangelone/collgo/elements"
)

type IMap interface {
	ICollection
	IIterable
	IMappable
	IFilterable
	Put(key , value elements.IElement)
	Get(key elements.IElement) (elements.IElement, error)
	Delete(key elements.IElement)
}
