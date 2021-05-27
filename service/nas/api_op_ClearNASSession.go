// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package nas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type ClearNASSessionInput struct {
	_ struct{} `type:"structure"`

	// NASInstanceIdentifier is a required field
	NASInstanceIdentifier *string `locationName:"NASInstanceIdentifier" type:"string" required:"true"`

	SessionClearType *string `locationName:"SessionClearType" type:"string"`
}

// String returns the string representation
func (s ClearNASSessionInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ClearNASSessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ClearNASSessionInput"}

	if s.NASInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("NASInstanceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ClearNASSessionOutput struct {
	_ struct{} `type:"structure"`

	NASInstance *NASInstance `type:"structure"`
}

// String returns the string representation
func (s ClearNASSessionOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opClearNASSession = "ClearNASSession"

// ClearNASSessionRequest returns a request value for making API operation for
// NIFCLOUD NAS.
//
//    // Example sending a request using ClearNASSessionRequest.
//    req := client.ClearNASSessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/nas/ClearNASSession.htm
func (c *Client) ClearNASSessionRequest(input *ClearNASSessionInput) ClearNASSessionRequest {
	op := &aws.Operation{
		Name:       opClearNASSession,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ClearNASSessionInput{}
	}

	req := c.newRequest(op, input, &ClearNASSessionOutput{})

	return ClearNASSessionRequest{Request: req, Input: input, Copy: c.ClearNASSessionRequest}
}

// ClearNASSessionRequest is the request type for the
// ClearNASSession API operation.
type ClearNASSessionRequest struct {
	*aws.Request
	Input *ClearNASSessionInput
	Copy  func(*ClearNASSessionInput) ClearNASSessionRequest
}

// Send marshals and sends the ClearNASSession API request.
func (r ClearNASSessionRequest) Send(ctx context.Context) (*ClearNASSessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ClearNASSessionResponse{
		ClearNASSessionOutput: r.Request.Data.(*ClearNASSessionOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ClearNASSessionResponse is the response type for the
// ClearNASSession API operation.
type ClearNASSessionResponse struct {
	*ClearNASSessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ClearNASSession request.
func (r *ClearNASSessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
