package restify

import (
	"encoding/json"
	"net/http"
)

type RestOutput struct {
	Code        int
	Description string
	Error       string
	Payload     interface{}
}

func Restify(w http.ResponseWriter, obj interface{}, err error) {

	var output RestOutput

	if err != nil {

		switch errorCode := err.Error(); errorCode {

		case http.StatusText(http.StatusNotFound):
			output.Code = http.StatusNotFound
			output.Description = http.StatusText(http.StatusNotFound)
			output.Payload = nil

		default:
			output.Code = http.StatusInternalServerError
			output.Description = http.StatusText(http.StatusInternalServerError)
			output.Error = err.Error()
		}

	} else {

		output.Code = http.StatusOK
		output.Description = http.StatusText(http.StatusOK)
		output.Payload = obj
		// output.Payload, _ = JWT.Tokerize(obj)
	}

	outputByte, _ := json.Marshal(output)

	w.Header().Set("Content-Type", "application/json")
	w.Write(outputByte)
}
