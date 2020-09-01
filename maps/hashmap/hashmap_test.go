package hashmap_test

import (
	"github.com/christiangelone/collgo/elements/primitives"
	"github.com/christiangelone/collgo/maps"
	"github.com/christiangelone/collgo/maps/hashmap"
	"github.com/realistschuckle/testify/assert"
	"testing"
)

func TestPuttingInHashMap(t *testing.T) {
	aMap := hashmap.New()
	aMap.Put(primitives.NewString("a"), primitives.NewInt(1))
	aMap.Put(primitives.NewString("a"),  primitives.NewInt(2))

	var v, _ = aMap.Get(primitives.NewString("a"))
	assert.Equal(t, primitives.NewInt(2), v)
}

func TestGettingInHashMap(t *testing.T) {
	aMap := hashmap.New()
	aMap.Put(primitives.NewString("a"), primitives.NewInt(1))

	var v1, _ = aMap.Get(primitives.NewString("a"))
	var v2, _ = aMap.Get(primitives.NewString("a"))
	assert.Equal(t, v1, v2)
}

func TestDeletingInHashMap(t *testing.T) {
	aMap := hashmap.New()
	aMap.Put(primitives.NewString("a"), primitives.NewInt(1))
	aMap.Delete(primitives.NewString("a"))

	var _, err = aMap.Get(primitives.NewString("a"))
	assert.NotNil(t, err)
	assert.IsType(t, &maps.NotFoundKeyError{}, err)
}