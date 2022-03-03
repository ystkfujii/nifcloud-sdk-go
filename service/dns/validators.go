// Code generated by smithy-go-codegen DO NOT EDIT.


package dns

import (
	"context"
	"fmt"
	"github.com/aws/smithy-go/middleware"
	smithy "github.com/aws/smithy-go"
	"github.com/nifcloud/nifcloud-sdk-go/service/dns/types"
)

type validateOpChangeResourceRecordSets struct {
}

func (*validateOpChangeResourceRecordSets) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpChangeResourceRecordSets) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ChangeResourceRecordSetsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpChangeResourceRecordSetsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateHostedZone struct {
}

func (*validateOpCreateHostedZone) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateHostedZone) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateHostedZoneInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateHostedZoneInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteHostedZone struct {
}

func (*validateOpDeleteHostedZone) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteHostedZone) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteHostedZoneInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteHostedZoneInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetChange struct {
}

func (*validateOpGetChange) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetChange) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetChangeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetChangeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetHostedZone struct {
}

func (*validateOpGetHostedZone) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetHostedZone) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetHostedZoneInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetHostedZoneInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListResourceRecordSets struct {
}

func (*validateOpListResourceRecordSets) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListResourceRecordSets) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListResourceRecordSetsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListResourceRecordSetsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpChangeResourceRecordSetsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpChangeResourceRecordSets{}, middleware.After)
}

func addOpCreateHostedZoneValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateHostedZone{}, middleware.After)
}

func addOpDeleteHostedZoneValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteHostedZone{}, middleware.After)
}

func addOpGetChangeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetChange{}, middleware.After)
}

func addOpGetHostedZoneValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetHostedZone{}, middleware.After)
}

func addOpListResourceRecordSetsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListResourceRecordSets{}, middleware.After)
}

func validateListOfRequestChanges(v []types.RequestChanges) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListOfRequestChanges"}
	for i := range v {
		if err := validateRequestChanges(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateRequestChange(v *types.RequestChange) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RequestChange"}
	if len(v.Action) == 0 {
	invalidParams.Add(smithy.NewErrParamRequired("Action"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateRequestChangeBatch(v *types.RequestChangeBatch) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RequestChangeBatch"}
	if v.ListOfRequestChanges == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ListOfRequestChanges"))
	} else if v.ListOfRequestChanges != nil {
		if err := validateListOfRequestChanges(v.ListOfRequestChanges); err != nil {
			invalidParams.AddNested("ListOfRequestChanges", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateRequestChanges(v *types.RequestChanges) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RequestChanges"}
	if v.RequestChange == nil {
	invalidParams.Add(smithy.NewErrParamRequired("RequestChange"))
	} else if v.RequestChange != nil {
		if err := validateRequestChange(v.RequestChange); err != nil {
			invalidParams.AddNested("RequestChange", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpChangeResourceRecordSetsInput(v *ChangeResourceRecordSetsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ChangeResourceRecordSetsInput"}
	if v.RequestChangeBatch == nil {
	invalidParams.Add(smithy.NewErrParamRequired("RequestChangeBatch"))
	} else if v.RequestChangeBatch != nil {
		if err := validateRequestChangeBatch(v.RequestChangeBatch); err != nil {
			invalidParams.AddNested("RequestChangeBatch", err.(smithy.InvalidParamsError))
		}
	}
	if v.ZoneID == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ZoneID"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpCreateHostedZoneInput(v *CreateHostedZoneInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateHostedZoneInput"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpDeleteHostedZoneInput(v *DeleteHostedZoneInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteHostedZoneInput"}
	if v.ZoneID == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ZoneID"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpGetChangeInput(v *GetChangeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetChangeInput"}
	if v.ChangeID == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ChangeID"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpGetHostedZoneInput(v *GetHostedZoneInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetHostedZoneInput"}
	if v.ZoneID == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ZoneID"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpListResourceRecordSetsInput(v *ListResourceRecordSetsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListResourceRecordSetsInput"}
	if v.ZoneID == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ZoneID"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}
