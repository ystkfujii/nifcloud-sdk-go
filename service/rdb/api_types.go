// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

var _ aws.Config
var _ = nifcloudutil.Prettify

type AvailabilityZones struct {
	_ struct{} `type:"structure"`

	Name *string `locationName:"Name" type:"string"`

	NiftyStorageTypes []int64 `locationName:"NiftyStorageTypes" locationNameList:"NiftyStorageType" type:"list"`

	ProvisionedIopsCapable *bool `locationName:"ProvisionedIopsCapable" type:"boolean"`
}

// String returns the string representation
func (s AvailabilityZones) String() string {
	return nifcloudutil.Prettify(s)
}

type Certificates struct {
	_ struct{} `type:"structure"`

	CertificateIdentifier *string `locationName:"CertificateIdentifier" type:"string"`

	CertificateType *string `locationName:"CertificateType" type:"string"`

	Thumbprint *string `locationName:"Thumbprint" type:"string"`

	ValidFrom *time.Time `locationName:"ValidFrom" type:"timestamp"`

	ValidTill *time.Time `locationName:"ValidTill" type:"timestamp"`
}

// String returns the string representation
func (s Certificates) String() string {
	return nifcloudutil.Prettify(s)
}

type DBEngineVersions struct {
	_ struct{} `type:"structure"`

	DBEngineDescription *string `locationName:"DBEngineDescription" type:"string"`

	DBEngineVersionDescription *string `locationName:"DBEngineVersionDescription" type:"string"`

	DBParameterGroupFamily *string `locationName:"DBParameterGroupFamily" type:"string"`

	Engine *string `locationName:"Engine" type:"string"`

	EngineVersion *string `locationName:"EngineVersion" type:"string"`
}

// String returns the string representation
func (s DBEngineVersions) String() string {
	return nifcloudutil.Prettify(s)
}

type DBInstance struct {
	_ struct{} `type:"structure"`

	AccountingType *string `locationName:"AccountingType" type:"string"`

	AllocatedStorage *int64 `locationName:"AllocatedStorage" type:"integer"`

	AutoMinorVersionUpgrade *bool `locationName:"AutoMinorVersionUpgrade" type:"boolean"`

	AvailabilityZone *string `locationName:"AvailabilityZone" type:"string"`

	BackupRetentionPeriod *int64 `locationName:"BackupRetentionPeriod" type:"integer"`

	BinlogRetentionPeriod *int64 `locationName:"BinlogRetentionPeriod" type:"integer"`

	CACertificateIdentifier *string `locationName:"CACertificateIdentifier" type:"string"`

	DBInstanceClass *string `locationName:"DBInstanceClass" type:"string"`

	DBInstanceIdentifier *string `locationName:"DBInstanceIdentifier" type:"string"`

	DBInstanceStatus *string `locationName:"DBInstanceStatus" type:"string"`

	DBName *string `locationName:"DBName" type:"string"`

	DBParameterGroups []DBParameterGroups `locationName:"DBParameterGroups" locationNameList:"DBParameterGroup" type:"list"`

	DBSecurityGroups []DBSecurityGroups `locationName:"DBSecurityGroups" locationNameList:"DBSecurityGroup" type:"list"`

	Endpoint *Endpoint `locationName:"Endpoint" type:"structure"`

	Engine *string `locationName:"Engine" type:"string"`

	EngineVersion *string `locationName:"EngineVersion" type:"string"`

	ExternalReplicationInfo *ExternalReplicationInfo `locationName:"ExternalReplicationInfo" type:"structure"`

	InstanceCreateTime *time.Time `locationName:"InstanceCreateTime" type:"timestamp"`

	LatestRestorableTime *time.Time `locationName:"LatestRestorableTime" type:"timestamp"`

	LicenseModel *string `locationName:"LicenseModel" type:"string"`

	MasterUsername *string `locationName:"MasterUsername" type:"string"`

	MultiAZ *bool `locationName:"MultiAZ" type:"boolean"`

	NextMonthAccountingType *string `locationName:"NextMonthAccountingType" type:"string"`

	NiftyMasterPrivateAddress *string `locationName:"NiftyMasterPrivateAddress" type:"string"`

	NiftyMultiAZType *string `locationName:"NiftyMultiAZType" type:"string"`

	NiftyNetworkId *string `locationName:"NiftyNetworkId" type:"string"`

	NiftySlavePrivateAddress *string `locationName:"NiftySlavePrivateAddress" type:"string"`

	NiftyStorageType *int64 `locationName:"NiftyStorageType" type:"integer"`

	OptionGroupMemberships []OptionGroupMemberships `locationName:"OptionGroupMemberships" locationNameList:"OptionGroupMembership" type:"list"`

	PendingModifiedValues *PendingModifiedValues `locationName:"PendingModifiedValues" type:"structure"`

	PreferredBackupWindow *string `locationName:"PreferredBackupWindow" type:"string"`

	PreferredMaintenanceWindow *string `locationName:"PreferredMaintenanceWindow" type:"string"`

	PubliclyAccessible *bool `locationName:"PubliclyAccessible" type:"boolean"`

	ReadReplicaDBInstanceIdentifiers []string `locationName:"ReadReplicaDBInstanceIdentifiers" locationNameList:"ReadReplicaDBInstanceIdentifier" type:"list"`

	ReadReplicaSourceDBInstanceIdentifier *string `locationName:"ReadReplicaSourceDBInstanceIdentifier" type:"string"`

	SecondaryAvailabilityZone *string `locationName:"SecondaryAvailabilityZone" type:"string"`

	StatusInfos []StatusInfos `locationName:"StatusInfos" locationNameList:"DBInstanceStatusInfo" type:"list"`

	VpcSecurityGroups *string `locationName:"VpcSecurityGroups" type:"string"`
}

