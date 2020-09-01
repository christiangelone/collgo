package lists

import (
	"github.com/christiangelone/collgo/elements"
)

type IIterator interface {
	First() (element elements.IElement, end bool)
	Next() (element elements.IElement, end bool)
	WithOffset(offset uint64) IIterator
}

type EachFn func(elements.IElement)
type IIterable interface {
	ForEach(fn EachFn)
	GetRightIterator() IIterator
	GetLeftIterator() IIterator
}

type MapFn func(element elements.IElement) elements.IElement
type IMappable interface {
	Map(fn MapFn) IList
}

type ReduceFn func(elementA elements.IElement, elementB elements.IElement) elements.IElement
type IReducible interface {
	LeftReduce(fn ReduceFn) elements.IElement
	LeftReduceWith(fn ReduceFn, initialElement elements.IElement) elements.IElement
	RightReduce(fn ReduceFn) elements.IElement
	RightReduceWith(fn ReduceFn, initialElement elements.IElement) elements.IElement
}

type FilterFn func(element elements.IElement) bool
type IFilterable interface {
	Filter(fn FilterFn) IList
}

type IZippable interface {
	Zip(list IList) IList
	AdjustedZip(list IList, fillValue interface{}) IList
}

type IInclusive interface {
	Has(element elements.IElement) bool
}

type CompareFn func(a,b elements.IElement) bool
type ISortable interface {
	Sort(fn CompareFn) IList
	IsSorted(fn CompareFn) bool
}
