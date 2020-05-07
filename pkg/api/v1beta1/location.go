package v1beta1

import "sync"

//Location is a city, country, county or division
//within a city
type Location string

type SafeLocationMap struct {
	Locations map[Location]string
	mux       sync.Mutex
}

func NewSafeLocationMap() *SafeLocationMap {
	l := map[Location]string{}
	return &SafeLocationMap{
		Locations: l,
	}
}

func (lm *SafeLocationMap) Insert(l Location) {
	lm.mux.Lock()
	lm.Locations[l] = ""
	lm.mux.Unlock()
}
