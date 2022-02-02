// Code generated by smithy-go-codegen DO NOT EDIT.

package computing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func (c *Client) NiftyDescribeCorporateInfoForCertificate(ctx context.Context, params *NiftyDescribeCorporateInfoForCertificateInput, optFns ...func(*Options)) (*NiftyDescribeCorporateInfoForCertificateOutput, error) {
	if params == nil {
		params = &NiftyDescribeCorporateInfoForCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "NiftyDescribeCorporateInfoForCertificate", params, optFns, c.addOperationNiftyDescribeCorporateInfoForCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*NiftyDescribeCorporateInfoForCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NiftyDescribeCorporateInfoForCertificateInput struct {
	noSmithyDocumentSerde
}

type NiftyDescribeCorporateInfoForCertificateOutput struct {
	AlphabetName1 *string

	AlphabetName2 *string

	City *string

	CorpGrade *string

	CorpName *string

	DivisionName *string

	EmailAddress *string

	KanaName1 *string

	KanaName2 *string

	Name1 *string

	Name2 *string

	PhoneNumber *string

	PostName *string

	Pref *string

	PresidentName1 *string

	PresidentName2 *string

	RequestId *string

	TdbCode *string

	Zip1 *string

	Zip2 *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationNiftyDescribeCorporateInfoForCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpNiftyDescribeCorporateInfoForCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpNiftyDescribeCorporateInfoForCertificate{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opNiftyDescribeCorporateInfoForCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opNiftyDescribeCorporateInfoForCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "NiftyDescribeCorporateInfoForCertificate",
	}
}
