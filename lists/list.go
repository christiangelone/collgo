package lists

import (
	. "github.com/christiangelone/collgo"
	"github.com/christiangelone/collgo/elements"
)

type IList interface {
	ICollection
	IIterable
	IMappable
	IFilterable
	IReducible
	IInclusive
	ISortable
	IZippable
	First() (elements.IElement, error)
	Last() (elements.IElement, error)
	Add(element elements.IElement)
}

type IStack interface {
	ICollection
	IIterable
	IMappable
	IFilterable
	IReducible
	IInclusive
	Top() (elements.IElement, error)
	Push(element elements.IElement)
	Pop() (elements.IElement, error)
}

type IQueue interface {
	ICollection
	IIterable
	IMappable
	IFilterable
	IReducible
	IInclusive
	Head() (elements.IElement, error)
	Enqueue(element elements.IElement)
	Dequeue() (elements.IElement, error)
}
