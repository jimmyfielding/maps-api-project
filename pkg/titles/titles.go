package titles

import (
	"net/http"

	"github.com/jimmyfielding/maps-api-project/pkg/api/v1beta1"
)

func (c *Client) GenerateTitles(metadata []v1beta1.ImageMetadata) ([]v1beta1.Title, *http.Response, error) {
	endpoint := "titles"
	response := new(v1beta1.TitlesWrapper)

	body := v1beta1.ImagesMetadataWrapper{
		ImageMetadata: metadata,
	}

	req, err := c.NewRequest("POST", endpoint, body)
	if err != nil {
		return response.Titles, nil, err
	}

	httpResp, err := c.Do(req, response)
	if err != nil {
		return response.Titles, httpResp, err
	}

	return response.Titles, httpResp, nil
}
