package main

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s *server) sendJSONResponse(wrapper interface{}, w http.ResponseWriter) error {
	if err := json.NewEncoder(w).Encode(wrapper); err != nil {
		s.log.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("failed to send JSON response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}
