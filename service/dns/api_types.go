// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dns

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

var _ aws.Config
var _ = nifcloudutil.Prettify

type ChangeInfo struct {
	_ struct{} `type:"structure"`

	Id *string `locationName:"Id" type:"string"`

	Status *string `locationName:"Status" type:"string"`

	SubmittedAt *string `locationName:"SubmittedAt" type:"string"`
}

// String returns the string representation
func (s ChangeInfo) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ChangeInfo) MarshalFields(e protocol.FieldEncoder) error {
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.StringValue(v), metadata)
	}
	if s.Status != nil {
		v := *s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Status", protocol.StringValue(v), metadata)
	}
	if s.SubmittedAt != nil {
		v := *s.SubmittedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SubmittedAt", protocol.StringValue(v), metadata)
	}
	return nil
}

type Config struct {
	_ struct{} `type:"structure"`

	Comment *string `locationName:"Comment" type:"string"`
}

// String returns the string representation
func (s Config) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Config) MarshalFields(e protocol.FieldEncoder) error {
	if s.Comment != nil {
		v := *s.Comment

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Comment", protocol.StringValue(v), metadata)
	}
	return nil
}

type DelegationSet struct {
	_ struct{} `type:"structure"`

	NameServers []string `locationName:"NameServers" locationNameList:"NameServer" type:"list"`
}

// String returns the string representation
func (s DelegationSet) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DelegationSet) MarshalFields(e protocol.FieldEncoder) error {
	if s.NameServers != nil {
		v := s.NameServers

		metadata := protocol.Metadata{ListLocationName: "NameServer"}
		ls0 := e.List(protocol.BodyTarget, "NameServers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.StringValue(v1))
		}
		ls0.End()

	}
	return nil
}

type HostedZone struct {
	_ struct{} `type:"structure"`

	CallerReference *string `locationName:"CallerReference" type:"string"`

	Config *Config `locationName:"Config" type:"structure"`

	Id *string `locationName:"Id" type:"string"`

	Name *string `locationName:"Name" type:"string"`

	ResourceRecordSetCount *int64 `locationName:"ResourceRecordSetCount" type:"integer"`
}

// String returns the string representation
func (s HostedZone) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s HostedZone) MarshalFields(e protocol.FieldEncoder) error {
	if s.CallerReference != nil {
		v := *s.CallerReference

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CallerReference", protocol.StringValue(v), metadata)
	}
	if s.Config != nil {
		v := s.Config

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Config", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.StringValue(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.StringValue(v), metadata)
	}
	if s.ResourceRecordSetCount != nil {
		v := *s.ResourceRecordSetCount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ResourceRecordSetCount", protocol.Int64Value(v), metadata)
	}
	return nil
}

type HostedZones struct {
	_ struct{} `type:"structure"`

	CallerReference *string `locationName:"CallerReference" type:"string"`

	Config *Config `locationName:"Config" type:"structure"`

	Id *string `locationName:"Id" type:"string"`

	Name *string `locationName:"Name" type:"string"`

	ResourceRecordSetCount *int64 `locationName:"ResourceRecordSetCount" type:"integer"`
}

// String returns the string representation
func (s HostedZones) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s HostedZones) MarshalFields(e protocol.FieldEncoder) error {
	if s.CallerReference != nil {
		v := *s.CallerReference

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CallerReference", protocol.StringValue(v), metadata)
	}
	if s.Config != nil {
		v := s.Config

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Config", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.StringValue(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.StringValue(v), metadata)
	}
	if s.ResourceRecordSetCount != nil {
		v := *s.ResourceRecordSetCount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ResourceRecordSetCount", protocol.Int64Value(v), metadata)
	}
	return nil
}

type RequestChange struct {
	_ struct{} `type:"structure"`

	// Action is a required field
	Action *string `locationName:"Action" type:"string" required:"true"`

	RequestResourceRecordSet *RequestResourceRecordSet `locationName:"ResourceRecordSet" type:"structure"`
}