// String returns the string representation
func (s DBInstance) String() string {
	return nifcloudutil.Prettify(s)
}

type DBInstances struct {
	_ struct{} `type:"structure"`

	AccountingType *string `locationName:"AccountingType" type:"string"`

	AllocatedStorage *int64 `locationName:"AllocatedStorage" type:"integer"`

	AutoMinorVersionUpgrade *bool `locationName:"AutoMinorVersionUpgrade" type:"boolean"`

	AvailabilityZone *string `locationName:"AvailabilityZone" type:"string"`

	BackupRetentionPeriod *int64 `locationName:"BackupRetentionPeriod" type:"integer"`

	BinlogRetentionPeriod *int64 `locationName:"BinlogRetentionPeriod" type:"integer"`

	CACertificateIdentifier *string `locationName:"CACertificateIdentifier" type:"string"`

	DBInstanceClass *string `locationName:"DBInstanceClass" type:"string"`

	DBInstanceIdentifier *string `locationName:"DBInstanceIdentifier" type:"string"`

	DBInstanceStatus *string `locationName:"DBInstanceStatus" type:"string"`

	DBName *string `locationName:"DBName" type:"string"`

	DBParameterGroups []DBParameterGroups `locationName:"DBParameterGroups" locationNameList:"DBParameterGroup" type:"list"`

	DBSecurityGroups []DBSecurityGroups `locationName:"DBSecurityGroups" locationNameList:"DBSecurityGroup" type:"list"`

	Endpoint *Endpoint `locationName:"Endpoint" type:"structure"`

	Engine *string `locationName:"Engine" type:"string"`

	EngineVersion *string `locationName:"EngineVersion" type:"string"`

	ExternalReplicationInfo *ExternalReplicationInfo `locationName:"ExternalReplicationInfo" type:"structure"`

	InstanceCreateTime *time.Time `locationName:"InstanceCreateTime" type:"timestamp"`

	LatestRestorableTime *time.Time `locationName:"LatestRestorableTime" type:"timestamp"`

	LicenseModel *string `locationName:"LicenseModel" type:"string"`

	MasterUsername *string `locationName:"MasterUsername" type:"string"`

	MultiAZ *bool `locationName:"MultiAZ" type:"boolean"`

	NextMonthAccountingType *string `locationName:"NextMonthAccountingType" type:"string"`

	NiftyMasterPrivateAddress *string `locationName:"NiftyMasterPrivateAddress" type:"string"`

	NiftyMultiAZType *string `locationName:"NiftyMultiAZType" type:"string"`

	NiftyNetworkId *string `locationName:"NiftyNetworkId" type:"string"`

	NiftySlavePrivateAddress *string `locationName:"NiftySlavePrivateAddress" type:"string"`

	NiftyStorageType *int64 `locationName:"NiftyStorageType" type:"integer"`

	OptionGroupMemberships []OptionGroupMemberships `locationName:"OptionGroupMemberships" locationNameList:"OptionGroupMembership" type:"list"`

	PendingModifiedValues *PendingModifiedValues `locationName:"PendingModifiedValues" type:"structure"`

	PreferredBackupWindow *string `locationName:"PreferredBackupWindow" type:"string"`

	PreferredMaintenanceWindow *string `locationName:"PreferredMaintenanceWindow" type:"string"`

	PubliclyAccessible *bool `locationName:"PubliclyAccessible" type:"boolean"`

	ReadReplicaDBInstanceIdentifiers []string `locationName:"ReadReplicaDBInstanceIdentifiers" locationNameList:"ReadReplicaDBInstanceIdentifier" type:"list"`

	ReadReplicaSourceDBInstanceIdentifier *string `locationName:"ReadReplicaSourceDBInstanceIdentifier" type:"string"`

	SecondaryAvailabilityZone *string `locationName:"SecondaryAvailabilityZone" type:"string"`

	StatusInfos []StatusInfos `locationName:"StatusInfos" locationNameList:"DBInstanceStatusInfo" type:"list"`

	VpcSecurityGroups *string `locationName:"VpcSecurityGroups" type:"string"`
}

