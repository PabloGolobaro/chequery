package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/pablogolobaro/chequery/cmd/client/client"
)

func main() {
	apiClient := client.NewHTTPClient(nil)
	probeOK, err := apiClient.Health.LiveProbe(nil)
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Println(probeOK.Code())
	fmt.Println(*probeOK.Payload.Message)

}
