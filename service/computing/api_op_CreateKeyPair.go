// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type CreateKeyPairInput struct {
	_ struct{} `type:"structure"`

	Description *string `locationName:"Description" type:"string"`

	// KeyName is a required field
	KeyName *string `locationName:"KeyName" type:"string" required:"true"`

	// Password is a required field
	Password *string `locationName:"Password" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateKeyPairInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateKeyPairInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateKeyPairInput"}

	if s.KeyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyName"))
	}

	if s.Password == nil {
		invalidParams.Add(aws.NewErrParamRequired("Password"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateKeyPairOutput struct {
	_ struct{} `type:"structure"`

	Description *string `locationName:"description" type:"string"`

	KeyFingerprint *string `locationName:"keyFingerprint" type:"string"`

	KeyMaterial *string `locationName:"keyMaterial" type:"string"`

	KeyName *string `locationName:"keyName" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s CreateKeyPairOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opCreateKeyPair = "CreateKeyPair"

// CreateKeyPairRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using CreateKeyPairRequest.
//    req := client.CreateKeyPairRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/CreateKeyPair.htm
func (c *Client) CreateKeyPairRequest(input *CreateKeyPairInput) CreateKeyPairRequest {
	op := &aws.Operation{
		Name:       opCreateKeyPair,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &CreateKeyPairInput{}
	}

	req := c.newRequest(op, input, &CreateKeyPairOutput{})

	return CreateKeyPairRequest{Request: req, Input: input, Copy: c.CreateKeyPairRequest}
}

// CreateKeyPairRequest is the request type for the
// CreateKeyPair API operation.
type CreateKeyPairRequest struct {
	*aws.Request
	Input *CreateKeyPairInput
	Copy  func(*CreateKeyPairInput) CreateKeyPairRequest
}

// Send marshals and sends the CreateKeyPair API request.
func (r CreateKeyPairRequest) Send(ctx context.Context) (*CreateKeyPairResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateKeyPairResponse{
		CreateKeyPairOutput: r.Request.Data.(*CreateKeyPairOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateKeyPairResponse is the response type for the
// CreateKeyPair API operation.
type CreateKeyPairResponse struct {
	*CreateKeyPairOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateKeyPair request.
func (r *CreateKeyPairResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
