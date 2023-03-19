package main

import (
	"context"
	"github.com/pablogolobaro/chequery/cmd/client/client"
	"github.com/pablogolobaro/chequery/cmd/client/client/order"
	"github.com/pablogolobaro/chequery/cmd/client/models"
	"github.com/stretchr/testify/assert"
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
