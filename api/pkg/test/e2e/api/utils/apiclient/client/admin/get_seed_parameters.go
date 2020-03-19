// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSeedParams creates a new GetSeedParams object
// with the default values initialized.
func NewGetSeedParams() *GetSeedParams {
	var ()
	return &GetSeedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSeedParamsWithTimeout creates a new GetSeedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSeedParamsWithTimeout(timeout time.Duration) *GetSeedParams {
	var ()
	return &GetSeedParams{

		timeout: timeout,
	}
}

// NewGetSeedParamsWithContext creates a new GetSeedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSeedParamsWithContext(ctx context.Context) *GetSeedParams {
	var ()
	return &GetSeedParams{

		Context: ctx,
	}
}

// NewGetSeedParamsWithHTTPClient creates a new GetSeedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSeedParamsWithHTTPClient(client *http.Client) *GetSeedParams {
	var ()
	return &GetSeedParams{
		HTTPClient: client,
	}
}

/*GetSeedParams contains all the parameters to send to the API endpoint
for the get seed operation typically these are written to a http.Request
*/
type GetSeedParams struct {

	/*SeedName*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get seed params
func (o *GetSeedParams) WithTimeout(timeout time.Duration) *GetSeedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get seed params
func (o *GetSeedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get seed params
func (o *GetSeedParams) WithContext(ctx context.Context) *GetSeedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get seed params
func (o *GetSeedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get seed params
func (o *GetSeedParams) WithHTTPClient(client *http.Client) *GetSeedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get seed params
func (o *GetSeedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the seedName to the get seed params
func (o *GetSeedParams) WithName(seedName string) *GetSeedParams {
	o.SetName(seedName)
	return o
}

// SetName adds the seedName to the get seed params
func (o *GetSeedParams) SetName(seedName string) {
	o.Name = seedName
}

// WriteToRequest writes these params to a swagger request
func (o *GetSeedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param seed_name
	if err := r.SetPathParam("seed_name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}