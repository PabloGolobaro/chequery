// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new order API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for order API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateOrder(params *CreateOrderParams, opts ...ClientOption) (*CreateOrderOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateOrder creates checks from new order

This will create new checks in db and starts generating pdf files from it.
*/
func (a *Client) CreateOrder(params *CreateOrderParams, opts ...ClientOption) (*CreateOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createOrder",
		Method:             "POST",
		PathPattern:        "/order",
		ProducesMediaTypes: []string{"application/json", "application/pdf"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateOrderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateOrderDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}