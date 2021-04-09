// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package computingiface provides an interface to enable mocking the NIFCLOUD Computing service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package computingiface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/service/computing"
)

// ClientAPI provides an interface to enable mocking the
// computing.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // computing.
//    func myFunc(svc computingiface.ClientAPI) bool {
//        // Make svc.AllocateAddress request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := computing.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        computingiface.ClientPI
//    }
//    func (m *mockClientClient) AllocateAddress(input *computing.AllocateAddressInput) (*computing.AllocateAddressOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	AllocateAddressRequest(*computing.AllocateAddressInput) computing.AllocateAddressRequest

	AssociateAddressRequest(*computing.AssociateAddressInput) computing.AssociateAddressRequest

	AssociateMultiIpAddressGroupRequest(*computing.AssociateMultiIpAddressGroupInput) computing.AssociateMultiIpAddressGroupRequest

	AssociateRouteTableRequest(*computing.AssociateRouteTableInput) computing.AssociateRouteTableRequest

	AssociateUsersRequest(*computing.AssociateUsersInput) computing.AssociateUsersRequest

	AttachIsoImageRequest(*computing.AttachIsoImageInput) computing.AttachIsoImageRequest

	AttachNetworkInterfaceRequest(*computing.AttachNetworkInterfaceInput) computing.AttachNetworkInterfaceRequest

	AttachVolumeRequest(*computing.AttachVolumeInput) computing.AttachVolumeRequest

	AuthorizeSecurityGroupIngressRequest(*computing.AuthorizeSecurityGroupIngressInput) computing.AuthorizeSecurityGroupIngressRequest

	CancelCopyInstancesRequest(*computing.CancelCopyInstancesInput) computing.CancelCopyInstancesRequest

	CancelUploadRequest(*computing.CancelUploadInput) computing.CancelUploadRequest

	ClearLoadBalancerSessionRequest(*computing.ClearLoadBalancerSessionInput) computing.ClearLoadBalancerSessionRequest

	ConfigureHealthCheckRequest(*computing.ConfigureHealthCheckInput) computing.ConfigureHealthCheckRequest

	CopyFromBackupInstanceRequest(*computing.CopyFromBackupInstanceInput) computing.CopyFromBackupInstanceRequest

	CopyInstancesRequest(*computing.CopyInstancesInput) computing.CopyInstancesRequest

	CreateBackupInstancesRequest(*computing.CreateBackupInstancesInput) computing.CreateBackupInstancesRequest

	CreateCustomerGatewayRequest(*computing.CreateCustomerGatewayInput) computing.CreateCustomerGatewayRequest

	CreateDhcpOptionsRequest(*computing.CreateDhcpOptionsInput) computing.CreateDhcpOptionsRequest

	CreateImageRequest(*computing.CreateImageInput) computing.CreateImageRequest

	CreateInstanceBackupRuleRequest(*computing.CreateInstanceBackupRuleInput) computing.CreateInstanceBackupRuleRequest

	CreateKeyPairRequest(*computing.CreateKeyPairInput) computing.CreateKeyPairRequest

	CreateLoadBalancerRequest(*computing.CreateLoadBalancerInput) computing.CreateLoadBalancerRequest

	CreateMultiIpAddressGroupRequest(*computing.CreateMultiIpAddressGroupInput) computing.CreateMultiIpAddressGroupRequest

	CreateNetworkInterfaceRequest(*computing.CreateNetworkInterfaceInput) computing.CreateNetworkInterfaceRequest

	CreateRouteRequest(*computing.CreateRouteInput) computing.CreateRouteRequest

	CreateRouteTableRequest(*computing.CreateRouteTableInput) computing.CreateRouteTableRequest

	CreateSecurityGroupRequest(*computing.CreateSecurityGroupInput) computing.CreateSecurityGroupRequest

	CreateSslCertificateRequest(*computing.CreateSslCertificateInput) computing.CreateSslCertificateRequest

	CreateVolumeRequest(*computing.CreateVolumeInput) computing.CreateVolumeRequest

	CreateVpnConnectionRequest(*computing.CreateVpnConnectionInput) computing.CreateVpnConnectionRequest

	CreateVpnGatewayRequest(*computing.CreateVpnGatewayInput) computing.CreateVpnGatewayRequest

	DeleteCustomerGatewayRequest(*computing.DeleteCustomerGatewayInput) computing.DeleteCustomerGatewayRequest

	DeleteDhcpOptionsRequest(*computing.DeleteDhcpOptionsInput) computing.DeleteDhcpOptionsRequest

	DeleteImageRequest(*computing.DeleteImageInput) computing.DeleteImageRequest

	DeleteInstanceBackupRuleRequest(*computing.DeleteInstanceBackupRuleInput) computing.DeleteInstanceBackupRuleRequest

	DeleteIsoImageRequest(*computing.DeleteIsoImageInput) computing.DeleteIsoImageRequest

	DeleteKeyPairRequest(*computing.DeleteKeyPairInput) computing.DeleteKeyPairRequest

	DeleteLoadBalancerRequest(*computing.DeleteLoadBalancerInput) computing.DeleteLoadBalancerRequest

	DeleteMultiIpAddressGroupRequest(*computing.DeleteMultiIpAddressGroupInput) computing.DeleteMultiIpAddressGroupRequest

	DeleteNetworkInterfaceRequest(*computing.DeleteNetworkInterfaceInput) computing.DeleteNetworkInterfaceRequest

	DeleteRouteRequest(*computing.DeleteRouteInput) computing.DeleteRouteRequest

	DeleteRouteTableRequest(*computing.DeleteRouteTableInput) computing.DeleteRouteTableRequest

	DeleteSecurityGroupRequest(*computing.DeleteSecurityGroupInput) computing.DeleteSecurityGroupRequest

	DeleteSslCertificateRequest(*computing.DeleteSslCertificateInput) computing.DeleteSslCertificateRequest

	DeleteVolumeRequest(*computing.DeleteVolumeInput) computing.DeleteVolumeRequest

	DeleteVpnConnectionRequest(*computing.DeleteVpnConnectionInput) computing.DeleteVpnConnectionRequest

	DeleteVpnGatewayRequest(*computing.DeleteVpnGatewayInput) computing.DeleteVpnGatewayRequest

	DeregisterInstancesFromLoadBalancerRequest(*computing.DeregisterInstancesFromLoadBalancerInput) computing.DeregisterInstancesFromLoadBalancerRequest

	DeregisterInstancesFromSecurityGroupRequest(*computing.DeregisterInstancesFromSecurityGroupInput) computing.DeregisterInstancesFromSecurityGroupRequest

	DescribeAddressesRequest(*computing.DescribeAddressesInput) computing.DescribeAddressesRequest

	DescribeAssociatedUsersRequest(*computing.DescribeAssociatedUsersInput) computing.DescribeAssociatedUsersRequest

	DescribeAvailabilityZonesRequest(*computing.DescribeAvailabilityZonesInput) computing.DescribeAvailabilityZonesRequest

	DescribeCustomerGatewaysRequest(*computing.DescribeCustomerGatewaysInput) computing.DescribeCustomerGatewaysRequest

	DescribeDhcpOptionsRequest(*computing.DescribeDhcpOptionsInput) computing.DescribeDhcpOptionsRequest

	DescribeImagesRequest(*computing.DescribeImagesInput) computing.DescribeImagesRequest

	DescribeInstanceAttributeRequest(*computing.DescribeInstanceAttributeInput) computing.DescribeInstanceAttributeRequest

	DescribeInstanceBackupRuleActivitiesRequest(*computing.DescribeInstanceBackupRuleActivitiesInput) computing.DescribeInstanceBackupRuleActivitiesRequest

	DescribeInstanceBackupRulesRequest(*computing.DescribeInstanceBackupRulesInput) computing.DescribeInstanceBackupRulesRequest

	DescribeInstanceHealthRequest(*computing.DescribeInstanceHealthInput) computing.DescribeInstanceHealthRequest

	DescribeInstancesRequest(*computing.DescribeInstancesInput) computing.DescribeInstancesRequest

	DescribeIsoImagesRequest(*computing.DescribeIsoImagesInput) computing.DescribeIsoImagesRequest

	DescribeKeyPairsRequest(*computing.DescribeKeyPairsInput) computing.DescribeKeyPairsRequest

	DescribeLoadBalancersRequest(*computing.DescribeLoadBalancersInput) computing.DescribeLoadBalancersRequest

	DescribeMultiIpAddressGroupsRequest(*computing.DescribeMultiIpAddressGroupsInput) computing.DescribeMultiIpAddressGroupsRequest

	DescribeNetworkInterfacesRequest(*computing.DescribeNetworkInterfacesInput) computing.DescribeNetworkInterfacesRequest

	DescribeRegionsRequest(*computing.DescribeRegionsInput) computing.DescribeRegionsRequest

	DescribeResourcesRequest(*computing.DescribeResourcesInput) computing.DescribeResourcesRequest

	DescribeRouteTablesRequest(*computing.DescribeRouteTablesInput) computing.DescribeRouteTablesRequest

	DescribeSecurityActivitiesRequest(*computing.DescribeSecurityActivitiesInput) computing.DescribeSecurityActivitiesRequest

	DescribeSecurityGroupsRequest(*computing.DescribeSecurityGroupsInput) computing.DescribeSecurityGroupsRequest

	DescribeServiceStatusRequest(*computing.DescribeServiceStatusInput) computing.DescribeServiceStatusRequest

	DescribeSslCertificateAttributeRequest(*computing.DescribeSslCertificateAttributeInput) computing.DescribeSslCertificateAttributeRequest

	DescribeSslCertificatesRequest(*computing.DescribeSslCertificatesInput) computing.DescribeSslCertificatesRequest

	DescribeUploadsRequest(*computing.DescribeUploadsInput) computing.DescribeUploadsRequest

	DescribeUsageRequest(*computing.DescribeUsageInput) computing.DescribeUsageRequest

	DescribeUserActivitiesRequest(*computing.DescribeUserActivitiesInput) computing.DescribeUserActivitiesRequest

	DescribeVolumesRequest(*computing.DescribeVolumesInput) computing.DescribeVolumesRequest

	DescribeVpnConnectionsRequest(*computing.DescribeVpnConnectionsInput) computing.DescribeVpnConnectionsRequest

	DescribeVpnGatewaysRequest(*computing.DescribeVpnGatewaysInput) computing.DescribeVpnGatewaysRequest

	DetachIsoImageRequest(*computing.DetachIsoImageInput) computing.DetachIsoImageRequest

	DetachNetworkInterfaceRequest(*computing.DetachNetworkInterfaceInput) computing.DetachNetworkInterfaceRequest

	DetachVolumeRequest(*computing.DetachVolumeInput) computing.DetachVolumeRequest

	DisassociateAddressRequest(*computing.DisassociateAddressInput) computing.DisassociateAddressRequest

	DisassociateMultiIpAddressGroupRequest(*computing.DisassociateMultiIpAddressGroupInput) computing.DisassociateMultiIpAddressGroupRequest

	DisassociateRouteTableRequest(*computing.DisassociateRouteTableInput) computing.DisassociateRouteTableRequest

	DissociateUsersRequest(*computing.DissociateUsersInput) computing.DissociateUsersRequest

	DownloadSslCertificateRequest(*computing.DownloadSslCertificateInput) computing.DownloadSslCertificateRequest

	ExtendVolumeSizeRequest(*computing.ExtendVolumeSizeInput) computing.ExtendVolumeSizeRequest

	ImportInstanceRequest(*computing.ImportInstanceInput) computing.ImportInstanceRequest

	ImportKeyPairRequest(*computing.ImportKeyPairInput) computing.ImportKeyPairRequest

	IncreaseMultiIpAddressCountRequest(*computing.IncreaseMultiIpAddressCountInput) computing.IncreaseMultiIpAddressCountRequest

	ModifyImageAttributeRequest(*computing.ModifyImageAttributeInput) computing.ModifyImageAttributeRequest

	ModifyInstanceAttributeRequest(*computing.ModifyInstanceAttributeInput) computing.ModifyInstanceAttributeRequest

	ModifyInstanceBackupRuleAttributeRequest(*computing.ModifyInstanceBackupRuleAttributeInput) computing.ModifyInstanceBackupRuleAttributeRequest

	ModifyMultiIpAddressGroupAttributeRequest(*computing.ModifyMultiIpAddressGroupAttributeInput) computing.ModifyMultiIpAddressGroupAttributeRequest

	ModifyNetworkInterfaceAttributeRequest(*computing.ModifyNetworkInterfaceAttributeInput) computing.ModifyNetworkInterfaceAttributeRequest

	ModifySslCertificateAttributeRequest(*computing.ModifySslCertificateAttributeInput) computing.ModifySslCertificateAttributeRequest

	ModifyVolumeAttributeRequest(*computing.ModifyVolumeAttributeInput) computing.ModifyVolumeAttributeRequest

	NiftyAssociateImageRequest(*computing.NiftyAssociateImageInput) computing.NiftyAssociateImageRequest

	NiftyAssociateNatTableRequest(*computing.NiftyAssociateNatTableInput) computing.NiftyAssociateNatTableRequest

	NiftyAssociateRouteTableWithElasticLoadBalancerRequest(*computing.NiftyAssociateRouteTableWithElasticLoadBalancerInput) computing.NiftyAssociateRouteTableWithElasticLoadBalancerRequest

	NiftyAssociateRouteTableWithVpnGatewayRequest(*computing.NiftyAssociateRouteTableWithVpnGatewayInput) computing.NiftyAssociateRouteTableWithVpnGatewayRequest

	NiftyConfigureElasticLoadBalancerHealthCheckRequest(*computing.NiftyConfigureElasticLoadBalancerHealthCheckInput) computing.NiftyConfigureElasticLoadBalancerHealthCheckRequest

	NiftyCreateAlarmRequest(*computing.NiftyCreateAlarmInput) computing.NiftyCreateAlarmRequest

	NiftyCreateAutoScalingGroupRequest(*computing.NiftyCreateAutoScalingGroupInput) computing.NiftyCreateAutoScalingGroupRequest

	NiftyCreateDhcpConfigRequest(*computing.NiftyCreateDhcpConfigInput) computing.NiftyCreateDhcpConfigRequest

	NiftyCreateDhcpIpAddressPoolRequest(*computing.NiftyCreateDhcpIpAddressPoolInput) computing.NiftyCreateDhcpIpAddressPoolRequest

	NiftyCreateDhcpStaticMappingRequest(*computing.NiftyCreateDhcpStaticMappingInput) computing.NiftyCreateDhcpStaticMappingRequest

	NiftyCreateElasticLoadBalancerRequest(*computing.NiftyCreateElasticLoadBalancerInput) computing.NiftyCreateElasticLoadBalancerRequest

	NiftyCreateInstanceSnapshotRequest(*computing.NiftyCreateInstanceSnapshotInput) computing.NiftyCreateInstanceSnapshotRequest

	NiftyCreateNatRuleRequest(*computing.NiftyCreateNatRuleInput) computing.NiftyCreateNatRuleRequest

	NiftyCreateNatTableRequest(*computing.NiftyCreateNatTableInput) computing.NiftyCreateNatTableRequest

	NiftyCreatePrivateLanRequest(*computing.NiftyCreatePrivateLanInput) computing.NiftyCreatePrivateLanRequest

	NiftyCreateRouterRequest(*computing.NiftyCreateRouterInput) computing.NiftyCreateRouterRequest

	NiftyCreateSeparateInstanceRuleRequest(*computing.NiftyCreateSeparateInstanceRuleInput) computing.NiftyCreateSeparateInstanceRuleRequest

	NiftyCreateWebProxyRequest(*computing.NiftyCreateWebProxyInput) computing.NiftyCreateWebProxyRequest

	NiftyDeleteAlarmRequest(*computing.NiftyDeleteAlarmInput) computing.NiftyDeleteAlarmRequest

	NiftyDeleteAutoScalingGroupRequest(*computing.NiftyDeleteAutoScalingGroupInput) computing.NiftyDeleteAutoScalingGroupRequest

	NiftyDeleteDhcpConfigRequest(*computing.NiftyDeleteDhcpConfigInput) computing.NiftyDeleteDhcpConfigRequest

	NiftyDeleteDhcpIpAddressPoolRequest(*computing.NiftyDeleteDhcpIpAddressPoolInput) computing.NiftyDeleteDhcpIpAddressPoolRequest

	NiftyDeleteDhcpStaticMappingRequest(*computing.NiftyDeleteDhcpStaticMappingInput) computing.NiftyDeleteDhcpStaticMappingRequest

	NiftyDeleteElasticLoadBalancerRequest(*computing.NiftyDeleteElasticLoadBalancerInput) computing.NiftyDeleteElasticLoadBalancerRequest

	NiftyDeleteInstanceSnapshotRequest(*computing.NiftyDeleteInstanceSnapshotInput) computing.NiftyDeleteInstanceSnapshotRequest

	NiftyDeleteNatRuleRequest(*computing.NiftyDeleteNatRuleInput) computing.NiftyDeleteNatRuleRequest

	NiftyDeleteNatTableRequest(*computing.NiftyDeleteNatTableInput) computing.NiftyDeleteNatTableRequest

	NiftyDeletePrivateLanRequest(*computing.NiftyDeletePrivateLanInput) computing.NiftyDeletePrivateLanRequest

	NiftyDeleteRouterRequest(*computing.NiftyDeleteRouterInput) computing.NiftyDeleteRouterRequest

	NiftyDeleteSeparateInstanceRuleRequest(*computing.NiftyDeleteSeparateInstanceRuleInput) computing.NiftyDeleteSeparateInstanceRuleRequest

	NiftyDeleteWebProxyRequest(*computing.NiftyDeleteWebProxyInput) computing.NiftyDeleteWebProxyRequest

	NiftyDeregisterInstancesFromElasticLoadBalancerRequest(*computing.NiftyDeregisterInstancesFromElasticLoadBalancerInput) computing.NiftyDeregisterInstancesFromElasticLoadBalancerRequest

	NiftyDeregisterInstancesFromSeparateInstanceRuleRequest(*computing.NiftyDeregisterInstancesFromSeparateInstanceRuleInput) computing.NiftyDeregisterInstancesFromSeparateInstanceRuleRequest

	NiftyDeregisterRoutersFromSecurityGroupRequest(*computing.NiftyDeregisterRoutersFromSecurityGroupInput) computing.NiftyDeregisterRoutersFromSecurityGroupRequest

	NiftyDeregisterVpnGatewaysFromSecurityGroupRequest(*computing.NiftyDeregisterVpnGatewaysFromSecurityGroupInput) computing.NiftyDeregisterVpnGatewaysFromSecurityGroupRequest

	NiftyDescribeAlarmHistoryRequest(*computing.NiftyDescribeAlarmHistoryInput) computing.NiftyDescribeAlarmHistoryRequest

	NiftyDescribeAlarmRulesActivitiesRequest(*computing.NiftyDescribeAlarmRulesActivitiesInput) computing.NiftyDescribeAlarmRulesActivitiesRequest

	NiftyDescribeAlarmsRequest(*computing.NiftyDescribeAlarmsInput) computing.NiftyDescribeAlarmsRequest

	NiftyDescribeAlarmsPartitionsRequest(*computing.NiftyDescribeAlarmsPartitionsInput) computing.NiftyDescribeAlarmsPartitionsRequest

	NiftyDescribeAutoScalingGroupsRequest(*computing.NiftyDescribeAutoScalingGroupsInput) computing.NiftyDescribeAutoScalingGroupsRequest

	NiftyDescribeCorporateInfoForCertificateRequest(*computing.NiftyDescribeCorporateInfoForCertificateInput) computing.NiftyDescribeCorporateInfoForCertificateRequest

	NiftyDescribeDhcpConfigsRequest(*computing.NiftyDescribeDhcpConfigsInput) computing.NiftyDescribeDhcpConfigsRequest

	NiftyDescribeDhcpStatusRequest(*computing.NiftyDescribeDhcpStatusInput) computing.NiftyDescribeDhcpStatusRequest

	NiftyDescribeElasticLoadBalancersRequest(*computing.NiftyDescribeElasticLoadBalancersInput) computing.NiftyDescribeElasticLoadBalancersRequest

	NiftyDescribeInstanceElasticLoadBalancerHealthRequest(*computing.NiftyDescribeInstanceElasticLoadBalancerHealthInput) computing.NiftyDescribeInstanceElasticLoadBalancerHealthRequest

	NiftyDescribeInstanceSnapshotsRequest(*computing.NiftyDescribeInstanceSnapshotsInput) computing.NiftyDescribeInstanceSnapshotsRequest

	NiftyDescribeLoadBalancerSSLPoliciesRequest(*computing.NiftyDescribeLoadBalancerSSLPoliciesInput) computing.NiftyDescribeLoadBalancerSSLPoliciesRequest

	NiftyDescribeNatTablesRequest(*computing.NiftyDescribeNatTablesInput) computing.NiftyDescribeNatTablesRequest

	NiftyDescribePerformanceChartRequest(*computing.NiftyDescribePerformanceChartInput) computing.NiftyDescribePerformanceChartRequest

	NiftyDescribePrivateLansRequest(*computing.NiftyDescribePrivateLansInput) computing.NiftyDescribePrivateLansRequest

	NiftyDescribeRoutersRequest(*computing.NiftyDescribeRoutersInput) computing.NiftyDescribeRoutersRequest

	NiftyDescribeScalingActivitiesRequest(*computing.NiftyDescribeScalingActivitiesInput) computing.NiftyDescribeScalingActivitiesRequest

	NiftyDescribeSeparateInstanceRulesRequest(*computing.NiftyDescribeSeparateInstanceRulesInput) computing.NiftyDescribeSeparateInstanceRulesRequest

	NiftyDescribeVpnGatewayActivitiesRequest(*computing.NiftyDescribeVpnGatewayActivitiesInput) computing.NiftyDescribeVpnGatewayActivitiesRequest

	NiftyDescribeWebProxiesRequest(*computing.NiftyDescribeWebProxiesInput) computing.NiftyDescribeWebProxiesRequest

	NiftyDisableDhcpRequest(*computing.NiftyDisableDhcpInput) computing.NiftyDisableDhcpRequest

	NiftyDisassociateNatTableRequest(*computing.NiftyDisassociateNatTableInput) computing.NiftyDisassociateNatTableRequest

	NiftyDisassociateRouteTableFromElasticLoadBalancerRequest(*computing.NiftyDisassociateRouteTableFromElasticLoadBalancerInput) computing.NiftyDisassociateRouteTableFromElasticLoadBalancerRequest

	NiftyDisassociateRouteTableFromVpnGatewayRequest(*computing.NiftyDisassociateRouteTableFromVpnGatewayInput) computing.NiftyDisassociateRouteTableFromVpnGatewayRequest

	NiftyEnableDhcpRequest(*computing.NiftyEnableDhcpInput) computing.NiftyEnableDhcpRequest

	NiftyModifyAddressAttributeRequest(*computing.NiftyModifyAddressAttributeInput) computing.NiftyModifyAddressAttributeRequest

	NiftyModifyCustomerGatewayAttributeRequest(*computing.NiftyModifyCustomerGatewayAttributeInput) computing.NiftyModifyCustomerGatewayAttributeRequest

	NiftyModifyElasticLoadBalancerAttributesRequest(*computing.NiftyModifyElasticLoadBalancerAttributesInput) computing.NiftyModifyElasticLoadBalancerAttributesRequest

	NiftyModifyInstanceSnapshotAttributeRequest(*computing.NiftyModifyInstanceSnapshotAttributeInput) computing.NiftyModifyInstanceSnapshotAttributeRequest

	NiftyModifyKeyPairAttributeRequest(*computing.NiftyModifyKeyPairAttributeInput) computing.NiftyModifyKeyPairAttributeRequest

	NiftyModifyPrivateLanAttributeRequest(*computing.NiftyModifyPrivateLanAttributeInput) computing.NiftyModifyPrivateLanAttributeRequest

	NiftyModifyRouterAttributeRequest(*computing.NiftyModifyRouterAttributeInput) computing.NiftyModifyRouterAttributeRequest

	NiftyModifyVpnGatewayAttributeRequest(*computing.NiftyModifyVpnGatewayAttributeInput) computing.NiftyModifyVpnGatewayAttributeRequest

	NiftyModifyWebProxyAttributeRequest(*computing.NiftyModifyWebProxyAttributeInput) computing.NiftyModifyWebProxyAttributeRequest

	NiftyRebootRoutersRequest(*computing.NiftyRebootRoutersInput) computing.NiftyRebootRoutersRequest

	NiftyRebootVpnGatewaysRequest(*computing.NiftyRebootVpnGatewaysInput) computing.NiftyRebootVpnGatewaysRequest

	NiftyRegisterInstancesWithElasticLoadBalancerRequest(*computing.NiftyRegisterInstancesWithElasticLoadBalancerInput) computing.NiftyRegisterInstancesWithElasticLoadBalancerRequest

	NiftyRegisterInstancesWithSeparateInstanceRuleRequest(*computing.NiftyRegisterInstancesWithSeparateInstanceRuleInput) computing.NiftyRegisterInstancesWithSeparateInstanceRuleRequest

	NiftyRegisterPortWithElasticLoadBalancerRequest(*computing.NiftyRegisterPortWithElasticLoadBalancerInput) computing.NiftyRegisterPortWithElasticLoadBalancerRequest

	NiftyRegisterRoutersWithSecurityGroupRequest(*computing.NiftyRegisterRoutersWithSecurityGroupInput) computing.NiftyRegisterRoutersWithSecurityGroupRequest

	NiftyRegisterVpnGatewaysWithSecurityGroupRequest(*computing.NiftyRegisterVpnGatewaysWithSecurityGroupInput) computing.NiftyRegisterVpnGatewaysWithSecurityGroupRequest

	NiftyReleaseRouterBackupStateRequest(*computing.NiftyReleaseRouterBackupStateInput) computing.NiftyReleaseRouterBackupStateRequest

	NiftyReleaseVpnGatewayBackupStateRequest(*computing.NiftyReleaseVpnGatewayBackupStateInput) computing.NiftyReleaseVpnGatewayBackupStateRequest

	NiftyReplaceDhcpConfigRequest(*computing.NiftyReplaceDhcpConfigInput) computing.NiftyReplaceDhcpConfigRequest

	NiftyReplaceDhcpOptionRequest(*computing.NiftyReplaceDhcpOptionInput) computing.NiftyReplaceDhcpOptionRequest

	NiftyReplaceElasticLoadBalancerLatestVersionRequest(*computing.NiftyReplaceElasticLoadBalancerLatestVersionInput) computing.NiftyReplaceElasticLoadBalancerLatestVersionRequest

	NiftyReplaceElasticLoadBalancerListenerSSLCertificateRequest(*computing.NiftyReplaceElasticLoadBalancerListenerSSLCertificateInput) computing.NiftyReplaceElasticLoadBalancerListenerSSLCertificateRequest

	NiftyReplaceNatRuleRequest(*computing.NiftyReplaceNatRuleInput) computing.NiftyReplaceNatRuleRequest

	NiftyReplaceNatTableAssociationRequest(*computing.NiftyReplaceNatTableAssociationInput) computing.NiftyReplaceNatTableAssociationRequest

	NiftyReplaceRouteTableAssociationWithElasticLoadBalancerRequest(*computing.NiftyReplaceRouteTableAssociationWithElasticLoadBalancerInput) computing.NiftyReplaceRouteTableAssociationWithElasticLoadBalancerRequest

	NiftyReplaceRouteTableAssociationWithVpnGatewayRequest(*computing.NiftyReplaceRouteTableAssociationWithVpnGatewayInput) computing.NiftyReplaceRouteTableAssociationWithVpnGatewayRequest

	NiftyReplaceRouterLatestVersionRequest(*computing.NiftyReplaceRouterLatestVersionInput) computing.NiftyReplaceRouterLatestVersionRequest

	NiftyReplaceVpnGatewayLatestVersionRequest(*computing.NiftyReplaceVpnGatewayLatestVersionInput) computing.NiftyReplaceVpnGatewayLatestVersionRequest

	NiftyRestoreInstanceSnapshotRequest(*computing.NiftyRestoreInstanceSnapshotInput) computing.NiftyRestoreInstanceSnapshotRequest

	NiftyRestoreRouterPreviousVersionRequest(*computing.NiftyRestoreRouterPreviousVersionInput) computing.NiftyRestoreRouterPreviousVersionRequest

	NiftyRestoreVpnGatewayPreviousVersionRequest(*computing.NiftyRestoreVpnGatewayPreviousVersionInput) computing.NiftyRestoreVpnGatewayPreviousVersionRequest

	NiftyRetryImportInstanceRequest(*computing.NiftyRetryImportInstanceInput) computing.NiftyRetryImportInstanceRequest

	NiftySetLoadBalancerSSLPoliciesOfListenerRequest(*computing.NiftySetLoadBalancerSSLPoliciesOfListenerInput) computing.NiftySetLoadBalancerSSLPoliciesOfListenerRequest

	NiftyUnsetLoadBalancerSSLPoliciesOfListenerRequest(*computing.NiftyUnsetLoadBalancerSSLPoliciesOfListenerInput) computing.NiftyUnsetLoadBalancerSSLPoliciesOfListenerRequest

	NiftyUpdateAlarmRequest(*computing.NiftyUpdateAlarmInput) computing.NiftyUpdateAlarmRequest

	NiftyUpdateAutoScalingGroupRequest(*computing.NiftyUpdateAutoScalingGroupInput) computing.NiftyUpdateAutoScalingGroupRequest

	NiftyUpdateElasticLoadBalancerRequest(*computing.NiftyUpdateElasticLoadBalancerInput) computing.NiftyUpdateElasticLoadBalancerRequest

	NiftyUpdateInstanceNetworkInterfacesRequest(*computing.NiftyUpdateInstanceNetworkInterfacesInput) computing.NiftyUpdateInstanceNetworkInterfacesRequest

	NiftyUpdateRouterNetworkInterfacesRequest(*computing.NiftyUpdateRouterNetworkInterfacesInput) computing.NiftyUpdateRouterNetworkInterfacesRequest

	NiftyUpdateSeparateInstanceRuleRequest(*computing.NiftyUpdateSeparateInstanceRuleInput) computing.NiftyUpdateSeparateInstanceRuleRequest

	NiftyUpdateVpnGatewayNetworkInterfacesRequest(*computing.NiftyUpdateVpnGatewayNetworkInterfacesInput) computing.NiftyUpdateVpnGatewayNetworkInterfacesRequest

	RebootInstancesRequest(*computing.RebootInstancesInput) computing.RebootInstancesRequest

	RefreshInstanceBackupRuleRequest(*computing.RefreshInstanceBackupRuleInput) computing.RefreshInstanceBackupRuleRequest

	RegisterCorporateInfoForCertificateRequest(*computing.RegisterCorporateInfoForCertificateInput) computing.RegisterCorporateInfoForCertificateRequest

	RegisterInstancesWithLoadBalancerRequest(*computing.RegisterInstancesWithLoadBalancerInput) computing.RegisterInstancesWithLoadBalancerRequest

	RegisterInstancesWithSecurityGroupRequest(*computing.RegisterInstancesWithSecurityGroupInput) computing.RegisterInstancesWithSecurityGroupRequest

	RegisterPortWithLoadBalancerRequest(*computing.RegisterPortWithLoadBalancerInput) computing.RegisterPortWithLoadBalancerRequest

	ReleaseAddressRequest(*computing.ReleaseAddressInput) computing.ReleaseAddressRequest

	ReleaseMultiIpAddressesRequest(*computing.ReleaseMultiIpAddressesInput) computing.ReleaseMultiIpAddressesRequest

	ReplaceRouteRequest(*computing.ReplaceRouteInput) computing.ReplaceRouteRequest

	ReplaceRouteTableAssociationRequest(*computing.ReplaceRouteTableAssociationInput) computing.ReplaceRouteTableAssociationRequest

	RevokeSecurityGroupIngressRequest(*computing.RevokeSecurityGroupIngressInput) computing.RevokeSecurityGroupIngressRequest

	RunInstancesRequest(*computing.RunInstancesInput) computing.RunInstancesRequest

	SetFilterForLoadBalancerRequest(*computing.SetFilterForLoadBalancerInput) computing.SetFilterForLoadBalancerRequest

	SetLoadBalancerListenerSSLCertificateRequest(*computing.SetLoadBalancerListenerSSLCertificateInput) computing.SetLoadBalancerListenerSSLCertificateRequest

	StartInstancesRequest(*computing.StartInstancesInput) computing.StartInstancesRequest

	StopInstancesRequest(*computing.StopInstancesInput) computing.StopInstancesRequest

	TerminateInstancesRequest(*computing.TerminateInstancesInput) computing.TerminateInstancesRequest

	UnsetLoadBalancerListenerSSLCertificateRequest(*computing.UnsetLoadBalancerListenerSSLCertificateInput) computing.UnsetLoadBalancerListenerSSLCertificateRequest

	UpdateLoadBalancerRequest(*computing.UpdateLoadBalancerInput) computing.UpdateLoadBalancerRequest

	UpdateLoadBalancerOptionRequest(*computing.UpdateLoadBalancerOptionInput) computing.UpdateLoadBalancerOptionRequest

	UpdateSecurityGroupRequest(*computing.UpdateSecurityGroupInput) computing.UpdateSecurityGroupRequest

	UploadIsoImageRequest(*computing.UploadIsoImageInput) computing.UploadIsoImageRequest

	UploadSslCertificateRequest(*computing.UploadSslCertificateInput) computing.UploadSslCertificateRequest

	WaitUntilCustomerGatewayAvailable(context.Context, *computing.DescribeCustomerGatewaysInput, ...aws.WaiterOption) error

	WaitUntilCustomerGatewayDeleted(context.Context, *computing.DescribeCustomerGatewaysInput, ...aws.WaiterOption) error

	WaitUntilCustomerGatewayExists(context.Context, *computing.DescribeCustomerGatewaysInput, ...aws.WaiterOption) error

	WaitUntilCustomerGatewayStopped(context.Context, *computing.DescribeCustomerGatewaysInput, ...aws.WaiterOption) error

	WaitUntilCustomerGatewayWarning(context.Context, *computing.DescribeCustomerGatewaysInput, ...aws.WaiterOption) error

	WaitUntilElasticLoadBalancerAvailable(context.Context, *computing.NiftyDescribeElasticLoadBalancersInput, ...aws.WaiterOption) error

	WaitUntilElasticLoadBalancerDeleted(context.Context, *computing.NiftyDescribeElasticLoadBalancersInput, ...aws.WaiterOption) error

	WaitUntilElasticLoadBalancerExists(context.Context, *computing.NiftyDescribeElasticLoadBalancersInput, ...aws.WaiterOption) error

	WaitUntilInstanceDeleted(context.Context, *computing.DescribeInstancesInput, ...aws.WaiterOption) error

	WaitUntilInstanceExists(context.Context, *computing.DescribeInstancesInput, ...aws.WaiterOption) error

	WaitUntilInstanceImportError(context.Context, *computing.DescribeInstancesInput, ...aws.WaiterOption) error

	WaitUntilInstanceRunning(context.Context, *computing.DescribeInstancesInput, ...aws.WaiterOption) error

	WaitUntilInstanceStopped(context.Context, *computing.DescribeInstancesInput, ...aws.WaiterOption) error

	WaitUntilInstanceSuspending(context.Context, *computing.DescribeInstancesInput, ...aws.WaiterOption) error

	WaitUntilInstanceWarning(context.Context, *computing.DescribeInstancesInput, ...aws.WaiterOption) error

	WaitUntilLoadBalancerDeleted(context.Context, *computing.DescribeLoadBalancersInput, ...aws.WaiterOption) error

	WaitUntilLoadBalancerExists(context.Context, *computing.DescribeLoadBalancersInput, ...aws.WaiterOption) error

	WaitUntilLoadBalancerInService(context.Context, *computing.DescribeLoadBalancersInput, ...aws.WaiterOption) error

	WaitUntilPrivateLanAvailable(context.Context, *computing.NiftyDescribePrivateLansInput, ...aws.WaiterOption) error

	WaitUntilPrivateLanDeleted(context.Context, *computing.NiftyDescribePrivateLansInput, ...aws.WaiterOption) error

	WaitUntilPrivateLanExists(context.Context, *computing.NiftyDescribePrivateLansInput, ...aws.WaiterOption) error

	WaitUntilRouterAvailable(context.Context, *computing.NiftyDescribeRoutersInput, ...aws.WaiterOption) error

	WaitUntilRouterDeleted(context.Context, *computing.NiftyDescribeRoutersInput, ...aws.WaiterOption) error

	WaitUntilRouterExists(context.Context, *computing.NiftyDescribeRoutersInput, ...aws.WaiterOption) error

	WaitUntilRouterStopped(context.Context, *computing.NiftyDescribeRoutersInput, ...aws.WaiterOption) error

	WaitUntilRouterWarning(context.Context, *computing.NiftyDescribeRoutersInput, ...aws.WaiterOption) error

	WaitUntilSecurityGroupApplied(context.Context, *computing.DescribeSecurityGroupsInput, ...aws.WaiterOption) error

	WaitUntilSecurityGroupDeleted(context.Context, *computing.DescribeSecurityGroupsInput, ...aws.WaiterOption) error

	WaitUntilSecurityGroupExists(context.Context, *computing.DescribeSecurityGroupsInput, ...aws.WaiterOption) error

	WaitUntilSnapshotDeleted(context.Context, *computing.NiftyDescribeInstanceSnapshotsInput, ...aws.WaiterOption) error

	WaitUntilSnapshotExists(context.Context, *computing.NiftyDescribeInstanceSnapshotsInput, ...aws.WaiterOption) error

	WaitUntilSnapshotNormal(context.Context, *computing.NiftyDescribeInstanceSnapshotsInput, ...aws.WaiterOption) error

	WaitUntilVolumeAttached(context.Context, *computing.DescribeVolumesInput, ...aws.WaiterOption) error

	WaitUntilVolumeAvailable(context.Context, *computing.DescribeVolumesInput, ...aws.WaiterOption) error

	WaitUntilVolumeDeleted(context.Context, *computing.DescribeVolumesInput, ...aws.WaiterOption) error

	WaitUntilVolumeExists(context.Context, *computing.DescribeVolumesInput, ...aws.WaiterOption) error

	WaitUntilVolumeInUse(context.Context, *computing.DescribeVolumesInput, ...aws.WaiterOption) error

	WaitUntilVpnConnectionAvailable(context.Context, *computing.DescribeVpnConnectionsInput, ...aws.WaiterOption) error

	WaitUntilVpnConnectionDeleted(context.Context, *computing.DescribeVpnConnectionsInput, ...aws.WaiterOption) error

	WaitUntilVpnConnectionExists(context.Context, *computing.DescribeVpnConnectionsInput, ...aws.WaiterOption) error

	WaitUntilVpnGatewayAvailable(context.Context, *computing.DescribeVpnGatewaysInput, ...aws.WaiterOption) error

	WaitUntilVpnGatewayDeleted(context.Context, *computing.DescribeVpnGatewaysInput, ...aws.WaiterOption) error

	WaitUntilVpnGatewayExists(context.Context, *computing.DescribeVpnGatewaysInput, ...aws.WaiterOption) error

	WaitUntilVpnGatewayStopped(context.Context, *computing.DescribeVpnGatewaysInput, ...aws.WaiterOption) error

	WaitUntilVpnGatewayWarning(context.Context, *computing.DescribeVpnGatewaysInput, ...aws.WaiterOption) error
}

var _ ClientAPI = (*computing.Client)(nil)
