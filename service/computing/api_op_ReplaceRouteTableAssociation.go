// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type ReplaceRouteTableAssociationInput struct {
	_ struct{} `type:"structure"`

	Agreement *bool `locationName:"Agreement" type:"boolean"`

	// AssociationId is a required field
	AssociationId *string `locationName:"AssociationId" type:"string" required:"true"`

	// RouteTableId is a required field
	RouteTableId *string `locationName:"RouteTableId" type:"string" required:"true"`
}

// String returns the string representation
func (s ReplaceRouteTableAssociationInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReplaceRouteTableAssociationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReplaceRouteTableAssociationInput"}

	if s.AssociationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssociationId"))
	}

	if s.RouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ReplaceRouteTableAssociationOutput struct {
	_ struct{} `type:"structure"`

	NewAssociationId *string `locationName:"newAssociationId" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s ReplaceRouteTableAssociationOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opReplaceRouteTableAssociation = "ReplaceRouteTableAssociation"

// ReplaceRouteTableAssociationRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using ReplaceRouteTableAssociationRequest.
//    req := client.ReplaceRouteTableAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/ReplaceRouteTableAssociation.htm
func (c *Client) ReplaceRouteTableAssociationRequest(input *ReplaceRouteTableAssociationInput) ReplaceRouteTableAssociationRequest {
	op := &aws.Operation{
		Name:       opReplaceRouteTableAssociation,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &ReplaceRouteTableAssociationInput{}
	}

	req := c.newRequest(op, input, &ReplaceRouteTableAssociationOutput{})

	return ReplaceRouteTableAssociationRequest{Request: req, Input: input, Copy: c.ReplaceRouteTableAssociationRequest}
}

// ReplaceRouteTableAssociationRequest is the request type for the
// ReplaceRouteTableAssociation API operation.
type ReplaceRouteTableAssociationRequest struct {
	*aws.Request
	Input *ReplaceRouteTableAssociationInput
	Copy  func(*ReplaceRouteTableAssociationInput) ReplaceRouteTableAssociationRequest
}

// Send marshals and sends the ReplaceRouteTableAssociation API request.
func (r ReplaceRouteTableAssociationRequest) Send(ctx context.Context) (*ReplaceRouteTableAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReplaceRouteTableAssociationResponse{
		ReplaceRouteTableAssociationOutput: r.Request.Data.(*ReplaceRouteTableAssociationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReplaceRouteTableAssociationResponse is the response type for the
// ReplaceRouteTableAssociation API operation.
type ReplaceRouteTableAssociationResponse struct {
	*ReplaceRouteTableAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReplaceRouteTableAssociation request.
func (r *ReplaceRouteTableAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
