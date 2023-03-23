package main

import (
	"fmt"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/gommon/log"
	"github.com/pablogolobaro/chequery/cmd/client/client"
	"os"
	"os/signal"
)

func main() {
	transport := httptransport.New("localhost", "api/v1", []string{"http"})
	transport.Consumers["application/pdf"] = runtime.ByteStreamConsumer()
	apiClient := client.New(transport, strfmt.Default)
	probeOK, err := apiClient.Health.LiveProbe(nil)
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Println(*probeOK.Payload.Message)

	worker := NewWorker(apiClient)

	go worker.Start()

	sigCh := make(chan os.Signal, 0)

	signal.Notify(sigCh)

	<-sigCh

	log.Print("App is stopping...")

	worker.Stop()

}