// String returns the string representation
func (s DBInstances) String() string {
	return nifcloudutil.Prettify(s)
}

type DBParameterGroup struct {
	_ struct{} `type:"structure"`

	DBParameterGroupFamily *string `locationName:"DBParameterGroupFamily" type:"string"`

	DBParameterGroupName *string `locationName:"DBParameterGroupName" type:"string"`

	Description *string `locationName:"Description" type:"string"`
}

// String returns the string representation
func (s DBParameterGroup) String() string {
	return nifcloudutil.Prettify(s)
}

type DBParameterGroups struct {
	_ struct{} `type:"structure"`

	DBParameterGroupName *string `locationName:"DBParameterGroupName" type:"string"`

	ParameterApplyStatus *string `locationName:"ParameterApplyStatus" type:"string"`
}

// String returns the string representation
func (s DBParameterGroups) String() string {
	return nifcloudutil.Prettify(s)
}

type DBParameterGroupsOfDescribeDBParameterGroups struct {
	_ struct{} `type:"structure"`

	DBParameterGroupFamily *string `locationName:"DBParameterGroupFamily" type:"string"`

	DBParameterGroupName *string `locationName:"DBParameterGroupName" type:"string"`

	Description *string `locationName:"Description" type:"string"`
}

// String returns the string representation
func (s DBParameterGroupsOfDescribeDBParameterGroups) String() string {
	return nifcloudutil.Prettify(s)
}

type DBSecurityGroup struct {
	_ struct{} `type:"structure"`

	DBSecurityGroupDescription *string `locationName:"DBSecurityGroupDescription" type:"string"`

	DBSecurityGroupName *string `locationName:"DBSecurityGroupName" type:"string"`

	EC2SecurityGroups []EC2SecurityGroups `locationName:"EC2SecurityGroups" locationNameList:"EC2SecurityGroup" type:"list"`

	IPRanges []IPRanges `locationName:"IPRanges" locationNameList:"IPRange" type:"list"`

	NiftyAvailabilityZone *string `locationName:"NiftyAvailabilityZone" type:"string"`

	OwnerId *string `locationName:"OwnerId" type:"string"`
}

