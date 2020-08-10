// Code generated by go-swagger; DO NOT EDIT.

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dreadl0ck/firecracker-go-sdk/client/models"
)

// PatchMmdsReader is a Reader for the PatchMmds structure.
type PatchMmdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchMmdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchMmdsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchMmdsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchMmdsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchMmdsNoContent creates a PatchMmdsNoContent with default headers values
func NewPatchMmdsNoContent() *PatchMmdsNoContent {
	return &PatchMmdsNoContent{}
}

/*PatchMmdsNoContent handles this case with default header values.

MMDS data store updated.
*/
type PatchMmdsNoContent struct {
}

func (o *PatchMmdsNoContent) Error() string {
	return fmt.Sprintf("[PATCH /mmds][%d] patchMmdsNoContent ", 204)
}

func (o *PatchMmdsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchMmdsBadRequest creates a PatchMmdsBadRequest with default headers values
func NewPatchMmdsBadRequest() *PatchMmdsBadRequest {
	return &PatchMmdsBadRequest{}
}

/*PatchMmdsBadRequest handles this case with default header values.

MMDS data store cannot be updated due to bad input.
*/
type PatchMmdsBadRequest struct {
	Payload *models.Error
}

func (o *PatchMmdsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /mmds][%d] patchMmdsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchMmdsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchMmdsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMmdsDefault creates a PatchMmdsDefault with default headers values
func NewPatchMmdsDefault(code int) *PatchMmdsDefault {
	return &PatchMmdsDefault{
		_statusCode: code,
	}
}

/*PatchMmdsDefault handles this case with default header values.

Internal server error
*/
type PatchMmdsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch mmds default response
func (o *PatchMmdsDefault) Code() int {
	return o._statusCode
}

func (o *PatchMmdsDefault) Error() string {
	return fmt.Sprintf("[PATCH /mmds][%d] PatchMmds default  %+v", o._statusCode, o.Payload)
}

func (o *PatchMmdsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchMmdsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
