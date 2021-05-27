// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package script

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type ExecuteScriptInput struct {
	_ struct{} `type:"structure"`

	Body *string `locationName:"Body" type:"string"`

	Header *string `locationName:"Header" type:"string"`

	// Method is a required field
	Method *string `locationName:"Method" type:"string" required:"true"`

	Query *string `locationName:"Query" type:"string"`

	// ScriptIdentifier is a required field
	ScriptIdentifier *string `locationName:"ScriptIdentifier" type:"string" required:"true"`
}

// String returns the string representation
func (s ExecuteScriptInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExecuteScriptInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ExecuteScriptInput"}

	if s.Method == nil {
		invalidParams.Add(aws.NewErrParamRequired("Method"))
	}

	if s.ScriptIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScriptIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ExecuteScriptOutput struct {
	_ struct{} `type:"structure"`

	Result *Result `type:"structure"`
}

// String returns the string representation
func (s ExecuteScriptOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opExecuteScript = "ExecuteScript"

// ExecuteScriptRequest returns a request value for making API operation for
// NIFCLOUD Script.
//
//    // Example sending a request using ExecuteScriptRequest.
//    req := client.ExecuteScriptRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/script/ExecuteScript.htm
func (c *Client) ExecuteScriptRequest(input *ExecuteScriptInput) ExecuteScriptRequest {
	op := &aws.Operation{
		Name:       opExecuteScript,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-09-01/",
	}

	if input == nil {
		input = &ExecuteScriptInput{}
	}

	req := c.newRequest(op, input, &ExecuteScriptOutput{})

	return ExecuteScriptRequest{Request: req, Input: input, Copy: c.ExecuteScriptRequest}
}

// ExecuteScriptRequest is the request type for the
// ExecuteScript API operation.
type ExecuteScriptRequest struct {
	*aws.Request
	Input *ExecuteScriptInput
	Copy  func(*ExecuteScriptInput) ExecuteScriptRequest
}

// Send marshals and sends the ExecuteScript API request.
func (r ExecuteScriptRequest) Send(ctx context.Context) (*ExecuteScriptResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ExecuteScriptResponse{
		ExecuteScriptOutput: r.Request.Data.(*ExecuteScriptOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ExecuteScriptResponse is the response type for the
// ExecuteScript API operation.
type ExecuteScriptResponse struct {
	*ExecuteScriptOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ExecuteScript request.
func (r *ExecuteScriptResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