// String returns the string representation
func (s DBSecurityGroup) String() string {
	return nifcloudutil.Prettify(s)
}

type DBSecurityGroups struct {
	_ struct{} `type:"structure"`

	DBSecurityGroupName *string `locationName:"DBSecurityGroupName" type:"string"`

	Status *string `locationName:"Status" type:"string"`
}

// String returns the string representation
func (s DBSecurityGroups) String() string {
	return nifcloudutil.Prettify(s)
}

type DBSecurityGroupsOfDescribeDBSecurityGroups struct {
	_ struct{} `type:"structure"`

	DBSecurityGroupDescription *string `locationName:"DBSecurityGroupDescription" type:"string"`

	DBSecurityGroupName *string `locationName:"DBSecurityGroupName" type:"string"`

	EC2SecurityGroups []EC2SecurityGroups `locationName:"EC2SecurityGroups" locationNameList:"EC2SecurityGroup" type:"list"`

	IPRanges []IPRanges `locationName:"IPRanges" locationNameList:"IPRange" type:"list"`

	NiftyAvailabilityZone *string `locationName:"NiftyAvailabilityZone" type:"string"`

	OwnerId *string `locationName:"OwnerId" type:"string"`
}

// String returns the string representation
func (s DBSecurityGroupsOfDescribeDBSecurityGroups) String() string {
	return nifcloudutil.Prettify(s)
}

type DBSnapshot struct {
	_ struct{} `type:"structure"`

	AllocatedStorage *int64 `locationName:"AllocatedStorage" type:"integer"`

	AvailabilityZone *string `locationName:"AvailabilityZone" type:"string"`

	DBInstanceIdentifier *string `locationName:"DBInstanceIdentifier" type:"string"`

	DBSnapshotIdentifier *string `locationName:"DBSnapshotIdentifier" type:"string"`

	Engine *string `locationName:"Engine" type:"string"`

	EngineVersion *string `locationName:"EngineVersion" type:"string"`

	InstanceCreateTime *time.Time `locationName:"InstanceCreateTime" type:"timestamp"`

	LicenseModel *string `locationName:"LicenseModel" type:"string"`

	MasterUsername *string `locationName:"MasterUsername" type:"string"`

	OptionGroupName *string `locationName:"OptionGroupName" type:"string"`

	Port *int64 `locationName:"Port" type:"integer"`

	SnapshotCreateTime *time.Time `locationName:"SnapshotCreateTime" type:"timestamp"`

	SnapshotType *string `locationName:"SnapshotType" type:"string"`

	Status *string `locationName:"Status" type:"string"`
}

// String returns the string representation
func (s DBSnapshot) String() string {
	return nifcloudutil.Prettify(s)
}

type DBSnapshots struct {
	_ struct{} `type:"structure"`

	AllocatedStorage *int64 `locationName:"AllocatedStorage" type:"integer"`

	AvailabilityZone *string `locationName:"AvailabilityZone" type:"string"`

	DBInstanceIdentifier *string `locationName:"DBInstanceIdentifier" type:"string"`

	DBSnapshotIdentifier *string `locationName:"DBSnapshotIdentifier" type:"string"`

	Engine *string `locationName:"Engine" type:"string"`

	EngineVersion *string `locationName:"EngineVersion" type:"string"`

	InstanceCreateTime *time.Time `locationName:"InstanceCreateTime" type:"timestamp"`

	LicenseModel *string `locationName:"LicenseModel" type:"string"`

	MasterUsername *string `locationName:"MasterUsername" type:"string"`

	OptionGroupName *string `locationName:"OptionGroupName" type:"string"`

	Port *int64 `locationName:"Port" type:"integer"`

	SnapshotCreateTime *time.Time `locationName:"SnapshotCreateTime" type:"timestamp"`

	SnapshotType *string `locationName:"SnapshotType" type:"string"`

	Status *string `locationName:"Status" type:"string"`
}

// String returns the string representation
func (s DBSnapshots) String() string {
	return nifcloudutil.Prettify(s)
}

