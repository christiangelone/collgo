package maps


import (
	"github.com/christiangelone/collgo/elements"
)

type EachFn func(key, value elements.IElement)
type IIterable interface {
	ForEach(fn EachFn)
}

type MapFn func(key, value elements.IElement) (elements.IElement, elements.IElement)
type IMappable interface {
	Map(fn MapFn) IMap
}

type FilterFn func(key, value elements.IElement) bool
type IFilterable interface {
	Filter(fn FilterFn) IMap
}
