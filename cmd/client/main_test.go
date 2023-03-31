package main

import (
	"bytes"
	"context"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/pablogolobaro/chequery/cmd/client/client"
	"github.com/pablogolobaro/chequery/cmd/client/client/auth_handlers"
	"github.com/pablogolobaro/chequery/cmd/client/client/check"
	"github.com/pablogolobaro/chequery/cmd/client/client/order"
	"github.com/pablogolobaro/chequery/cmd/client/models"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

const tokenJWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkpvaG5Tbm93IiwiZXhwIjoxNjgwNTQ2NTk2fQ.u1pURDXEat-8zPl80el6zLBbjbovVdbQS44pPTbm3Aw"

func TestAuthPostData(t *testing.T) {
	apiClient := client.NewHTTPClient(nil)
	tokenHandlerOK, err := apiClient.AuthHandlers.TokenHandler(&auth_handlers.TokenHandlerParams{
		Body: &models.AuthDataRequest{
			Password: "password",
			Username: "JohnSnow",
		},
		Context: context.Background(),
	})
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, true, tokenHandlerOK.IsCode(200))
	t.Log(tokenHandlerOK.GetPayload().Token)

}

func TestLiveProbe(t *testing.T) {
	apiClient := client.NewHTTPClient(nil)
	probeOK, err := apiClient.Health.LiveProbe(nil)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, true, probeOK.IsCode(200))
	assert.Equal(t, "I'm alive", *probeOK.Payload.Message)

}

func TestOrderCreation(t *testing.T) {
	requestOrder := &models.Order{
		PointID: 1,
		Products: []*models.Product{
			{Name: "Meat", Quantity: 3, Price: 145},
			{Name: "vegetables", Quantity: 2, Price: 32},
			{Name: "Juice", Quantity: 1, Price: 48}},
	}

	apiClient := client.NewHTTPClient(nil)
	createOrder, err := apiClient.Order.CreateOrder(
		&order.CreateOrderParams{
			Body:    &models.OrderCreateRequest{Order: requestOrder},
			Context: context.Background(),
		},
		httptransport.BearerToken(tokenJWT),
	)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, true, createOrder.IsCode(201))
	t.Log(createOrder.String())

}

func TestGetGenerated(t *testing.T) {
	apiClient := client.NewHTTPClient(nil)
	getGenerated, err := apiClient.Check.GetGenerated(
		nil,
		httptransport.BearerToken(tokenJWT),
	)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, true, getGenerated.IsCode(200))
	t.Log(getGenerated.String())

}

func TestUpdateStatus(t *testing.T) {
	apiClient := client.NewHTTPClient(nil)
	getGenerated, err := apiClient.Check.UpdateChecksStatus(
		&check.UpdateChecksStatusParams{IDs: []int64{16, 17}, Context: context.Background()},
		httptransport.BearerToken(tokenJWT),
	)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, true, getGenerated.IsCode(200))
	t.Log(getGenerated.String())
}

func TestGetPDF(t *testing.T) {
	transport := httptransport.New("localhost", "", []string{"http"})
	transport.Consumers["application/pdf"] = runtime.ByteStreamConsumer()

	apiClient := client.New(transport, strfmt.Default)
	buffer := &bytes.Buffer{}

	getPDF, err := apiClient.Check.GetPDF(
		&check.GetPDFParams{CheckID: 16, Context: context.Background()},
		httptransport.BearerToken(tokenJWT),
		buffer,
	)
	if err != nil {
		t.Log(err.Error())
	}
	assert.Equal(t, true, getPDF.IsCode(200))
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
