package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type MalformedRequest struct {
	Status int
	Msg    string
}

func (mr *MalformedRequest) Error() string {
	return mr.Msg
}

func DecodeJSON(w http.ResponseWriter, r *http.Request, dst interface{}) {
	err := decodeJSONBody(w, r, &dst)
	if err != nil {
		var mr *MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
}

func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	contentType := r.Header.Get("Content-Type")
	if contentType != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(contentType, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-Type header is not application/json"
			return &MalformedRequest{Status: http.StatusUnsupportedMediaType, Msg: msg}
		}
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &MalformedRequest{Status: http.StatusRequestEntityTooLarge, Msg: msg}
		default:
			return err
		}
	}
	err = decoder.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		msg := "Request body must only contain a single JSON object"
		return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
	}
	return nil
}
