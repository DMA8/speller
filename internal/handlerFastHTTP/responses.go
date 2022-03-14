package handlerFastHTTP

import (
	"encoding/json"
	"log"
	"spellCheck/internal/storage"
)

type badResponse struct {
	ErrorMassage string `json:"error"`
}

type resultResponse struct {
	Result storage.Spelling `json:"result"`
}

type addmitResponse struct {
	Result string `json:"result"`
}

func makeJSONErrorResponse(msg string) []byte {
	result, err := json.Marshal(badResponse{msg})
	if err != nil {
		log.Print(err)
	}
	return result
}

func makeJSONResultResponse(resultSpell storage.Spelling) []byte{
	result, err := json.Marshal(resultResponse{resultSpell})
	if err != nil {
		log.Print(err)
	}
	return result
}

func makeJSONAddmit(msg string) []byte {
	result, err := json.Marshal(addmitResponse{msg})
	if err != nil {
		log.Print(err)
	}
	return result
}