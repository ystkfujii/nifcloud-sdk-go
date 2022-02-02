// Code generated by smithy-go-codegen DO NOT EDIT.

package computing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func (c *Client) NiftyDeleteDhcpIpAddressPool(ctx context.Context, params *NiftyDeleteDhcpIpAddressPoolInput, optFns ...func(*Options)) (*NiftyDeleteDhcpIpAddressPoolOutput, error) {
	if params == nil {
		params = &NiftyDeleteDhcpIpAddressPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "NiftyDeleteDhcpIpAddressPool", params, optFns, c.addOperationNiftyDeleteDhcpIpAddressPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*NiftyDeleteDhcpIpAddressPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NiftyDeleteDhcpIpAddressPoolInput struct {

	// This member is required.
	DhcpConfigId *string

	// This member is required.
	StartIpAddress *string

	// This member is required.
	StopIpAddress *string

	noSmithyDocumentSerde
}

type NiftyDeleteDhcpIpAddressPoolOutput struct {
	RequestId *string

	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationNiftyDeleteDhcpIpAddressPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpNiftyDeleteDhcpIpAddressPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpNiftyDeleteDhcpIpAddressPool{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV2Middleware(stack, options); err != nil {
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
	if err = addOpNiftyDeleteDhcpIpAddressPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opNiftyDeleteDhcpIpAddressPool(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opNiftyDeleteDhcpIpAddressPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "NiftyDeleteDhcpIpAddressPool",
	}
}
