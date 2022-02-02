// Code generated by smithy-go-codegen DO NOT EDIT.

package computing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/nifcloud/nifcloud-sdk-go/service/computing/types"
)

func (c *Client) NiftyRegisterRoutersWithSecurityGroup(ctx context.Context, params *NiftyRegisterRoutersWithSecurityGroupInput, optFns ...func(*Options)) (*NiftyRegisterRoutersWithSecurityGroupOutput, error) {
	if params == nil {
		params = &NiftyRegisterRoutersWithSecurityGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "NiftyRegisterRoutersWithSecurityGroup", params, optFns, c.addOperationNiftyRegisterRoutersWithSecurityGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*NiftyRegisterRoutersWithSecurityGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NiftyRegisterRoutersWithSecurityGroupInput struct {

	// This member is required.
	GroupName *string

	RouterSet []types.RequestRouterSet

	noSmithyDocumentSerde
}

type NiftyRegisterRoutersWithSecurityGroupOutput struct {
	RequestId *string

	RouterSet []types.RouterSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationNiftyRegisterRoutersWithSecurityGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpNiftyRegisterRoutersWithSecurityGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpNiftyRegisterRoutersWithSecurityGroup{}, middleware.After)
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
	if err = addOpNiftyRegisterRoutersWithSecurityGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opNiftyRegisterRoutersWithSecurityGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opNiftyRegisterRoutersWithSecurityGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "NiftyRegisterRoutersWithSecurityGroup",
	}
}
