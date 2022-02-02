// Code generated by smithy-go-codegen DO NOT EDIT.

package nas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/nifcloud/nifcloud-sdk-go/service/nas/types"
)

func (c *Client) AuthorizeNASSecurityGroupIngress(ctx context.Context, params *AuthorizeNASSecurityGroupIngressInput, optFns ...func(*Options)) (*AuthorizeNASSecurityGroupIngressOutput, error) {
	if params == nil {
		params = &AuthorizeNASSecurityGroupIngressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AuthorizeNASSecurityGroupIngress", params, optFns, c.addOperationAuthorizeNASSecurityGroupIngressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AuthorizeNASSecurityGroupIngressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AuthorizeNASSecurityGroupIngressInput struct {

	// This member is required.
	NASSecurityGroupName *string

	CIDRIP *string

	SecurityGroupName *string

	noSmithyDocumentSerde
}

type AuthorizeNASSecurityGroupIngressOutput struct {
	NASSecurityGroup *types.NASSecurityGroup

	ResponseMetadata *types.ResponseMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAuthorizeNASSecurityGroupIngressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAuthorizeNASSecurityGroupIngress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAuthorizeNASSecurityGroupIngress{}, middleware.After)
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
	if err = addOpAuthorizeNASSecurityGroupIngressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAuthorizeNASSecurityGroupIngress(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAuthorizeNASSecurityGroupIngress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nas",
		OperationName: "AuthorizeNASSecurityGroupIngress",
	}
}
