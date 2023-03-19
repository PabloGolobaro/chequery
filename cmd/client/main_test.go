package main

import (
	"bytes"
	"context"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/pablogolobaro/chequery/cmd/client/client"
	"github.com/pablogolobaro/chequery/cmd/client/client/check"
	"github.com/pablogolobaro/chequery/cmd/client/client/order"
	"github.com/pablogolobaro/chequery/cmd/client/models"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestLiveProbe(t *testing.T) {
	apiClient := client.NewHTTPClient(nil)
	probeOK, err := apiClient.Health.LiveProbe(nil)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, 200, probeOK.Code())
	assert.Equal(t, "I'm alive", *probeOK.Payload.Message)

}

func TestOrderCreation(t *testing.T) {
	apiClient := client.NewHTTPClient(nil)
	createOrder, err := apiClient.Order.CreateOrder(&order.CreateOrderParams{
		Context: context.Background(),
		PointID: 1,
		Body: &models.OrderCreateRequest{
			Order: "{\"meat\":\"25kg\"}",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, 201, createOrder.Code())
	t.Log(createOrder.String())

}

func TestGetGenerated(t *testing.T) {
	apiClient := client.NewHTTPClient(nil)
	getGenerated, err := apiClient.Check.GetGenerated(nil)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, 200, getGenerated.Code())
	t.Log(getGenerated.String())

}

func TestUpdateStatus(t *testing.T) {
	apiClient := client.NewHTTPClient(nil)
	getGenerated, err := apiClient.Check.UpdateChecksStatus(&check.UpdateChecksStatusParams{IDs: []int64{16, 17}, Context: context.Background()})
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, 200, getGenerated.Code())
	t.Log(getGenerated.String())
}

func TestGetPDF(t *testing.T) {
	transport := httptransport.New("localhost", "api/v1", []string{"http"})
	transport.Consumers["application/pdf"] = runtime.ByteStreamConsumer()

	apiClient := client.New(transport, strfmt.Default)
	buffer := &bytes.Buffer{}

	getPDF, err := apiClient.Check.GetPDF(&check.GetPDFParams{CheckID: 16, Context: context.Background()}, buffer)
	if err != nil {
		t.Log(err.Error())
	}
	assert.Equal(t, 200, getPDF.Code())
	create, err := os.Create("example.pdf")
	if err != nil {
		t.Error(err)
		return
	}
	c, err := io.Copy(create, buffer)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(c)
}
