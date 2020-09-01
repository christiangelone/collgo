package hashmap

import (
	"github.com/christiangelone/collgo/elements"
	"github.com/christiangelone/collgo/maps"
)

func New() *HashMap {
	return &HashMap{
		size: 0,
		mapping: map[string] []elements.IElement{},
	}
}

type HashMap struct {
	size     uint64
	mapping  map[string] []elements.IElement
}

func (h *HashMap) ForEach(eachFn maps.EachFn) {
	for k := range h.mapping{
		val := h.mapping[k]
		eachFn(val[0], val[1])
	}
}

func (h *HashMap) Map(mapFn maps.MapFn) maps.IMap {
	aMap := New()
	h.ForEach(func(key, value elements.IElement) {
		mapKey, mapVal := mapFn(key, value)
		aMap.Put(mapKey, mapVal)
	})
	return aMap
}

func (h *HashMap) Filter(filterFn maps.FilterFn) maps.IMap {
	aMap := New()
	h.ForEach(func(key, value elements.IElement) {
		if filterFn(key, value) {
			aMap.Put(key, value)
		}
	})
	return aMap
}

func (h *HashMap) Size() uint64 {
	return h.size
}

func (h *HashMap) IsEmpty() bool {
	return h.size == 0
}

func (h *HashMap) Put(key, value elements.IElement) {
	h.mapping[key.ToString().Value().(string)] = []elements.IElement{ key, value }
	h.size++
}

func (h *HashMap) Get(key elements.IElement) (elements.IElement, error) {
	keyStr := key.ToString().Value().(string)
	if e, ok := h.mapping[keyStr]; ok {
		return e[1], nil
	}
	return nil, &maps.NotFoundKeyError{Key: keyStr}
}

func (h *HashMap) Delete(key elements.IElement) {
	if len(h.mapping) > 0 {
		delete(h.mapping, key.ToString().Value().(string))
		h.size--
	}
}
