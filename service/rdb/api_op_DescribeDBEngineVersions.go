// Code generated by smithy-go-codegen DO NOT EDIT.

package rdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/nifcloud/nifcloud-sdk-go/service/rdb/types"
)

func (c *Client) DescribeDBEngineVersions(ctx context.Context, params *DescribeDBEngineVersionsInput, optFns ...func(*Options)) (*DescribeDBEngineVersionsOutput, error) {
	if params == nil {
		params = &DescribeDBEngineVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBEngineVersions", params, optFns, c.addOperationDescribeDBEngineVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBEngineVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBEngineVersionsInput struct {
	DBParameterGroupFamily types.DBParameterGroupFamilyOfDescribeDBEngineVersionsRequest

	DefaultOnly *bool

	Engine types.EngineOfDescribeDBEngineVersionsRequest

	EngineVersion *string

	IncludeAll *bool

	ListSupportedCharacterSets *bool

	Marker *string

	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeDBEngineVersionsOutput struct {
	DBEngineVersions []types.DBEngineVersions

	Marker *string

	ResponseMetadata *types.ResponseMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBEngineVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBEngineVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBEngineVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBEngineVersions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeDBEngineVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rdb",
		OperationName: "DescribeDBEngineVersions",
	}
}
