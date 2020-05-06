package v1beta1

import "time"

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
