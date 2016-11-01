package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTaskParams creates a new GetTaskParams object
// with the default values initialized.
func NewGetTaskParams() *GetTaskParams {
	var ()
	return &GetTaskParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTaskParamsWithTimeout creates a new GetTaskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTaskParamsWithTimeout(timeout time.Duration) *GetTaskParams {
	var ()
	return &GetTaskParams{

		timeout: timeout,
	}
}

// NewGetTaskParamsWithContext creates a new GetTaskParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTaskParamsWithContext(ctx context.Context) *GetTaskParams {
	var ()
	return &GetTaskParams{

		Context: ctx,
	}
}

/*GetTaskParams contains all the parameters to send to the API endpoint
for the get task operation typically these are written to a http.Request
*/
type GetTaskParams struct {

	/*Arn
	  ARN of the task to fetch

	*/
	Arn string
	/*Cluster
	  Cluster name of the task to fetch

	*/
	Cluster string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get task params
func (o *GetTaskParams) WithTimeout(timeout time.Duration) *GetTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get task params
func (o *GetTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get task params
func (o *GetTaskParams) WithContext(ctx context.Context) *GetTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get task params
func (o *GetTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithArn adds the arn to the get task params
func (o *GetTaskParams) WithArn(arn string) *GetTaskParams {
	o.SetArn(arn)
	return o
}

// SetArn adds the arn to the get task params
func (o *GetTaskParams) SetArn(arn string) {
	o.Arn = arn
}

// WithCluster adds the cluster to the get task params
func (o *GetTaskParams) WithCluster(cluster string) *GetTaskParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the get task params
func (o *GetTaskParams) SetCluster(cluster string) {
	o.Cluster = cluster
}

// WriteToRequest writes these params to a swagger request
func (o *GetTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param arn
	if err := r.SetPathParam("arn", o.Arn); err != nil {
		return err
	}

	// path param cluster
	if err := r.SetPathParam("cluster", o.Cluster); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