type Datapoints struct {
	_ struct{} `type:"structure"`

	NiftyTargetName *string `locationName:"NiftyTargetName" type:"string"`

	SampleCount *int64 `locationName:"SampleCount" type:"integer"`

	Sum *float64 `locationName:"Sum" type:"double"`

	Timestamp *time.Time `locationName:"Timestamp" type:"timestamp"`
}

// String returns the string representation
func (s Datapoints) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeDBLogFiles struct {
	_ struct{} `type:"structure"`

	LastWritten *int64 `locationName:"LastWritten" type:"long"`

	LogFileName *string `locationName:"LogFileName" type:"string"`

	Size *int64 `locationName:"Size" type:"long"`
}

// String returns the string representation
func (s DescribeDBLogFiles) String() string {
	return nifcloudutil.Prettify(s)
}

type EC2SecurityGroups struct {
	_ struct{} `type:"structure"`

	EC2SecurityGroupName *string `locationName:"EC2SecurityGroupName" type:"string"`

	EC2SecurityGroupOwnerId *string `locationName:"EC2SecurityGroupOwnerId" type:"string"`

	Status *string `locationName:"Status" type:"string"`
}

// String returns the string representation
func (s EC2SecurityGroups) String() string {
	return nifcloudutil.Prettify(s)
}

type Endpoint struct {
	_ struct{} `type:"structure"`

	Address *string `locationName:"Address" type:"string"`

	NiftyPrivateAddress *string `locationName:"NiftyPrivateAddress" type:"string"`

	Port *int64 `locationName:"Port" type:"integer"`
}

// String returns the string representation
func (s Endpoint) String() string {
	return nifcloudutil.Prettify(s)
}

type EngineDefaults struct {
	_ struct{} `type:"structure"`

	DBParameterGroupFamily *string `locationName:"DBParameterGroupFamily" type:"string"`

	Marker *string `locationName:"Marker" type:"string"`

	Parameters []Parameters `locationName:"Parameters" locationNameList:"Parameter" type:"list"`
}

// String returns the string representation
func (s EngineDefaults) String() string {
	return nifcloudutil.Prettify(s)
}

type EventCategoriesMapList struct {
	_ struct{} `type:"structure"`

	EventCategories []string `locationName:"EventCategories" locationNameList:"EventCategory" type:"list"`

	SourceType *string `locationName:"SourceType" type:"string"`
}

// String returns the string representation
func (s EventCategoriesMapList) String() string {
	return nifcloudutil.Prettify(s)
}

type EventSubscription struct {
	_ struct{} `type:"structure"`

	CustSubscriptionId *string `locationName:"CustSubscriptionId" type:"string"`

	Enabled *bool `locationName:"Enabled" type:"boolean"`

	EventCategoriesList []string `locationName:"EventCategoriesList" locationNameList:"EventCategory" type:"list"`

	NiftyDescription *string `locationName:"NiftyDescription" type:"string"`

	NiftyEmailAddressesList []string `locationName:"NiftyEmailAddressesList" locationNameList:"NiftyEmailAddress" type:"list"`

	SourceIdsList []string `locationName:"SourceIdsList" locationNameList:"SourceId" type:"list"`

	SourceType *string `locationName:"SourceType" type:"string"`

	Status *string `locationName:"Status" type:"string"`

	SubscriptionCreationTime *string `locationName:"SubscriptionCreationTime" type:"string"`
}

// String returns the string representation
func (s EventSubscription) String() string {
	return nifcloudutil.Prettify(s)
}

type EventSubscriptionsList struct {
	_ struct{} `type:"structure"`

	CustSubscriptionId *string `locationName:"CustSubscriptionId" type:"string"`

	Enabled *bool `locationName:"Enabled" type:"boolean"`

	EventCategoriesList []string `locationName:"EventCategoriesList" locationNameList:"EventCategory" type:"list"`

	NiftyDescription *string `locationName:"NiftyDescription" type:"string"`

	NiftyEmailAddressesList []string `locationName:"NiftyEmailAddressesList" locationNameList:"NiftyEmailAddress" type:"list"`

	SourceIdsList []string `locationName:"SourceIdsList" locationNameList:"SourceId" type:"list"`

	SourceType *string `locationName:"SourceType" type:"string"`

	Status *string `locationName:"Status" type:"string"`

	SubscriptionCreationTime *string `locationName:"SubscriptionCreationTime" type:"string"`
}

