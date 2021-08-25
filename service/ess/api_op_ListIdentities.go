// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ess

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type ListIdentitiesInput struct {
	_ struct{} `type:"structure"`

	IdentityType *string `locationName:"IdentityType" type:"string"`

	MaxItems *int64 `locationName:"MaxItems" type:"integer"`

	NextToken *string `locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListIdentitiesInput) String() string {
	return nifcloudutil.Prettify(s)
}

type ListIdentitiesOutput struct {
	_ struct{} `type:"structure"`

	Identities []string `locationName:"Identities" locationNameList:"member" type:"list"`

	NextToken *string `locationName:"NextToken" type:"string"`

	ResponseMetadata *ResponseMetadata `locationName:"ResponseMetadata" type:"structure"`
}

// String returns the string representation
func (s ListIdentitiesOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opListIdentities = "ListIdentities"

// ListIdentitiesRequest returns a request value for making API operation for
// NIFCLOUD ESS.
//
//    // Example sending a request using ListIdentitiesRequest.
//    req := client.ListIdentitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/ess/ListIdentities.htm
func (c *Client) ListIdentitiesRequest(input *ListIdentitiesInput) ListIdentitiesRequest {
	op := &aws.Operation{
		Name:       opListIdentities,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListIdentitiesInput{}
	}

	req := c.newRequest(op, input, &ListIdentitiesOutput{})

	return ListIdentitiesRequest{Request: req, Input: input, Copy: c.ListIdentitiesRequest}
}

// ListIdentitiesRequest is the request type for the
// ListIdentities API operation.
type ListIdentitiesRequest struct {
	*aws.Request
	Input *ListIdentitiesInput
	Copy  func(*ListIdentitiesInput) ListIdentitiesRequest
}

// Send marshals and sends the ListIdentities API request.
func (r ListIdentitiesRequest) Send(ctx context.Context) (*ListIdentitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListIdentitiesResponse{
		ListIdentitiesOutput: r.Request.Data.(*ListIdentitiesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListIdentitiesResponse is the response type for the
// ListIdentities API operation.
type ListIdentitiesResponse struct {
	*ListIdentitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListIdentities request.
func (r *ListIdentitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
