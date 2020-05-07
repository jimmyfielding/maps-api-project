package cache

import (
	"sync"

	"github.com/jimmyfielding/maps-api-project/pkg/api/v1beta1"
)

type Cache struct {
	mem map[string][]v1beta1.Location
	mux sync.Mutex
}

func NewCache() *Cache {
	m := map[string][]v1beta1.Location{}
	return &Cache{
		mem: m,
	}
}

func (c *Cache) Insert(latlng string, locations []v1beta1.Location) {
	c.mux.Lock()
	c.mem[latlng] = locations
	c.mux.Unlock()
}

func (c *Cache) Check(latlng string) ([]v1beta1.Location, bool) {
	var locs []v1beta1.Location
	var found bool
	if l, ok := c.mem[latlng]; ok {
		locs = l
		found = true
	}

	return locs, found
}
