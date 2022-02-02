// Code generated by smithy-go-codegen DO NOT EDIT.

package endpoints

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/smithy-go/logging"
	endpoints "github.com/nifcloud/nifcloud-sdk-go/internal/endpoints/v2"
	"regexp"
)

// Options is the endpoint resolver configuration options
type Options struct {
	// Logger is a logging implementation that log events should be sent to.
	Logger logging.Logger

	// LogDeprecated indicates that deprecated endpoints should be logged to the
	// provided logger.
	LogDeprecated bool

	// ResolvedRegion is used to override the region to be resolved, rather then the
	// using the value passed to the ResolveEndpoint method. This value is used by the
	// SDK to translate regions like fips-us-east-1 or us-east-1-fips to an alternative
	// name. You must not set this value directly in your application.
	ResolvedRegion string

	// DisableHTTPS informs the resolver to return an endpoint that does not use the
	// HTTPS scheme.
	DisableHTTPS bool

	// UseDualStackEndpoint specifies the resolver must resolve a dual-stack endpoint.
	UseDualStackEndpoint aws.DualStackEndpointState

	// UseFIPSEndpoint specifies the resolver must resolve a FIPS endpoint.
	UseFIPSEndpoint aws.FIPSEndpointState
}

func (o Options) GetResolvedRegion() string {
	return o.ResolvedRegion
}

func (o Options) GetDisableHTTPS() bool {
	return o.DisableHTTPS
}

func (o Options) GetUseDualStackEndpoint() aws.DualStackEndpointState {
	return o.UseDualStackEndpoint
}

func (o Options) GetUseFIPSEndpoint() aws.FIPSEndpointState {
	return o.UseFIPSEndpoint
}

func transformToSharedOptions(options Options) endpoints.Options {
	return endpoints.Options{
		Logger:               options.Logger,
		LogDeprecated:        options.LogDeprecated,
		ResolvedRegion:       options.ResolvedRegion,
		DisableHTTPS:         options.DisableHTTPS,
		UseDualStackEndpoint: options.UseDualStackEndpoint,
		UseFIPSEndpoint:      options.UseFIPSEndpoint,
	}
}

// Resolver ess endpoint resolver
type Resolver struct {
	partitions endpoints.Partitions
}

// ResolveEndpoint resolves the service endpoint for the given region and options
func (r *Resolver) ResolveEndpoint(region string, options Options) (endpoint aws.Endpoint, err error) {
	if len(region) == 0 {
		return endpoint, &aws.MissingRegionError{}
	}

	opt := transformToSharedOptions(options)
	return r.partitions.ResolveEndpoint(region, opt)
}

// New returns a new Resolver
func New() *Resolver {
	return &Resolver{
		partitions: defaultPartitions,
	}
}

var partitionRegexp = struct {
	Nifcloud *regexp.Regexp
}{

	Nifcloud: regexp.MustCompile("^(us|jp)\\-\\w+\\-\\d+$"),
}

var defaultPartitions = endpoints.Partitions{
	{
		ID: "nifcloud",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			endpoints.DefaultKey{
				Variant: 0,
			}: {
				Hostname:          "{region}.ess.api.nifcloud.com",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4", "v3", "v2"},
			},
		},
		RegionRegex:       partitionRegexp.Nifcloud,
		IsRegionalized:    false,
		PartitionEndpoint: "aws-global",
		Endpoints: endpoints.Endpoints{
			endpoints.EndpointKey{
				Region: "aws-global",
			}: endpoints.Endpoint{
				Hostname: "ess.api.nifcloud.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "east-1",
				},
			},
		},
	},
}
