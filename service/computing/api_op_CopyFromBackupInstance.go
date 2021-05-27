// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type CopyFromBackupInstanceInput struct {
	_ struct{} `type:"structure"`

	AccountingType AccountingTypeOfCopyFromBackupInstanceRequest `locationName:"AccountingType" type:"string" enum:"true"`

	// BackupInstanceUniqueId is a required field
	BackupInstanceUniqueId *string `locationName:"BackupInstanceUniqueId" type:"string" required:"true"`

	Description *string `locationName:"Description" type:"string"`

	DisableApiTermination *bool `locationName:"DisableApiTermination" type:"boolean"`

	InstanceId *string `locationName:"InstanceId" type:"string"`

	InstanceType InstanceTypeOfCopyFromBackupInstanceRequest `locationName:"InstanceType" type:"string" enum:"true"`

	NetworkInterface []RequestNetworkInterfaceOfCopyFromBackupInstance `locationName:"NetworkInterface" type:"list"`

	SecurityGroup []string `locationName:"SecurityGroup" type:"list"`
}

// String returns the string representation
func (s CopyFromBackupInstanceInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopyFromBackupInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CopyFromBackupInstanceInput"}

	if s.BackupInstanceUniqueId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupInstanceUniqueId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CopyFromBackupInstanceOutput struct {
	_ struct{} `type:"structure"`

	GroupSet []GroupSet `locationName:"groupSet" locationNameList:"item" type:"list"`

	Instance *Instance `locationName:"instance" type:"structure"`

	OwnerId *string `locationName:"ownerId" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`

	ReservationId *string `locationName:"reservationId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s CopyFromBackupInstanceOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opCopyFromBackupInstance = "CopyFromBackupInstance"

// CopyFromBackupInstanceRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using CopyFromBackupInstanceRequest.
//    req := client.CopyFromBackupInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/CopyFromBackupInstance.htm
func (c *Client) CopyFromBackupInstanceRequest(input *CopyFromBackupInstanceInput) CopyFromBackupInstanceRequest {
	op := &aws.Operation{
		Name:       opCopyFromBackupInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &CopyFromBackupInstanceInput{}
	}

	req := c.newRequest(op, input, &CopyFromBackupInstanceOutput{})

	return CopyFromBackupInstanceRequest{Request: req, Input: input, Copy: c.CopyFromBackupInstanceRequest}
}

// CopyFromBackupInstanceRequest is the request type for the
// CopyFromBackupInstance API operation.
type CopyFromBackupInstanceRequest struct {
	*aws.Request
	Input *CopyFromBackupInstanceInput
	Copy  func(*CopyFromBackupInstanceInput) CopyFromBackupInstanceRequest
}

// Send marshals and sends the CopyFromBackupInstance API request.
func (r CopyFromBackupInstanceRequest) Send(ctx context.Context) (*CopyFromBackupInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CopyFromBackupInstanceResponse{
		CopyFromBackupInstanceOutput: r.Request.Data.(*CopyFromBackupInstanceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CopyFromBackupInstanceResponse is the response type for the
// CopyFromBackupInstance API operation.
type CopyFromBackupInstanceResponse struct {
	*CopyFromBackupInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CopyFromBackupInstance request.
func (r *CopyFromBackupInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
