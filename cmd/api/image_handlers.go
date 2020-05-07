package main

import (
	"encoding/json"
	"net/http"

	"github.com/jimmyfielding/maps-api-project/internal/title"
	"github.com/jimmyfielding/maps-api-project/pkg/api/v1beta1"
	"github.com/jimmyfielding/maps-api-project/pkg/cache"
	"github.com/sirupsen/logrus"
)

func (s *server) generateTitles(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var body v1beta1.ImagesMetadataWrapper
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response := v1beta1.ErrorWrapper{Error: v1beta1.Error{Message: err.Error()}}
		if err = s.sendJSONResponse(response, w); err != nil {
			s.log.WithFields(logrus.Fields{
				"err": err.Error(),
			}).Error("failed to send image metadata")
		}
		return
	}

	cache := cache.NewCache()
	titleGenerator := title.NewTitleGenerator(cache, s.log, s.mapsClient)
	titles, err := titleGenerator.GenerateTitles(body.ImageMetadata)
	if err != nil {
		response := v1beta1.ErrorWrapper{Error: v1beta1.Error{Message: err.Error()}}
		if err = s.sendJSONResponse(response, w); err != nil {
			s.log.WithFields(logrus.Fields{
				"err": err.Error(),
			}).Error("failed to generate titles")
		}
	}

	titleWrapper := v1beta1.TitlesWrapper{
		Titles: titles,
	}

	if err := json.NewEncoder(w).Encode(titleWrapper); err != nil {
		s.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
