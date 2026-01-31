package main

import (
	"fmt"
	"log"

	greenapi "github.com/green-api/telegram-api-client-golang"
)

func main() {
	GreenAPI := greenapi.GreenAPI{
		APIURL:           "https://4100.api.green-api.com",
		MediaURL:         "https://4100.api.green-api.com",
		IDInstance:       "4100000000",
		APITokenInstance: "d75b3a66374942c5b3c019c698abc2067e151558acbd412345",
	}

	response, err := GreenAPI.Account().Qr()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status: %v %s \n\rResponse: %s\n\rTimestamp: %s\n\r", response.StatusCode,
		response.StatusMessage,
		response.Body,
		response.Timestamp.Format("15:04:05.000"))
}
