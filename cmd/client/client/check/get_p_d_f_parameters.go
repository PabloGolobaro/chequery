// Code generated by go-swagger; DO NOT EDIT.

package check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetPDFParams creates a new GetPDFParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPDFParams() *GetPDFParams {
	return &GetPDFParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPDFParamsWithTimeout creates a new GetPDFParams object
// with the ability to set a timeout on a request.
func NewGetPDFParamsWithTimeout(timeout time.Duration) *GetPDFParams {
	return &GetPDFParams{
		timeout: timeout,
	}
}

// NewGetPDFParamsWithContext creates a new GetPDFParams object
// with the ability to set a context for a request.
func NewGetPDFParamsWithContext(ctx context.Context) *GetPDFParams {
	return &GetPDFParams{
		Context: ctx,
	}
}

// NewGetPDFParamsWithHTTPClient creates a new GetPDFParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPDFParamsWithHTTPClient(client *http.Client) *GetPDFParams {
	return &GetPDFParams{
		HTTPClient: client,
	}
}

/*
GetPDFParams contains all the parameters to send to the API endpoint

	for the get p d f operation.

	Typically these are written to a http.Request.
*/
type GetPDFParams struct {

	/* CheckID.

	   The ID of a check

	   Format: int64
	*/
	CheckID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get p d f params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPDFParams) WithDefaults() *GetPDFParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get p d f params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPDFParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get p d f params
func (o *GetPDFParams) WithTimeout(timeout time.Duration) *GetPDFParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get p d f params
func (o *GetPDFParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get p d f params
func (o *GetPDFParams) WithContext(ctx context.Context) *GetPDFParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get p d f params
func (o *GetPDFParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get p d f params
func (o *GetPDFParams) WithHTTPClient(client *http.Client) *GetPDFParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get p d f params
func (o *GetPDFParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCheckID adds the checkID to the get p d f params
func (o *GetPDFParams) WithCheckID(checkID int64) *GetPDFParams {
	o.SetCheckID(checkID)
	return o
}

// SetCheckID adds the checkId to the get p d f params
func (o *GetPDFParams) SetCheckID(checkID int64) {
	o.CheckID = checkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPDFParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param check_id
	if err := r.SetPathParam("check_id", swag.FormatInt64(o.CheckID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