// String returns the string representation
func (s RequestChange) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RequestChange) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RequestChange"}

	if s.Action == nil {
		invalidParams.Add(aws.NewErrParamRequired("Action"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RequestChange) MarshalFields(e protocol.FieldEncoder) error {
	if s.Action != nil {
		v := *s.Action

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Action", protocol.StringValue(v), metadata)
	}
	if s.RequestResourceRecordSet != nil {
		v := s.RequestResourceRecordSet

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ResourceRecordSet", v, metadata)
	}
	return nil
}

type RequestChangeBatch struct {
	_ struct{} `type:"structure"`

	// RequestChanges is a required field
	RequestChanges *RequestChanges `locationName:"Changes" type:"structure" required:"true"`
}

// String returns the string representation
func (s RequestChangeBatch) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RequestChangeBatch) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RequestChangeBatch"}

	if s.RequestChanges == nil {
		invalidParams.Add(aws.NewErrParamRequired("RequestChanges"))
	}
	if s.RequestChanges != nil {
		if err := s.RequestChanges.Validate(); err != nil {
			invalidParams.AddNested("RequestChanges", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RequestChangeBatch) MarshalFields(e protocol.FieldEncoder) error {
	if s.RequestChanges != nil {
		v := s.RequestChanges

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Changes", v, metadata)
	}
	return nil
}

type RequestChanges struct {
	_ struct{} `type:"structure"`

	// RequestChange is a required field
	RequestChange *RequestChange `locationName:"Change" type:"structure" required:"true"`
}

// String returns the string representation
func (s RequestChanges) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RequestChanges) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RequestChanges"}

	if s.RequestChange == nil {
		invalidParams.Add(aws.NewErrParamRequired("RequestChange"))
	}
	if s.RequestChange != nil {
		if err := s.RequestChange.Validate(); err != nil {
			invalidParams.AddNested("RequestChange", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RequestChanges) MarshalFields(e protocol.FieldEncoder) error {
	if s.RequestChange != nil {
		v := s.RequestChange

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Change", v, metadata)
	}
	return nil
}

type RequestHostedZoneConfig struct {
	_ struct{} `type:"structure"`

	Comment *string `locationName:"Comment" type:"string"`
}

// String returns the string representation
func (s RequestHostedZoneConfig) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RequestHostedZoneConfig) MarshalFields(e protocol.FieldEncoder) error {
	if s.Comment != nil {
		v := *s.Comment

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Comment", protocol.StringValue(v), metadata)
	}
	return nil
}

type RequestResourceRecord struct {
	_ struct{} `type:"structure"`

	Value *string `locationName:"Value" type:"string"`
}

// String returns the string representation
func (s RequestResourceRecord) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RequestResourceRecord) MarshalFields(e protocol.FieldEncoder) error {
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Value", protocol.StringValue(v), metadata)
	}
	return nil
}

type RequestResourceRecordSet struct {
	_ struct{} `type:"structure"`

	Failover *string `locationName:"Failover" type:"string"`

	Name *string `locationName:"Name" type:"string"`

	Region *string `locationName:"Region" type:"string"`

	RequestResourceRecords *RequestResourceRecords `locationName:"ResourceRecords" type:"structure"`

	RequestXniftyHealthCheckConfig *RequestXniftyHealthCheckConfig `locationName:"XniftyHealthCheckConfig" type:"structure"`

	SetIdentifier *string `locationName:"SetIdentifier" type:"string"`

	TTL *int64 `locationName:"TTL" type:"integer"`

	Type *string `locationName:"Type" type:"string"`

	Weight *int64 `locationName:"Weight" type:"integer"`

	XniftyComment *string `locationName:"XniftyComment" type:"string"`

	XniftyDefaultHost *string `locationName:"XniftyDefaultHost" type:"string"`
}

// String returns the string representation
func (s RequestResourceRecordSet) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RequestResourceRecordSet) MarshalFields(e protocol.FieldEncoder) error {
	if s.Failover != nil {
		v := *s.Failover

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Failover", protocol.StringValue(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.StringValue(v), metadata)
	}
	if s.Region != nil {
		v := *s.Region

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Region", protocol.StringValue(v), metadata)
	}
	if s.RequestResourceRecords != nil {
		v := s.RequestResourceRecords

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ResourceRecords", v, metadata)
	}
	if s.RequestXniftyHealthCheckConfig != nil {
		v := s.RequestXniftyHealthCheckConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "XniftyHealthCheckConfig", v, metadata)
	}
	if s.SetIdentifier != nil {
		v := *s.SetIdentifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SetIdentifier", protocol.StringValue(v), metadata)
	}
	if s.TTL != nil {
		v := *s.TTL

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TTL", protocol.Int64Value(v), metadata)
	}
	if s.Type != nil {
		v := *s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Type", protocol.StringValue(v), metadata)
	}
	if s.Weight != nil {
		v := *s.Weight

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Weight", protocol.Int64Value(v), metadata)
	}
	if s.XniftyComment != nil {
		v := *s.XniftyComment

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "XniftyComment", protocol.StringValue(v), metadata)
	}
	if s.XniftyDefaultHost != nil {
		v := *s.XniftyDefaultHost

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "XniftyDefaultHost", protocol.StringValue(v), metadata)
	}
	return nil
}

type RequestResourceRecords struct {
	_ struct{} `type:"structure"`

	RequestResourceRecord *RequestResourceRecord `locationName:"ResourceRecord" type:"structure"`
}

// String returns the string representation
func (s RequestResourceRecords) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RequestResourceRecords) MarshalFields(e protocol.FieldEncoder) error {
	if s.RequestResourceRecord != nil {
		v := s.RequestResourceRecord

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ResourceRecord", v, metadata)
	}
	return nil
}

type RequestXniftyHealthCheckConfig struct {
	_ struct{} `type:"structure"`

	FullyQualifiedDomainName *string `locationName:"FullyQualifiedDomainName" type:"string"`

	IPAddress *string `locationName:"IPAddress" type:"string"`

	Port *int64 `locationName:"Port" type:"integer"`

	Protocol *string `locationName:"Protocol" type:"string"`

	ResourcePath *string `locationName:"ResourcePath" type:"string"`
}

