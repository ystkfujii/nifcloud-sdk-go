// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type NiftyDescribeElasticLoadBalancersInput struct {
	_ struct{} `type:"structure"`

	ElasticLoadBalancers *RequestElasticLoadBalancers `locationName:"ElasticLoadBalancers" type:"structure"`

	Filter []RequestFilterOfNiftyDescribeElasticLoadBalancers `locationName:"Filter" type:"list"`
}

// String returns the string representation
func (s NiftyDescribeElasticLoadBalancersInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyDescribeElasticLoadBalancersOutput struct {
	_ struct{} `type:"structure"`

	NiftyDescribeElasticLoadBalancersResult *NiftyDescribeElasticLoadBalancersResult `locationName:"NiftyDescribeElasticLoadBalancersResult" type:"structure"`

	ResponseMetadata *ResponseMetadata `locationName:"ResponseMetadata" type:"structure"`
}

// String returns the string representation
func (s NiftyDescribeElasticLoadBalancersOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyDescribeElasticLoadBalancers = "NiftyDescribeElasticLoadBalancers"

// NiftyDescribeElasticLoadBalancersRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyDescribeElasticLoadBalancersRequest.
//    req := client.NiftyDescribeElasticLoadBalancersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/NiftyDescribeElasticLoadBalancers.htm
func (c *Client) NiftyDescribeElasticLoadBalancersRequest(input *NiftyDescribeElasticLoadBalancersInput) NiftyDescribeElasticLoadBalancersRequest {
	op := &aws.Operation{
		Name:       opNiftyDescribeElasticLoadBalancers,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyDescribeElasticLoadBalancersInput{}
	}

	req := c.newRequest(op, input, &NiftyDescribeElasticLoadBalancersOutput{})

	return NiftyDescribeElasticLoadBalancersRequest{Request: req, Input: input, Copy: c.NiftyDescribeElasticLoadBalancersRequest}
}

// NiftyDescribeElasticLoadBalancersRequest is the request type for the
// NiftyDescribeElasticLoadBalancers API operation.
type NiftyDescribeElasticLoadBalancersRequest struct {
	*aws.Request
	Input *NiftyDescribeElasticLoadBalancersInput
	Copy  func(*NiftyDescribeElasticLoadBalancersInput) NiftyDescribeElasticLoadBalancersRequest
}

// Send marshals and sends the NiftyDescribeElasticLoadBalancers API request.
func (r NiftyDescribeElasticLoadBalancersRequest) Send(ctx context.Context) (*NiftyDescribeElasticLoadBalancersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyDescribeElasticLoadBalancersResponse{
		NiftyDescribeElasticLoadBalancersOutput: r.Request.Data.(*NiftyDescribeElasticLoadBalancersOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyDescribeElasticLoadBalancersResponse is the response type for the
// NiftyDescribeElasticLoadBalancers API operation.
type NiftyDescribeElasticLoadBalancersResponse struct {
	*NiftyDescribeElasticLoadBalancersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyDescribeElasticLoadBalancers request.
func (r *NiftyDescribeElasticLoadBalancersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
