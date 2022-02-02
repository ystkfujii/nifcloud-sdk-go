// Code generated by smithy-go-codegen DO NOT EDIT.

package computing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func (c *Client) NiftyDisassociateRouteTableFromElasticLoadBalancer(ctx context.Context, params *NiftyDisassociateRouteTableFromElasticLoadBalancerInput, optFns ...func(*Options)) (*NiftyDisassociateRouteTableFromElasticLoadBalancerOutput, error) {
	if params == nil {
		params = &NiftyDisassociateRouteTableFromElasticLoadBalancerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "NiftyDisassociateRouteTableFromElasticLoadBalancer", params, optFns, c.addOperationNiftyDisassociateRouteTableFromElasticLoadBalancerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*NiftyDisassociateRouteTableFromElasticLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NiftyDisassociateRouteTableFromElasticLoadBalancerInput struct {

	// This member is required.
	AssociationId *string

	noSmithyDocumentSerde
}

type NiftyDisassociateRouteTableFromElasticLoadBalancerOutput struct {
	RequestId *string

	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationNiftyDisassociateRouteTableFromElasticLoadBalancerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpNiftyDisassociateRouteTableFromElasticLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpNiftyDisassociateRouteTableFromElasticLoadBalancer{}, middleware.After)
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
	if err = addOpNiftyDisassociateRouteTableFromElasticLoadBalancerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opNiftyDisassociateRouteTableFromElasticLoadBalancer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opNiftyDisassociateRouteTableFromElasticLoadBalancer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "NiftyDisassociateRouteTableFromElasticLoadBalancer",
	}
}
