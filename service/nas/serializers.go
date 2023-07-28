// Code generated by smithy-go-codegen DO NOT EDIT.

package nas

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/query"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/nifcloud/nifcloud-sdk-go/internal/serializers"
	"github.com/nifcloud/nifcloud-sdk-go/service/nas/types"
	"path"
	"reflect"
)

type awsAwsquery_serializeOpAuthorizeNASSecurityGroupIngress struct {
}

func (*awsAwsquery_serializeOpAuthorizeNASSecurityGroupIngress) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpAuthorizeNASSecurityGroupIngress) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*AuthorizeNASSecurityGroupIngressInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("AuthorizeNASSecurityGroupIngress")
	body.Key("Version").String("N2016-02-24")

	if err := awsAwsquery_serializeOpDocumentAuthorizeNASSecurityGroupIngressInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(bodyWriter.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpClearNASSession struct {
}

func (*awsAwsquery_serializeOpClearNASSession) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpClearNASSession) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ClearNASSessionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("ClearNASSession")
	body.Key("Version").String("N2016-02-24")

	if err := awsAwsquery_serializeOpDocumentClearNASSessionInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(bodyWriter.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpCreateNASInstance struct {
}

func (*awsAwsquery_serializeOpCreateNASInstance) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpCreateNASInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateNASInstanceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("CreateNASInstance")
	body.Key("Version").String("N2016-02-24")

	if err := awsAwsquery_serializeOpDocumentCreateNASInstanceInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(bodyWriter.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpCreateNASSecurityGroup struct {
}

func (*awsAwsquery_serializeOpCreateNASSecurityGroup) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpCreateNASSecurityGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateNASSecurityGroupInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("CreateNASSecurityGroup")
	body.Key("Version").String("N2016-02-24")

	if err := awsAwsquery_serializeOpDocumentCreateNASSecurityGroupInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(bodyWriter.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpDeleteNASInstance struct {
}

func (*awsAwsquery_serializeOpDeleteNASInstance) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpDeleteNASInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteNASInstanceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("DeleteNASInstance")
	body.Key("Version").String("N2016-02-24")

	if err := awsAwsquery_serializeOpDocumentDeleteNASInstanceInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(bodyWriter.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpDeleteNASSecurityGroup struct {
}

func (*awsAwsquery_serializeOpDeleteNASSecurityGroup) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpDeleteNASSecurityGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteNASSecurityGroupInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("DeleteNASSecurityGroup")
	body.Key("Version").String("N2016-02-24")

	if err := awsAwsquery_serializeOpDocumentDeleteNASSecurityGroupInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(bodyWriter.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpDescribeNASInstances struct {
}

func (*awsAwsquery_serializeOpDescribeNASInstances) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpDescribeNASInstances) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeNASInstancesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("DescribeNASInstances")
	body.Key("Version").String("N2016-02-24")

	if err := awsAwsquery_serializeOpDocumentDescribeNASInstancesInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(bodyWriter.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpDescribeNASSecurityGroups struct {
}

func (*awsAwsquery_serializeOpDescribeNASSecurityGroups) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpDescribeNASSecurityGroups) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeNASSecurityGroupsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("DescribeNASSecurityGroups")
	body.Key("Version").String("N2016-02-24")

	if err := awsAwsquery_serializeOpDocumentDescribeNASSecurityGroupsInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(bodyWriter.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpGetMetricStatistics struct {
}

func (*awsAwsquery_serializeOpGetMetricStatistics) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpGetMetricStatistics) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetMetricStatisticsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("GetMetricStatistics")
	body.Key("Version").String("N2016-02-24")

	if err := awsAwsquery_serializeOpDocumentGetMetricStatisticsInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(bodyWriter.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpModifyNASInstance struct {
}

func (*awsAwsquery_serializeOpModifyNASInstance) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpModifyNASInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ModifyNASInstanceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("ModifyNASInstance")
	body.Key("Version").String("N2016-02-24")

	if err := awsAwsquery_serializeOpDocumentModifyNASInstanceInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(bodyWriter.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpModifyNASSecurityGroup struct {
}

func (*awsAwsquery_serializeOpModifyNASSecurityGroup) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpModifyNASSecurityGroup) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ModifyNASSecurityGroupInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("ModifyNASSecurityGroup")
	body.Key("Version").String("N2016-02-24")

	if err := awsAwsquery_serializeOpDocumentModifyNASSecurityGroupInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(bodyWriter.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpRevokeNASSecurityGroupIngress struct {
}

func (*awsAwsquery_serializeOpRevokeNASSecurityGroupIngress) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpRevokeNASSecurityGroupIngress) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*RevokeNASSecurityGroupIngressInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("RevokeNASSecurityGroupIngress")
	body.Key("Version").String("N2016-02-24")

	if err := awsAwsquery_serializeOpDocumentRevokeNASSecurityGroupIngressInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(bodyWriter.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpUpgradeNASInstance struct {
}

func (*awsAwsquery_serializeOpUpgradeNASInstance) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpUpgradeNASInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpgradeNASInstanceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("UpgradeNASInstance")
	body.Key("Version").String("N2016-02-24")

	if err := awsAwsquery_serializeOpDocumentUpgradeNASInstanceInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(bodyWriter.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsquery_serializeDocumentListOfRequestDimensions(v []types.RequestDimensions, value query.Value) error {
	array := value.Array("member")

	for i := range v {
		av := array.Value()
		if err := awsAwsquery_serializeDocumentRequestDimensions(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsquery_serializeDocumentListOfRequestNASSecurityGroups(v []string, value query.Value) error {
	array := value.Array("member")

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsquery_serializeDocumentRequestDimensions(v *types.RequestDimensions, value query.Value) error {
	object := value.Object()
	_ = object

	if v.Name != nil {
		objectKey := object.Key("Name")
		objectKey.String(*v.Name)
	}

	if v.Value != nil {
		objectKey := object.Key("Value")
		objectKey.String(*v.Value)
	}

	return nil
}

func awsAwsquery_serializeOpDocumentAuthorizeNASSecurityGroupIngressInput(v *AuthorizeNASSecurityGroupIngressInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.CIDRIP != nil {
		objectKey := object.Key("CIDRIP")
		objectKey.String(*v.CIDRIP)
	}

	if v.NASSecurityGroupName != nil {
		objectKey := object.Key("NASSecurityGroupName")
		objectKey.String(*v.NASSecurityGroupName)
	}

	if v.SecurityGroupName != nil {
		objectKey := object.Key("SecurityGroupName")
		objectKey.String(*v.SecurityGroupName)
	}

	return nil
}

func awsAwsquery_serializeOpDocumentClearNASSessionInput(v *ClearNASSessionInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.NASInstanceIdentifier != nil {
		objectKey := object.Key("NASInstanceIdentifier")
		objectKey.String(*v.NASInstanceIdentifier)
	}

	if len(v.SessionClearType) > 0 {
		objectKey := object.Key("SessionClearType")
		objectKey.String(string(v.SessionClearType))
	}

	return nil
}

func awsAwsquery_serializeOpDocumentCreateNASInstanceInput(v *CreateNASInstanceInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.AllocatedStorage != nil {
		objectKey := object.Key("AllocatedStorage")
		objectKey.Integer(*v.AllocatedStorage)
	}

	if v.AvailabilityZone != nil {
		objectKey := object.Key("AvailabilityZone")
		objectKey.String(*v.AvailabilityZone)
	}

	if v.MasterPrivateAddress != nil {
		objectKey := object.Key("MasterPrivateAddress")
		objectKey.String(*v.MasterPrivateAddress)
	}

	if v.MasterUsername != nil {
		objectKey := object.Key("MasterUsername")
		objectKey.String(*v.MasterUsername)
	}

	if v.MasterUserPassword != nil {
		objectKey := object.Key("MasterUserPassword")
		objectKey.String(*v.MasterUserPassword)
	}

	if v.NASInstanceDescription != nil {
		objectKey := object.Key("NASInstanceDescription")
		objectKey.String(*v.NASInstanceDescription)
	}

	if v.NASInstanceIdentifier != nil {
		objectKey := object.Key("NASInstanceIdentifier")
		objectKey.String(*v.NASInstanceIdentifier)
	}

	if v.NASInstanceType != nil {
		objectKey := object.Key("NASInstanceType")
		objectKey.Integer(*v.NASInstanceType)
	}

	if v.NASSecurityGroups != nil {
		objectKey := object.Key("NASSecurityGroups")
		if err := awsAwsquery_serializeDocumentListOfRequestNASSecurityGroups(v.NASSecurityGroups, objectKey); err != nil {
			return err
		}
	}

	if v.NetworkId != nil {
		objectKey := object.Key("NetworkId")
		objectKey.String(*v.NetworkId)
	}

	if len(v.Protocol) > 0 {
		objectKey := object.Key("Protocol")
		objectKey.String(string(v.Protocol))
	}

	return nil
}

func awsAwsquery_serializeOpDocumentCreateNASSecurityGroupInput(v *CreateNASSecurityGroupInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.AvailabilityZone != nil {
		objectKey := object.Key("AvailabilityZone")
		objectKey.String(*v.AvailabilityZone)
	}

	if v.NASSecurityGroupDescription != nil {
		objectKey := object.Key("NASSecurityGroupDescription")
		objectKey.String(*v.NASSecurityGroupDescription)
	}

	if v.NASSecurityGroupName != nil {
		objectKey := object.Key("NASSecurityGroupName")
		objectKey.String(*v.NASSecurityGroupName)
	}

	return nil
}

func awsAwsquery_serializeOpDocumentDeleteNASInstanceInput(v *DeleteNASInstanceInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.NASInstanceIdentifier != nil {
		objectKey := object.Key("NASInstanceIdentifier")
		objectKey.String(*v.NASInstanceIdentifier)
	}

	return nil
}

func awsAwsquery_serializeOpDocumentDeleteNASSecurityGroupInput(v *DeleteNASSecurityGroupInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.NASSecurityGroupName != nil {
		objectKey := object.Key("NASSecurityGroupName")
		objectKey.String(*v.NASSecurityGroupName)
	}

	return nil
}

func awsAwsquery_serializeOpDocumentDescribeNASInstancesInput(v *DescribeNASInstancesInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.NASInstanceIdentifier != nil {
		objectKey := object.Key("NASInstanceIdentifier")
		objectKey.String(*v.NASInstanceIdentifier)
	}

	return nil
}

func awsAwsquery_serializeOpDocumentDescribeNASSecurityGroupsInput(v *DescribeNASSecurityGroupsInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.NASSecurityGroupName != nil {
		objectKey := object.Key("NASSecurityGroupName")
		objectKey.String(*v.NASSecurityGroupName)
	}

	return nil
}

func awsAwsquery_serializeOpDocumentGetMetricStatisticsInput(v *GetMetricStatisticsInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.Dimensions != nil {
		objectKey := object.Key("Dimensions")
		if err := awsAwsquery_serializeDocumentListOfRequestDimensions(v.Dimensions, objectKey); err != nil {
			return err
		}
	}

	if v.EndTime != nil {
		objectKey := object.Key("EndTime")
		objectKey.String(serializers.FormatGetMetricStatisticsDateTime(reflect.ValueOf(*v).Type().Name(), *v.EndTime))
	}

	if len(v.MetricName) > 0 {
		objectKey := object.Key("MetricName")
		objectKey.String(string(v.MetricName))
	}

	if v.StartTime != nil {
		objectKey := object.Key("StartTime")
		objectKey.String(serializers.FormatGetMetricStatisticsDateTime(reflect.ValueOf(*v).Type().Name(), *v.StartTime))
	}

	return nil
}

func awsAwsquery_serializeOpDocumentModifyNASInstanceInput(v *ModifyNASInstanceInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.AllocatedStorage != nil {
		objectKey := object.Key("AllocatedStorage")
		objectKey.Integer(*v.AllocatedStorage)
	}

	if v.AuthenticationType != nil {
		objectKey := object.Key("AuthenticationType")
		objectKey.Integer(*v.AuthenticationType)
	}

	if v.MasterPrivateAddress != nil {
		objectKey := object.Key("MasterPrivateAddress")
		objectKey.String(*v.MasterPrivateAddress)
	}

	if v.MasterUserPassword != nil {
		objectKey := object.Key("MasterUserPassword")
		objectKey.String(*v.MasterUserPassword)
	}

	if v.NASInstanceDescription != nil {
		objectKey := object.Key("NASInstanceDescription")
		objectKey.String(*v.NASInstanceDescription)
	}

	if v.NASInstanceIdentifier != nil {
		objectKey := object.Key("NASInstanceIdentifier")
		objectKey.String(*v.NASInstanceIdentifier)
	}

	if v.NASSecurityGroups != nil {
		objectKey := object.Key("NASSecurityGroups")
		if err := awsAwsquery_serializeDocumentListOfRequestNASSecurityGroups(v.NASSecurityGroups, objectKey); err != nil {
			return err
		}
	}

	if v.NetworkId != nil {
		objectKey := object.Key("NetworkId")
		objectKey.String(*v.NetworkId)
	}

	if v.NewNASInstanceIdentifier != nil {
		objectKey := object.Key("NewNASInstanceIdentifier")
		objectKey.String(*v.NewNASInstanceIdentifier)
	}

	if v.NoRootSquash != nil {
		objectKey := object.Key("NoRootSquash")
		objectKey.Boolean(*v.NoRootSquash)
	}

	return nil
}

func awsAwsquery_serializeOpDocumentModifyNASSecurityGroupInput(v *ModifyNASSecurityGroupInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.NASSecurityGroupDescription != nil {
		objectKey := object.Key("NASSecurityGroupDescription")
		objectKey.String(*v.NASSecurityGroupDescription)
	}

	if v.NASSecurityGroupName != nil {
		objectKey := object.Key("NASSecurityGroupName")
		objectKey.String(*v.NASSecurityGroupName)
	}

	if v.NewNASSecurityGroupName != nil {
		objectKey := object.Key("NewNASSecurityGroupName")
		objectKey.String(*v.NewNASSecurityGroupName)
	}

	return nil
}

func awsAwsquery_serializeOpDocumentRevokeNASSecurityGroupIngressInput(v *RevokeNASSecurityGroupIngressInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.CIDRIP != nil {
		objectKey := object.Key("CIDRIP")
		objectKey.String(*v.CIDRIP)
	}

	if v.NASSecurityGroupName != nil {
		objectKey := object.Key("NASSecurityGroupName")
		objectKey.String(*v.NASSecurityGroupName)
	}

	if v.SecurityGroupName != nil {
		objectKey := object.Key("SecurityGroupName")
		objectKey.String(*v.SecurityGroupName)
	}

	return nil
}

func awsAwsquery_serializeOpDocumentUpgradeNASInstanceInput(v *UpgradeNASInstanceInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.NASInstanceIdentifier != nil {
		objectKey := object.Key("NASInstanceIdentifier")
		objectKey.String(*v.NASInstanceIdentifier)
	}

	return nil
}
