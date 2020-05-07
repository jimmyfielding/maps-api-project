package title

import (
	"fmt"
	"sync"
	"time"

	"github.com/jimmyfielding/maps-api-project/internal/googlemaps"
	"github.com/jimmyfielding/maps-api-project/pkg/api/v1beta1"
	"github.com/jimmyfielding/maps-api-project/pkg/cache"
	"github.com/sirupsen/logrus"
)

type titleGenerator struct {
	cache      *cache.Cache
	log        *logrus.Logger
	mapsClient googlemaps.IClient
}

//NewTitleGenerator returns a pointer to a titleGenerator with cache, logger
//and googlemaps client
func NewTitleGenerator(c *cache.Cache, l *logrus.Logger, gm googlemaps.IClient) *titleGenerator {
	return &titleGenerator{
		cache:      c,
		log:        l,
		mapsClient: gm,
	}
}

//GenerateTitles returns a list of location centric titles
//for the image metadata passed in metadata
func (tg *titleGenerator) GenerateTitles(metadata []v1beta1.ImageMetadata) ([]v1beta1.Title, error) {
	var wg sync.WaitGroup
	dates := map[string]time.Time{}
	locations := v1beta1.NewSafeLocationMap() //Concurrency safe map
	var date string
	for _, m := range metadata {
		date = fmt.Sprintf("%d-%d-%d", m.Time.Year(), m.Time.Month(), m.Time.Day())
		dates[date] = *m.Time
		latlng := m.LatlngToString()
		locs, found := tg.checkCache(latlng)
		if !found {
			wg.Add(1)
			go tg.getLocations(m, locations, &wg)
			continue
		}

		for _, l := range locs {
			locations.Insert(l)
		}
	}

	wg.Wait()
	if len(locations.Locations) == 0 {
		return []v1beta1.Title{}, fmt.Errorf("failed to get locations from image metadata")
	}

	templates := generateTemplates(dates)
	titles := []v1beta1.Title{}
	var title string
	for l := range locations.Locations {
		for _, t := range templates {
			title = fmt.Sprintf(t, l)
			titles = append(titles, v1beta1.Title(title))
		}
	}

	return titles, nil
}

//checkCache returns a slice of locations and true if cache hit or
//an empty slice and false if miss. We cache on a rounded latlong
//pair concatenated as follows: 40.8-74.0 grouping all coordinates
//in a 10km zone
func (tg *titleGenerator) checkCache(latlng string) ([]v1beta1.Location, bool) {
	tg.log.Infof("checking cache for latlng pair %s", latlng)
	if locs, ok := tg.cache.Check(latlng); ok {
		tg.log.Infof("cache hit for latlng pair %s", latlng)
		return locs, true
	}
	tg.log.Infof("cache miss for latlng pair %s", latlng)
	return []v1beta1.Location{}, false
}

//getLocations gets the locations pointed to by the lines
//of latitude and longitude in the image metadata using
//reverse geocoding
func (tg *titleGenerator) getLocations(metadata v1beta1.ImageMetadata, lm *v1beta1.SafeLocationMap, wg *sync.WaitGroup) {
	defer wg.Done()
	locs, err := tg.mapsClient.ReverseGeocode(metadata.Latitude, metadata.Longitude)
	if err != nil {
		return
	}

	for _, l := range locs {
		lm.Insert(l)
	}

	tg.cache.Insert(metadata.LatlngToString(), locs)

}

//generateTemplates generates a number of string formatting
//templates to be used to provide uniform titles depending
//on date range
func generateTemplates(dates map[string]time.Time) []string {
	templates := []string{"A trip to %s"}
	months := make(map[string]string)
	weekendDays := make(map[string]string)
	for _, v := range dates {
		if _, ok := months[v.Month().String()]; !ok {
			months[v.Month().String()] = ""
		}
		if v.Weekday() == time.Saturday || v.Weekday() == time.Sunday {
			weekendDays[v.Weekday().String()] = ""
		}
	}
	if len(months) == 1 {
		for k := range months {
			templates = append(templates, "%s in "+k)
		}
	}

	if len(weekendDays) == 2 {
		templates = append(templates, "A weekend in %s")
	}

	return templates
}