// String returns the string representation
func (s RequestXniftyHealthCheckConfig) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RequestXniftyHealthCheckConfig) MarshalFields(e protocol.FieldEncoder) error {
	if s.FullyQualifiedDomainName != nil {
		v := *s.FullyQualifiedDomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FullyQualifiedDomainName", protocol.StringValue(v), metadata)
	}
	if s.IPAddress != nil {
		v := *s.IPAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IPAddress", protocol.StringValue(v), metadata)
	}
	if s.Port != nil {
		v := *s.Port

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Port", protocol.Int64Value(v), metadata)
	}
	if s.Protocol != nil {
		v := *s.Protocol

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Protocol", protocol.StringValue(v), metadata)
	}
	if s.ResourcePath != nil {
		v := *s.ResourcePath

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ResourcePath", protocol.StringValue(v), metadata)
	}
	return nil
}

type ResourceRecordSets struct {
	_ struct{} `type:"structure"`

	Failover *string `locationName:"Failover" type:"string"`

	Name *string `locationName:"Name" type:"string"`

	Region *string `locationName:"Region" type:"string"`

	ResourceRecords []ResourceRecords `locationName:"ResourceRecords" locationNameList:"ResourceRecord" type:"list"`

	SetIdentifier *string `locationName:"SetIdentifier" type:"string"`

	TTL *int64 `locationName:"TTL" type:"integer"`

	Type *string `locationName:"Type" type:"string"`

	Weight *int64 `locationName:"Weight" type:"integer"`

	XniftyComment *string `locationName:"XniftyComment" type:"string"`

	XniftyDefaultHost *string `locationName:"XniftyDefaultHost" type:"string"`

	XniftyHealthCheckConfig *XniftyHealthCheckConfig `locationName:"XniftyHealthCheckConfig" type:"structure"`
}

// String returns the string representation
func (s ResourceRecordSets) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ResourceRecordSets) MarshalFields(e protocol.FieldEncoder) error {
	if s.Failover != nil {
		v := *s.Failover

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Failover", protocol.StringValue(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.StringValue(v), metadata)
	}
	if s.Region != nil {
		v := *s.Region

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Region", protocol.StringValue(v), metadata)
	}
	if s.ResourceRecords != nil {
		v := s.ResourceRecords

		metadata := protocol.Metadata{ListLocationName: "ResourceRecord"}
		ls0 := e.List(protocol.BodyTarget, "ResourceRecords", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.SetIdentifier != nil {
		v := *s.SetIdentifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SetIdentifier", protocol.StringValue(v), metadata)
	}
	if s.TTL != nil {
		v := *s.TTL

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TTL", protocol.Int64Value(v), metadata)
	}
	if s.Type != nil {
		v := *s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Type", protocol.StringValue(v), metadata)
	}
	if s.Weight != nil {
		v := *s.Weight

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Weight", protocol.Int64Value(v), metadata)
	}
	if s.XniftyComment != nil {
		v := *s.XniftyComment

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "XniftyComment", protocol.StringValue(v), metadata)
	}
	if s.XniftyDefaultHost != nil {
		v := *s.XniftyDefaultHost

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "XniftyDefaultHost", protocol.StringValue(v), metadata)
	}
	if s.XniftyHealthCheckConfig != nil {
		v := s.XniftyHealthCheckConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "XniftyHealthCheckConfig", v, metadata)
	}
	return nil
}

type ResourceRecords struct {
	_ struct{} `type:"structure"`

	Value *string `locationName:"Value" type:"string"`
}

// String returns the string representation
func (s ResourceRecords) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ResourceRecords) MarshalFields(e protocol.FieldEncoder) error {
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Value", protocol.StringValue(v), metadata)
	}
	return nil
}

type XniftyHealthCheckConfig struct {
	_ struct{} `type:"structure"`

	FullyQualifiedDomainName *string `locationName:"FullyQualifiedDomainName" type:"string"`

	IPAddress *string `locationName:"IPAddress" type:"string"`

	Port *int64 `locationName:"Port" type:"integer"`

	Protocol *string `locationName:"Protocol" type:"string"`

	ResourcePath *string `locationName:"ResourcePath" type:"string"`
}

// String returns the string representation
func (s XniftyHealthCheckConfig) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s XniftyHealthCheckConfig) MarshalFields(e protocol.FieldEncoder) error {
	if s.FullyQualifiedDomainName != nil {
		v := *s.FullyQualifiedDomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FullyQualifiedDomainName", protocol.StringValue(v), metadata)
	}
	if s.IPAddress != nil {
		v := *s.IPAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IPAddress", protocol.StringValue(v), metadata)
	}
	if s.Port != nil {
		v := *s.Port

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Port", protocol.Int64Value(v), metadata)
	}
	if s.Protocol != nil {
		v := *s.Protocol

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Protocol", protocol.StringValue(v), metadata)
	}
	if s.ResourcePath != nil {
		v := *s.ResourcePath

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ResourcePath", protocol.StringValue(v), metadata)
	}
	return nil
}
