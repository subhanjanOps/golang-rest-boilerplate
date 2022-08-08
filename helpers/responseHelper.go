package helpers

import (
	"encoding/json"
	statuscodes "golang-rest/statusCodes"
	"log"
)

func ParseResponse(data statuscodes.ResponseBody) []byte {
	res, err := json.Marshal(&data)
	if err != nil {
		log.Println("Error marshalling data:", err)
		return nil
	}
	return res
}