// String returns the string representation
func (s EventSubscriptionsList) String() string {
	return nifcloudutil.Prettify(s)
}

type Events struct {
	_ struct{} `type:"structure"`

	Date *time.Time `locationName:"Date" type:"timestamp"`

	EventCategories []string `locationName:"EventCategories" locationNameList:"EventCategory" type:"list"`

	Message *string `locationName:"Message" type:"string"`

	SourceIdentifier *string `locationName:"SourceIdentifier" type:"string"`

	SourceType *string `locationName:"SourceType" type:"string"`
}

// String returns the string representation
func (s Events) String() string {
	return nifcloudutil.Prettify(s)
}

type ExternalReplicationInfo struct {
	_ struct{} `type:"structure"`

	ExternalMasterAddress *string `locationName:"ExternalMasterAddress" type:"string"`

	ExternalReplicationMessage *string `locationName:"ExternalReplicationMessage" type:"string"`

	ExternalReplicationStatus *string `locationName:"ExternalReplicationStatus" type:"string"`

	ReplicationAddresses []string `locationName:"ReplicationAddresses" locationNameList:"ReplicationAddress" type:"list"`

	ReplicationPrivateAddresses []string `locationName:"ReplicationPrivateAddresses" locationNameList:"ReplicationPrivateAddress" type:"list"`
}

// String returns the string representation
func (s ExternalReplicationInfo) String() string {
	return nifcloudutil.Prettify(s)
}

type IPRanges struct {
	_ struct{} `type:"structure"`

	CIDRIP *string `locationName:"CIDRIP" type:"string"`

	Status *string `locationName:"Status" type:"string"`
}

// String returns the string representation
func (s IPRanges) String() string {
	return nifcloudutil.Prettify(s)
}

type OptionGroupMemberships struct {
	_ struct{} `type:"structure"`

	OptionGroupName *string `locationName:"OptionGroupName" type:"string"`

	Status *string `locationName:"Status" type:"string"`
}

// String returns the string representation
func (s OptionGroupMemberships) String() string {
	return nifcloudutil.Prettify(s)
}

type OrderableDBInstanceOptions struct {
	_ struct{} `type:"structure"`

	AvailabilityZones []AvailabilityZones `locationName:"AvailabilityZones" locationNameList:"AvailabilityZone" type:"list"`

	DBInstanceClass *string `locationName:"DBInstanceClass" type:"string"`

	Engine *string `locationName:"Engine" type:"string"`

	EngineVersion *string `locationName:"EngineVersion" type:"string"`

	LicenseModel *string `locationName:"LicenseModel" type:"string"`

	MultiAZCapable *bool `locationName:"MultiAZCapable" type:"boolean"`

	ReadReplicaCapable *bool `locationName:"ReadReplicaCapable" type:"boolean"`

	Vpc *bool `locationName:"Vpc" type:"boolean"`
}

// String returns the string representation
func (s OrderableDBInstanceOptions) String() string {
	return nifcloudutil.Prettify(s)
}

type Parameters struct {
	_ struct{} `type:"structure"`

	AllowedValues *string `locationName:"AllowedValues" type:"string"`

	ApplyMethod *string `locationName:"ApplyMethod" type:"string"`

	ApplyType *string `locationName:"ApplyType" type:"string"`

	DataType *string `locationName:"DataType" type:"string"`

	Description *string `locationName:"Description" type:"string"`

	IsModifiable *bool `locationName:"IsModifiable" type:"boolean"`

	MinimumEngineVersion *string `locationName:"MinimumEngineVersion" type:"string"`

	ParameterName *string `locationName:"ParameterName" type:"string"`

	ParameterValue *string `locationName:"ParameterValue" type:"string"`

	Source *string `locationName:"Source" type:"string"`
}

