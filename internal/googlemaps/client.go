package googlemaps

import (
	"context"

	"googlemaps.github.io/maps"

	"github.com/jimmyfielding/maps-api-project/pkg/api/v1beta1"
)

type client struct {
	gm  *maps.Client
	ctx context.Context
}

type IClient interface {
	ReverseGeocode(lat, lng float64) ([]v1beta1.Location, error)
}

//NewClient returns a new client for interacting with the google maps API
//using the passed API key
func NewClient(ctx context.Context, apiKey string) (IClient, error) {
	mapsClient, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return &client{}, err
	}

	return &client{
		gm:  mapsClient,
		ctx: ctx,
	}, nil
}

//ReverseGeocode returns a slice of locations pointed to by the lines of latitude
//and longitude passed that are considered to be countries, cities, counties or
//small divisions within a city
func (c *client) ReverseGeocode(lat, lng float64) ([]v1beta1.Location, error) {
	req := maps.GeocodingRequest{
		LatLng: &maps.LatLng{
			Lat: lat,
			Lng: lng,
		},
		ResultType: []string{
			"country",
			"locality",
			"sublocality_level_1",
		},
	}

	res, err := c.gm.Geocode(c.ctx, &req)
	if err != nil {
		return []v1beta1.Location{}, err
	}

	locations := []v1beta1.Location{}
	for _, r := range res {
		for _, l := range r.AddressComponents {
			locations = append(locations, v1beta1.Location(l.LongName))
		}
	}

	return locations, nil
}
