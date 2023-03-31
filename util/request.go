package util

import (
	"encoding/json"
	LOGGER "hdfc-backend/logger"
	"io/ioutil"
	"net/http"
)

// ReadRequestBody - read raw body from request
func ReadRequestBody(r *http.Request, body interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		LOGGER.Warn("ReadRequestBody", err)
		return err
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, body)
	if err != nil {
		LOGGER.Warn("ReadRequestBody", err)
		return err
	}
	return nil
}
