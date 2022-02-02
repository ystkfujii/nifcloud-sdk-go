// Code generated by smithy-go-codegen DO NOT EDIT.

package computing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/nifcloud/nifcloud-sdk-go/service/computing/types"
)

func (c *Client) NiftyModifyElasticLoadBalancerAttributes(ctx context.Context, params *NiftyModifyElasticLoadBalancerAttributesInput, optFns ...func(*Options)) (*NiftyModifyElasticLoadBalancerAttributesOutput, error) {
	if params == nil {
		params = &NiftyModifyElasticLoadBalancerAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "NiftyModifyElasticLoadBalancerAttributes", params, optFns, c.addOperationNiftyModifyElasticLoadBalancerAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*NiftyModifyElasticLoadBalancerAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NiftyModifyElasticLoadBalancerAttributesInput struct {

	// This member is required.
	ElasticLoadBalancerPort *int32

	// This member is required.
	InstancePort *int32

	// This member is required.
	Protocol types.ProtocolOfNiftyModifyElasticLoadBalancerAttributesRequest

	ElasticLoadBalancerId *string

	ElasticLoadBalancerName *string

	LoadBalancerAttributes *types.RequestLoadBalancerAttributes

	noSmithyDocumentSerde
}

type NiftyModifyElasticLoadBalancerAttributesOutput struct {
	ResponseMetadata *types.ResponseMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationNiftyModifyElasticLoadBalancerAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpNiftyModifyElasticLoadBalancerAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpNiftyModifyElasticLoadBalancerAttributes{}, middleware.After)
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
	if err = addOpNiftyModifyElasticLoadBalancerAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opNiftyModifyElasticLoadBalancerAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opNiftyModifyElasticLoadBalancerAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "NiftyModifyElasticLoadBalancerAttributes",
	}
}