// String returns the string representation
func (s Parameters) String() string {
	return nifcloudutil.Prettify(s)
}

type PendingModifiedValues struct {
	_ struct{} `type:"structure"`

	AllocatedStorage *int64 `locationName:"AllocatedStorage" type:"integer"`

	BackupRetentionPeriod *int64 `locationName:"BackupRetentionPeriod" type:"integer"`

	DBInstanceClass *string `locationName:"DBInstanceClass" type:"string"`

	DBInstanceIdentifier *string `locationName:"DBInstanceIdentifier" type:"string"`

	EngineVersion *string `locationName:"EngineVersion" type:"string"`

	MasterUserPassword *string `locationName:"MasterUserPassword" type:"string"`

	MultiAZ *bool `locationName:"MultiAZ" type:"boolean"`

	Port *int64 `locationName:"Port" type:"integer"`
}

// String returns the string representation
func (s PendingModifiedValues) String() string {
	return nifcloudutil.Prettify(s)
}

type RequestDimensions struct {
	_ struct{} `type:"structure"`

	// Name is a required field
	Name *string `locationName:"Name" type:"string" required:"true"`

	// Value is a required field
	Value *string `locationName:"Value" type:"string" required:"true"`
}

// String returns the string representation
func (s RequestDimensions) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RequestDimensions) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RequestDimensions"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RequestFilter struct {
	_ struct{} `type:"structure"`

	FilterName FilterNameOfNiftyFiltersForDescribeEventSubscriptions `locationName:"FilterName" type:"string" enum:"true"`

	FilterValue *string `locationName:"FilterValue" type:"string"`
}

// String returns the string representation
func (s RequestFilter) String() string {
	return nifcloudutil.Prettify(s)
}

type RequestNiftyFilters struct {
	_ struct{} `type:"structure"`

	ListOfRequestFilter []RequestFilter `locationName:"Filter" type:"list"`
}

// String returns the string representation
func (s RequestNiftyFilters) String() string {
	return nifcloudutil.Prettify(s)
}

type RequestParameters struct {
	_ struct{} `type:"structure"`

	// ApplyMethod is a required field
	ApplyMethod ApplyMethodOfParametersForModifyDBParameterGroup `locationName:"ApplyMethod" type:"string" required:"true" enum:"true"`

	// ParameterName is a required field
	ParameterName *string `locationName:"ParameterName" type:"string" required:"true"`

	// ParameterValue is a required field
	ParameterValue *string `locationName:"ParameterValue" type:"string" required:"true"`
}

// String returns the string representation
func (s RequestParameters) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RequestParameters) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RequestParameters"}
	if len(s.ApplyMethod) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ApplyMethod"))
	}

	if s.ParameterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParameterName"))
	}

	if s.ParameterValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParameterValue"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RequestParametersOfResetDBParameterGroup struct {
	_ struct{} `type:"structure"`

	ApplyMethod ApplyMethodOfParametersForResetDBParameterGroup `locationName:"ApplyMethod" type:"string" enum:"true"`

	ParameterName *string `locationName:"ParameterName" type:"string"`
}

// String returns the string representation
func (s RequestParametersOfResetDBParameterGroup) String() string {
	return nifcloudutil.Prettify(s)
}

type ResponseMetadata struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"RequestId" type:"string"`
}

// String returns the string representation
func (s ResponseMetadata) String() string {
	return nifcloudutil.Prettify(s)
}

type StatusInfos struct {
	_ struct{} `type:"structure"`

	Message *string `locationName:"Message" type:"string"`

	Normal *bool `locationName:"Normal" type:"boolean"`

	Status *string `locationName:"Status" type:"string"`

	StatusType *string `locationName:"StatusType" type:"string"`
}

// String returns the string representation
func (s StatusInfos) String() string {
	return nifcloudutil.Prettify(s)
}
