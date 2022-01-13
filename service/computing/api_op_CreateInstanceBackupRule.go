// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type CreateInstanceBackupRuleInput struct {
	_ struct{} `type:"structure"`

	// BackupInstanceMaxCount is a required field
	BackupInstanceMaxCount *int64 `locationName:"BackupInstanceMaxCount" type:"integer" required:"true"`

	Description *string `locationName:"Description" type:"string"`

	InstanceBackupRuleName *string `locationName:"InstanceBackupRuleName" type:"string"`

	// InstanceUniqueId is a required field
	InstanceUniqueId []string `locationName:"InstanceUniqueId" type:"list" required:"true"`

	// TimeSlotId is a required field
	TimeSlotId TimeSlotIdOfCreateInstanceBackupRuleRequest `locationName:"TimeSlotId" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateInstanceBackupRuleInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateInstanceBackupRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateInstanceBackupRuleInput"}

	if s.BackupInstanceMaxCount == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupInstanceMaxCount"))
	}

	if s.InstanceUniqueId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceUniqueId"))
	}
	if len(s.TimeSlotId) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("TimeSlotId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateInstanceBackupRuleOutput struct {
	_ struct{} `type:"structure"`

	InstanceBackupRule *InstanceBackupRuleOfCreateInstanceBackupRule `locationName:"instanceBackupRule" type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s CreateInstanceBackupRuleOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opCreateInstanceBackupRule = "CreateInstanceBackupRule"

// CreateInstanceBackupRuleRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using CreateInstanceBackupRuleRequest.
//    req := client.CreateInstanceBackupRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/CreateInstanceBackupRule.htm
func (c *Client) CreateInstanceBackupRuleRequest(input *CreateInstanceBackupRuleInput) CreateInstanceBackupRuleRequest {
	op := &aws.Operation{
		Name:       opCreateInstanceBackupRule,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &CreateInstanceBackupRuleInput{}
	}

	req := c.newRequest(op, input, &CreateInstanceBackupRuleOutput{})

	return CreateInstanceBackupRuleRequest{Request: req, Input: input, Copy: c.CreateInstanceBackupRuleRequest}
}

// CreateInstanceBackupRuleRequest is the request type for the
// CreateInstanceBackupRule API operation.
type CreateInstanceBackupRuleRequest struct {
	*aws.Request
	Input *CreateInstanceBackupRuleInput
	Copy  func(*CreateInstanceBackupRuleInput) CreateInstanceBackupRuleRequest
}

// Send marshals and sends the CreateInstanceBackupRule API request.
func (r CreateInstanceBackupRuleRequest) Send(ctx context.Context) (*CreateInstanceBackupRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateInstanceBackupRuleResponse{
		CreateInstanceBackupRuleOutput: r.Request.Data.(*CreateInstanceBackupRuleOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateInstanceBackupRuleResponse is the response type for the
// CreateInstanceBackupRule API operation.
type CreateInstanceBackupRuleResponse struct {
	*CreateInstanceBackupRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateInstanceBackupRule request.
func (r *CreateInstanceBackupRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
