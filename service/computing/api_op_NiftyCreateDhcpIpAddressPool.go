// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type NiftyCreateDhcpIpAddressPoolInput struct {
	_ struct{} `type:"structure"`

	Description *string `locationName:"Description" type:"string"`

	// DhcpConfigId is a required field
	DhcpConfigId *string `locationName:"DhcpConfigId" type:"string" required:"true"`

	// StartIpAddress is a required field
	StartIpAddress *string `locationName:"StartIpAddress" type:"string" required:"true"`

	// StopIpAddress is a required field
	StopIpAddress *string `locationName:"StopIpAddress" type:"string" required:"true"`
}

// String returns the string representation
func (s NiftyCreateDhcpIpAddressPoolInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NiftyCreateDhcpIpAddressPoolInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "NiftyCreateDhcpIpAddressPoolInput"}

	if s.DhcpConfigId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DhcpConfigId"))
	}

	if s.StartIpAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartIpAddress"))
	}

	if s.StopIpAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("StopIpAddress"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type NiftyCreateDhcpIpAddressPoolOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s NiftyCreateDhcpIpAddressPoolOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyCreateDhcpIpAddressPool = "NiftyCreateDhcpIpAddressPool"

// NiftyCreateDhcpIpAddressPoolRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyCreateDhcpIpAddressPoolRequest.
//    req := client.NiftyCreateDhcpIpAddressPoolRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/NiftyCreateDhcpIpAddressPool.htm
func (c *Client) NiftyCreateDhcpIpAddressPoolRequest(input *NiftyCreateDhcpIpAddressPoolInput) NiftyCreateDhcpIpAddressPoolRequest {
	op := &aws.Operation{
		Name:       opNiftyCreateDhcpIpAddressPool,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyCreateDhcpIpAddressPoolInput{}
	}

	req := c.newRequest(op, input, &NiftyCreateDhcpIpAddressPoolOutput{})

	return NiftyCreateDhcpIpAddressPoolRequest{Request: req, Input: input, Copy: c.NiftyCreateDhcpIpAddressPoolRequest}
}

// NiftyCreateDhcpIpAddressPoolRequest is the request type for the
// NiftyCreateDhcpIpAddressPool API operation.
type NiftyCreateDhcpIpAddressPoolRequest struct {
	*aws.Request
	Input *NiftyCreateDhcpIpAddressPoolInput
	Copy  func(*NiftyCreateDhcpIpAddressPoolInput) NiftyCreateDhcpIpAddressPoolRequest
}

// Send marshals and sends the NiftyCreateDhcpIpAddressPool API request.
func (r NiftyCreateDhcpIpAddressPoolRequest) Send(ctx context.Context) (*NiftyCreateDhcpIpAddressPoolResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyCreateDhcpIpAddressPoolResponse{
		NiftyCreateDhcpIpAddressPoolOutput: r.Request.Data.(*NiftyCreateDhcpIpAddressPoolOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyCreateDhcpIpAddressPoolResponse is the response type for the
// NiftyCreateDhcpIpAddressPool API operation.
type NiftyCreateDhcpIpAddressPoolResponse struct {
	*NiftyCreateDhcpIpAddressPoolOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyCreateDhcpIpAddressPool request.
func (r *NiftyCreateDhcpIpAddressPoolResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
