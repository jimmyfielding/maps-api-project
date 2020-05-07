package v1beta1

import (
	"fmt"
	"time"
)

//ImageMetadata is the time/date, latitude and longitude
//information about an image
type ImageMetadata struct {
	Time      *time.Time `json:"time"`
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
}

type ImageMetadataWrapper struct {
	ImageMetadata ImageMetadata `json:"metadata"`
}

type ImagesMetadataWrapper struct {
	ImageMetadata []ImageMetadata `json:"metadata"`
}

func (i ImageMetadata) LatlngToString() string {
	return fmt.Sprintf("%.1f%.1f", i.Latitude, i.Longitude)
}
