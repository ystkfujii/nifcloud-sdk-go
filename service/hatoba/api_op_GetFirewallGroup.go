// Code generated by smithy-go-codegen DO NOT EDIT.

package hatoba

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"github.com/nifcloud/nifcloud-sdk-go/service/hatoba/types"
	"strconv"
	"time"
)

func (c *Client) GetFirewallGroup(ctx context.Context, params *GetFirewallGroupInput, optFns ...func(*Options)) (*GetFirewallGroupOutput, error) {
	if params == nil {
		params = &GetFirewallGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFirewallGroup", params, optFns, c.addOperationGetFirewallGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFirewallGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFirewallGroupInput struct {

	// This member is required.
	FirewallGroupName *string

	noSmithyDocumentSerde
}

type GetFirewallGroupOutput struct {
	FirewallGroup *types.FirewallGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFirewallGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFirewallGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFirewallGroup{}, middleware.After)
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
	if err = addOpGetFirewallGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFirewallGroup(options.Region), middleware.Before); err != nil {
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

// GetFirewallGroupAPIClient is a client that implements the GetFirewallGroup
// operation.
type GetFirewallGroupAPIClient interface {
	GetFirewallGroup(context.Context, *GetFirewallGroupInput, ...func(*Options)) (*GetFirewallGroupOutput, error)
}

var _ GetFirewallGroupAPIClient = (*Client)(nil)

// FirewallRuleAuthorizedWaiterOptions are waiter options for
// FirewallRuleAuthorizedWaiter
type FirewallRuleAuthorizedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// FirewallRuleAuthorizedWaiter will use default minimum delay of 20 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, FirewallRuleAuthorizedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetFirewallGroupInput, *GetFirewallGroupOutput, error) (bool, error)
}

// FirewallRuleAuthorizedWaiter defines the waiters for FirewallRuleAuthorized
type FirewallRuleAuthorizedWaiter struct {
	client GetFirewallGroupAPIClient

	options FirewallRuleAuthorizedWaiterOptions
}

// NewFirewallRuleAuthorizedWaiter constructs a FirewallRuleAuthorizedWaiter.
func NewFirewallRuleAuthorizedWaiter(client GetFirewallGroupAPIClient, optFns ...func(*FirewallRuleAuthorizedWaiterOptions)) *FirewallRuleAuthorizedWaiter {
	options := FirewallRuleAuthorizedWaiterOptions{}
	options.MinDelay = 20 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = firewallRuleAuthorizedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &FirewallRuleAuthorizedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for FirewallRuleAuthorized waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *FirewallRuleAuthorizedWaiter) Wait(ctx context.Context, params *GetFirewallGroupInput, maxWaitDur time.Duration, optFns ...func(*FirewallRuleAuthorizedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for FirewallRuleAuthorized waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *FirewallRuleAuthorizedWaiter) WaitForOutput(ctx context.Context, params *GetFirewallGroupInput, maxWaitDur time.Duration, optFns ...func(*FirewallRuleAuthorizedWaiterOptions)) (*GetFirewallGroupOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetFirewallGroup(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for FirewallRuleAuthorized waiter")
}

func firewallRuleAuthorizedStateRetryable(ctx context.Context, input *GetFirewallGroupInput, output *GetFirewallGroupOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("FirewallGroup.Rules[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "AUTHORIZED"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("length(FirewallGroup.Rules[]) == `0`", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "true"
		bv, err := strconv.ParseBool(expectedValue)
		if err != nil {
			return false, fmt.Errorf("error parsing boolean from string %w", err)
		}
		value, ok := pathValue.(bool)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected bool value got %T", pathValue)
		}

		if value == bv {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetFirewallGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "hatoba",
		OperationName: "GetFirewallGroup",
	}
}
