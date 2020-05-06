package main

import (
	"encoding/json"
	"net/http"

	"github.com/jimmyfielding/maps-api-project/pkg/api/v1beta1"
	"github.com/sirupsen/logrus"
)

func (s *server) getTitles(w http.ResponseWriter, r *http.Request) {
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

	titles, err := s.mapsClient.GenerateTitles(body.ImageMetadata)
	if err != nil {
		response := v1beta1.ErrorWrapper{Error: err.Error()}
		if err = s.sendJSONResponse(response, w); err != nil {
			s.log.WithFields(logrus.Fields{
				"err": err.Error(),
			}).Error("failed to generate titles")
		}
	}

	titleWrapper := v1beta1.TitlesWrapper{
		Titles: titles,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(titleWrapper); err != nil {
		s.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
