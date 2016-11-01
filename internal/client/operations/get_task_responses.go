package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/aws/amazon-ecs-event-stream-handler/internal/models"
)

// GetTaskReader is a Reader for the GetTask structure.
type GetTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTaskOK creates a GetTaskOK with default headers values
func NewGetTaskOK() *GetTaskOK {
	return &GetTaskOK{}
}

/*GetTaskOK handles this case with default header values.

Get task using cluster name and task ARN - success
*/
type GetTaskOK struct {
	Payload *models.TaskModel
}

func (o *GetTaskOK) Error() string {
	return fmt.Sprintf("[GET /tasks/{cluster}/{arn}][%d] getTaskOK  %+v", 200, o.Payload)
}

func (o *GetTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskNotFound creates a GetTaskNotFound with default headers values
func NewGetTaskNotFound() *GetTaskNotFound {
	return &GetTaskNotFound{}
}

/*GetTaskNotFound handles this case with default header values.

Get task using cluster name and task ARN - task not found
*/
type GetTaskNotFound struct {
	Payload string
}

func (o *GetTaskNotFound) Error() string {
	return fmt.Sprintf("[GET /tasks/{cluster}/{arn}][%d] getTaskNotFound  %+v", 404, o.Payload)
}

func (o *GetTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskInternalServerError creates a GetTaskInternalServerError with default headers values
func NewGetTaskInternalServerError() *GetTaskInternalServerError {
	return &GetTaskInternalServerError{}
}

/*GetTaskInternalServerError handles this case with default header values.

Get task using cluster name and task ARN - unexpected error
*/
type GetTaskInternalServerError struct {
	Payload string
}

func (o *GetTaskInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tasks/{cluster}/{arn}][%d] getTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
