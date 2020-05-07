package cmd

import (
	"net/http"

	"github.com/jimmyfielding/maps-api-project/pkg/api/v1beta1"
	"github.com/jimmyfielding/maps-api-project/pkg/titles"
	"github.com/sirupsen/logrus"
)

type TitlesClient struct {
	client  titles.Client
	baseURL string
}

// NewTitlesClient is a light wrapper around the titles
// client library.
func NewTitlesClient(baseURL string, log logrus.FieldLogger) (*TitlesClient, error) {
	return &TitlesClient{
		client:  *titles.NewClient(baseURL, logrus.New()),
		baseURL: baseURL,
	}, nil
}

//GenerateTitles returns a slice of location centric titles given a slice
//of image metadata
func (t *TitlesClient) GenerateTitles(metadata []v1beta1.ImageMetadata) ([]v1beta1.Title, *http.Response, error) {
	return t.client.GenerateTitles(metadata)
}
