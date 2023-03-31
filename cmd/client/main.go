package main

import (
	"context"
	"fmt"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/gommon/log"
	"github.com/pablogolobaro/chequery/cmd/client/client"
	"github.com/pablogolobaro/chequery/cmd/client/client/auth_handlers"
	"github.com/pablogolobaro/chequery/cmd/client/models"
	"os"
	"os/signal"
)

func main() {
	transport := httptransport.New("localhost", "", []string{"http"})
	transport.Consumers["application/pdf"] = runtime.ByteStreamConsumer()
	apiClient := client.New(transport, strfmt.Default)
	probeOK, err := apiClient.Health.LiveProbe(nil)
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Println(*probeOK.Payload.Message)

	tokenHandlerOK, err := apiClient.AuthHandlers.TokenHandler(&auth_handlers.TokenHandlerParams{
		Body: &models.AuthDataRequest{
			Password: "password",
			Username: "JohnSnow",
		},
		Context: context.Background(),
	})
	if err != nil {
		log.Error(err)
		return
	}
	token := tokenHandlerOK.GetPayload().Token

	worker := NewWorker(apiClient, token)

	go worker.Start()

	sigCh := make(chan os.Signal, 0)

	signal.Notify(sigCh)

	<-sigCh

	log.Print("App is stopping...")

	worker.Stop()

}
