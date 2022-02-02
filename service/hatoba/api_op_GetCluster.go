// Code generated by smithy-go-codegen DO NOT EDIT.

package hatoba

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"github.com/nifcloud/nifcloud-sdk-go/service/hatoba/types"
	"time"
)

func (c *Client) GetCluster(ctx context.Context, params *GetClusterInput, optFns ...func(*Options)) (*GetClusterOutput, error) {
	if params == nil {
		params = &GetClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCluster", params, optFns, c.addOperationGetClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetClusterInput struct {

	// This member is required.
	ClusterName *string

	noSmithyDocumentSerde
}

type GetClusterOutput struct {
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCluster{}, middleware.After)
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
	if err = addOpGetClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCluster(options.Region), middleware.Before); err != nil {
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

// GetClusterAPIClient is a client that implements the GetCluster operation.
type GetClusterAPIClient interface {
	GetCluster(context.Context, *GetClusterInput, ...func(*Options)) (*GetClusterOutput, error)
}

var _ GetClusterAPIClient = (*Client)(nil)

// ClusterRunningWaiterOptions are waiter options for ClusterRunningWaiter
type ClusterRunningWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ClusterRunningWaiter will use default minimum delay of 60 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ClusterRunningWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *GetClusterInput, *GetClusterOutput, error) (bool, error)
}

// ClusterRunningWaiter defines the waiters for ClusterRunning
type ClusterRunningWaiter struct {
	client GetClusterAPIClient

	options ClusterRunningWaiterOptions
}

// NewClusterRunningWaiter constructs a ClusterRunningWaiter.
func NewClusterRunningWaiter(client GetClusterAPIClient, optFns ...func(*ClusterRunningWaiterOptions)) *ClusterRunningWaiter {
	options := ClusterRunningWaiterOptions{}
	options.MinDelay = 60 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = clusterRunningStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ClusterRunningWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ClusterRunning waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ClusterRunningWaiter) Wait(ctx context.Context, params *GetClusterInput, maxWaitDur time.Duration, optFns ...func(*ClusterRunningWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ClusterRunning waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *ClusterRunningWaiter) WaitForOutput(ctx context.Context, params *GetClusterInput, maxWaitDur time.Duration, optFns ...func(*ClusterRunningWaiterOptions)) (*GetClusterOutput, error) {
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

		out, err := w.client.GetCluster(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for ClusterRunning waiter")
}

func clusterRunningStateRetryable(ctx context.Context, input *GetClusterInput, output *GetClusterOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Cluster.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "RUNNING"
		value, ok := pathValue.(*string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
		}

		if string(*value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Cluster.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ERROR"
		value, ok := pathValue.(*string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
		}

		if string(*value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Cluster.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DEGRADED"
		value, ok := pathValue.(*string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
		}

		if string(*value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Cluster.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "STOPPING"
		value, ok := pathValue.(*string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
		}

		if string(*value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// ClusterDeletedWaiterOptions are waiter options for ClusterDeletedWaiter
type ClusterDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ClusterDeletedWaiter will use default minimum delay of 30 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ClusterDeletedWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *GetClusterInput, *GetClusterOutput, error) (bool, error)
}

// ClusterDeletedWaiter defines the waiters for ClusterDeleted
type ClusterDeletedWaiter struct {
	client GetClusterAPIClient

	options ClusterDeletedWaiterOptions
}

// NewClusterDeletedWaiter constructs a ClusterDeletedWaiter.
func NewClusterDeletedWaiter(client GetClusterAPIClient, optFns ...func(*ClusterDeletedWaiterOptions)) *ClusterDeletedWaiter {
	options := ClusterDeletedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = clusterDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ClusterDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ClusterDeleted waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ClusterDeletedWaiter) Wait(ctx context.Context, params *GetClusterInput, maxWaitDur time.Duration, optFns ...func(*ClusterDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ClusterDeleted waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *ClusterDeletedWaiter) WaitForOutput(ctx context.Context, params *GetClusterInput, maxWaitDur time.Duration, optFns ...func(*ClusterDeletedWaiterOptions)) (*GetClusterOutput, error) {
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

		out, err := w.client.GetCluster(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for ClusterDeleted waiter")
}

func clusterDeletedStateRetryable(ctx context.Context, input *GetClusterInput, output *GetClusterOutput, err error) (bool, error) {

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "Client.InvalidParameterNotFound.Cluster" == apiErr.ErrorCode() {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Cluster.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "RUNNING"
		value, ok := pathValue.(*string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
		}

		if string(*value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Cluster.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ERROR"
		value, ok := pathValue.(*string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
		}

		if string(*value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Cluster.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DEGRADED"
		value, ok := pathValue.(*string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
		}

		if string(*value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "hatoba",
		OperationName: "GetCluster",
	}
}
