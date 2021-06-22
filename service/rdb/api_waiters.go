// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilDBInstanceAvailable uses the rdb API operation
// DescribeDBInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBInstanceAvailable(ctx context.Context, input *DescribeDBInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBInstanceAvailable",
		MaxAttempts: 80,
		Delay:       aws.ConstantWaiterDelay(40 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "available",
			},
			{
				State:   aws.FailureWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBInstanceDeleted uses the rdb API operation
// DescribeDBInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBInstanceDeleted(ctx context.Context, input *DescribeDBInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBInstanceDeleted",
		MaxAttempts: 80,
		Delay:       aws.ConstantWaiterDelay(40 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "Client.InvalidParameterNotFound.DBInstance",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "length(DBInstances[]) > `0`",
				Expected: true,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBInstanceExists uses the rdb API operation
// DescribeDBInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBInstanceExists(ctx context.Context, input *DescribeDBInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBInstanceExists",
		MaxAttempts: 80,
		Delay:       aws.ConstantWaiterDelay(40 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "length(DBInstances[]) > `0`",
				Expected: true,
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "Client.InvalidParameterNotFound.DBInstance",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBInstanceFailed uses the rdb API operation
// DescribeDBInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBInstanceFailed(ctx context.Context, input *DescribeDBInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBInstanceFailed",
		MaxAttempts: 80,
		Delay:       aws.ConstantWaiterDelay(40 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBInstanceStorageFull uses the rdb API operation
// DescribeDBInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBInstanceStorageFull(ctx context.Context, input *DescribeDBInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBInstanceStorageFull",
		MaxAttempts: 80,
		Delay:       aws.ConstantWaiterDelay(40 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "DBInstances[].DBInstanceStatus",
				Expected: "storage-full",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBSecurityGroupDeleted uses the rdb API operation
// DescribeDBSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBSecurityGroupDeleted(ctx context.Context, input *DescribeDBSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBSecurityGroupDeleted",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "Client.InvalidParameterNotFound.DBSecurityGroup",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "length(DBSecurityGroups[]) > `0`",
				Expected: true,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBSecurityGroupEC2SecurityGroupsAuthFailed uses the rdb API operation
// DescribeDBSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBSecurityGroupEC2SecurityGroupsAuthFailed(ctx context.Context, input *DescribeDBSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBSecurityGroupEC2SecurityGroupsAuthFailed",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "DBSecurityGroups[].EC2SecurityGroups[].Status",
				Expected: "auth-failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBSecurityGroupEC2SecurityGroupsAuthorized uses the rdb API operation
// DescribeDBSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBSecurityGroupEC2SecurityGroupsAuthorized(ctx context.Context, input *DescribeDBSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBSecurityGroupEC2SecurityGroupsAuthorized",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "DBSecurityGroups[].EC2SecurityGroups[].Status",
				Expected: "authorized",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBSecurityGroupEC2SecurityGroupsEmptied uses the rdb API operation
// DescribeDBSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBSecurityGroupEC2SecurityGroupsEmptied(ctx context.Context, input *DescribeDBSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBSecurityGroupEC2SecurityGroupsEmptied",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "length(DBSecurityGroups[0].EC2SecurityGroups[]) == `0`",
				Expected: true,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBSecurityGroupEC2SecurityGroupsRevokeFailed uses the rdb API operation
// DescribeDBSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBSecurityGroupEC2SecurityGroupsRevokeFailed(ctx context.Context, input *DescribeDBSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBSecurityGroupEC2SecurityGroupsRevokeFailed",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "DBSecurityGroups[].EC2SecurityGroups[].Status",
				Expected: "revoke-failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBSecurityGroupExists uses the rdb API operation
// DescribeDBSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBSecurityGroupExists(ctx context.Context, input *DescribeDBSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBSecurityGroupExists",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "length(DBSecurityGroups[]) > `0`",
				Expected: true,
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "Client.InvalidParameterNotFound.DBSecurityGroup",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBSecurityGroupIPRangesAuthFailed uses the rdb API operation
// DescribeDBSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBSecurityGroupIPRangesAuthFailed(ctx context.Context, input *DescribeDBSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBSecurityGroupIPRangesAuthFailed",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "DBSecurityGroups[].IPRanges[].Status",
				Expected: "auth-failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBSecurityGroupIPRangesAuthorized uses the rdb API operation
// DescribeDBSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBSecurityGroupIPRangesAuthorized(ctx context.Context, input *DescribeDBSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBSecurityGroupIPRangesAuthorized",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "DBSecurityGroups[].IPRanges[].Status",
				Expected: "authorized",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBSecurityGroupIPRangesEmptied uses the rdb API operation
// DescribeDBSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBSecurityGroupIPRangesEmptied(ctx context.Context, input *DescribeDBSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBSecurityGroupIPRangesEmptied",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "length(DBSecurityGroups[0].IPRanges[]) == `0`",
				Expected: true,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilDBSecurityGroupIPRangesRevokeFailed uses the rdb API operation
// DescribeDBSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilDBSecurityGroupIPRangesRevokeFailed(ctx context.Context, input *DescribeDBSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilDBSecurityGroupIPRangesRevokeFailed",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "DBSecurityGroups[].IPRanges[].Status",
				Expected: "revoke-failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeDBSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeDBSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
