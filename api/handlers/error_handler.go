package handlers

import "encoding/json"

func formatJSONerror(message string) []byte {
	appError := struct {
		Message string `json:"message"`
	}{
		message,
	}

	response, err := json.Marshal(appError)
	if err != nil {
		return []byte(err.Error())
	}

	return response
}
