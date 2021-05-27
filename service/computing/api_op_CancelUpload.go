// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type CancelUploadInput struct {
	_ struct{} `type:"structure"`

	// ConversionTaskId is a required field
	ConversionTaskId *string `locationName:"ConversionTaskId" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelUploadInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelUploadInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelUploadInput"}

	if s.ConversionTaskId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConversionTaskId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CancelUploadOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s CancelUploadOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opCancelUpload = "CancelUpload"

// CancelUploadRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using CancelUploadRequest.
//    req := client.CancelUploadRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/CancelUpload.htm
func (c *Client) CancelUploadRequest(input *CancelUploadInput) CancelUploadRequest {
	op := &aws.Operation{
		Name:       opCancelUpload,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &CancelUploadInput{}
	}

	req := c.newRequest(op, input, &CancelUploadOutput{})

	return CancelUploadRequest{Request: req, Input: input, Copy: c.CancelUploadRequest}
}

// CancelUploadRequest is the request type for the
// CancelUpload API operation.
type CancelUploadRequest struct {
	*aws.Request
	Input *CancelUploadInput
	Copy  func(*CancelUploadInput) CancelUploadRequest
}

// Send marshals and sends the CancelUpload API request.
func (r CancelUploadRequest) Send(ctx context.Context) (*CancelUploadResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelUploadResponse{
		CancelUploadOutput: r.Request.Data.(*CancelUploadOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelUploadResponse is the response type for the
// CancelUpload API operation.
type CancelUploadResponse struct {
	*CancelUploadOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelUpload request.
func (r *CancelUploadResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
