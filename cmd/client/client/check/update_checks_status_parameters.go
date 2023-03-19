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

// NewUpdateChecksStatusParams creates a new UpdateChecksStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateChecksStatusParams() *UpdateChecksStatusParams {
	return &UpdateChecksStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateChecksStatusParamsWithTimeout creates a new UpdateChecksStatusParams object
// with the ability to set a timeout on a request.
func NewUpdateChecksStatusParamsWithTimeout(timeout time.Duration) *UpdateChecksStatusParams {
	return &UpdateChecksStatusParams{
		timeout: timeout,
	}
}

// NewUpdateChecksStatusParamsWithContext creates a new UpdateChecksStatusParams object
// with the ability to set a context for a request.
func NewUpdateChecksStatusParamsWithContext(ctx context.Context) *UpdateChecksStatusParams {
	return &UpdateChecksStatusParams{
		Context: ctx,
	}
}

// NewUpdateChecksStatusParamsWithHTTPClient creates a new UpdateChecksStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateChecksStatusParamsWithHTTPClient(client *http.Client) *UpdateChecksStatusParams {
	return &UpdateChecksStatusParams{
		HTTPClient: client,
	}
}

/*
UpdateChecksStatusParams contains all the parameters to send to the API endpoint

	for the update checks status operation.

	Typically these are written to a http.Request.
*/
type UpdateChecksStatusParams struct {

	/* ID.

	   The IDs of a checks set status printed
	*/
	IDs []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update checks status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateChecksStatusParams) WithDefaults() *UpdateChecksStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update checks status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateChecksStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update checks status params
func (o *UpdateChecksStatusParams) WithTimeout(timeout time.Duration) *UpdateChecksStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update checks status params
func (o *UpdateChecksStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update checks status params
func (o *UpdateChecksStatusParams) WithContext(ctx context.Context) *UpdateChecksStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update checks status params
func (o *UpdateChecksStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update checks status params
func (o *UpdateChecksStatusParams) WithHTTPClient(client *http.Client) *UpdateChecksStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update checks status params
func (o *UpdateChecksStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIDs adds the id to the update checks status params
func (o *UpdateChecksStatusParams) WithIDs(id []string) *UpdateChecksStatusParams {
	o.SetIDs(id)
	return o
}

// SetIDs adds the id to the update checks status params
func (o *UpdateChecksStatusParams) SetIDs(id []string) {
	o.IDs = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateChecksStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IDs != nil {

		// binding items for id
		joinedID := o.bindParamID(reg)

		// query array param id
		if err := r.SetQueryParam("id", joinedID...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamUpdateChecksStatus binds the parameter id
func (o *UpdateChecksStatusParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.IDs

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []string

		iDIIV := iDIIR // string as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: ""
	iDIS := swag.JoinByFormat(iDIC, "")

	return iDIS
}
