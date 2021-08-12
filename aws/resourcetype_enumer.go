// Code generated by "enumer -type ResourceType -addprefix aws_ -transform snake -linecomment"; DO NOT EDIT.

package aws

import (
	"fmt"
	"strings"
)

const _ResourceTypeName = "aws_instanceaws_albaws_alb_listeneraws_alb_listener_certificateaws_alb_listener_ruleaws_alb_target_groupaws_alb_target_group_attachmentaws_api_gateway_deploymentaws_api_gateway_resourceaws_api_gateway_rest_apiaws_api_gateway_stageaws_athena_workgroupaws_autoscaling_groupaws_autoscaling_policyaws_batch_job_definitionaws_cloudfront_distributionaws_cloudfront_origin_access_identityaws_cloudfront_public_keyaws_cloudwatch_metric_alarmaws_dax_clusteraws_db_instanceaws_db_parameter_groupaws_db_subnet_groupaws_directory_service_directoryaws_dms_replication_instanceaws_dx_gatewayaws_dynamodb_global_tableaws_dynamodb_tableaws_ebs_volumeaws_ecs_clusteraws_ecs_serviceaws_efs_file_systemaws_eipaws_eks_clusteraws_elasticache_clusteraws_elasticache_replication_groupaws_elastic_beanstalk_applicationaws_elasticsearch_domainaws_elasticsearch_domain_policyaws_elbaws_emr_clusteraws_fsx_lustre_file_systemaws_glue_catalog_databaseaws_glue_catalog_tableaws_iam_access_keyaws_iam_account_aliasaws_iam_account_password_policyaws_iam_groupaws_iam_group_membershipaws_iam_group_policyaws_iam_group_policy_attachmentaws_iam_instance_profileaws_iam_openid_connect_provideraws_iam_policyaws_iam_roleaws_iam_role_policyaws_iam_role_policy_attachmentaws_iam_saml_provideraws_iam_server_certificateaws_iam_useraws_iam_user_group_membershipaws_iam_user_policyaws_iam_user_policy_attachmentaws_iam_user_ssh_keyaws_internet_gatewayaws_key_pairaws_kinesis_streamaws_lambda_functionaws_launch_configurationaws_launch_templateaws_lbaws_lb_cookie_stickiness_policyaws_lb_listeneraws_lb_listener_certificateaws_lb_listener_ruleaws_lb_target_groupaws_lb_target_group_attachmentaws_lightsail_instanceaws_media_store_containeraws_mq_brokeraws_nat_gatewayaws_neptune_clusteraws_rds_clusteraws_rds_global_clusteraws_redshift_clusteraws_route53_delegation_setaws_route53_health_checkaws_route53_query_logaws_route53_recordaws_route53_resolver_endpointaws_route53_resolver_rule_associationaws_route53_zoneaws_route53_zone_associationaws_s3_bucketaws_security_groupaws_ses_active_receipt_rule_setaws_ses_configuration_setaws_ses_domain_dkimaws_ses_domain_identityaws_ses_domain_mail_fromaws_ses_identity_notification_topicaws_ses_receipt_filteraws_ses_receipt_ruleaws_ses_receipt_rule_setaws_ses_templateaws_sqs_queueaws_storagegateway_gatewayaws_subnetaws_volume_attachmentaws_vpcaws_vpc_peering_connectionaws_vpn_gateway"

var _ResourceTypeIndex = [...]uint16{0, 12, 19, 35, 63, 84, 104, 135, 161, 185, 209, 230, 250, 271, 293, 317, 344, 381, 406, 433, 448, 463, 485, 504, 535, 563, 577, 602, 620, 634, 649, 664, 683, 690, 705, 728, 761, 794, 818, 849, 856, 871, 897, 922, 944, 962, 983, 1014, 1027, 1051, 1071, 1102, 1126, 1157, 1171, 1183, 1202, 1232, 1253, 1279, 1291, 1320, 1339, 1369, 1389, 1409, 1421, 1439, 1458, 1482, 1501, 1507, 1538, 1553, 1580, 1600, 1619, 1649, 1671, 1696, 1709, 1724, 1743, 1758, 1780, 1800, 1826, 1850, 1871, 1889, 1918, 1955, 1971, 1999, 2012, 2030, 2061, 2086, 2105, 2128, 2152, 2187, 2209, 2229, 2253, 2269, 2282, 2308, 2318, 2339, 2346, 2372, 2387}

const _ResourceTypeLowerName = "aws_instanceaws_albaws_alb_listeneraws_alb_listener_certificateaws_alb_listener_ruleaws_alb_target_groupaws_alb_target_group_attachmentaws_api_gateway_deploymentaws_api_gateway_resourceaws_api_gateway_rest_apiaws_api_gateway_stageaws_athena_workgroupaws_autoscaling_groupaws_autoscaling_policyaws_batch_job_definitionaws_cloudfront_distributionaws_cloudfront_origin_access_identityaws_cloudfront_public_keyaws_cloudwatch_metric_alarmaws_dax_clusteraws_db_instanceaws_db_parameter_groupaws_db_subnet_groupaws_directory_service_directoryaws_dms_replication_instanceaws_dx_gatewayaws_dynamodb_global_tableaws_dynamodb_tableaws_ebs_volumeaws_ecs_clusteraws_ecs_serviceaws_efs_file_systemaws_eipaws_eks_clusteraws_elasticache_clusteraws_elasticache_replication_groupaws_elastic_beanstalk_applicationaws_elasticsearch_domainaws_elasticsearch_domain_policyaws_elbaws_emr_clusteraws_fsx_lustre_file_systemaws_glue_catalog_databaseaws_glue_catalog_tableaws_iam_access_keyaws_iam_account_aliasaws_iam_account_password_policyaws_iam_groupaws_iam_group_membershipaws_iam_group_policyaws_iam_group_policy_attachmentaws_iam_instance_profileaws_iam_openid_connect_provideraws_iam_policyaws_iam_roleaws_iam_role_policyaws_iam_role_policy_attachmentaws_iam_saml_provideraws_iam_server_certificateaws_iam_useraws_iam_user_group_membershipaws_iam_user_policyaws_iam_user_policy_attachmentaws_iam_user_ssh_keyaws_internet_gatewayaws_key_pairaws_kinesis_streamaws_lambda_functionaws_launch_configurationaws_launch_templateaws_lbaws_lb_cookie_stickiness_policyaws_lb_listeneraws_lb_listener_certificateaws_lb_listener_ruleaws_lb_target_groupaws_lb_target_group_attachmentaws_lightsail_instanceaws_media_store_containeraws_mq_brokeraws_nat_gatewayaws_neptune_clusteraws_rds_clusteraws_rds_global_clusteraws_redshift_clusteraws_route53_delegation_setaws_route53_health_checkaws_route53_query_logaws_route53_recordaws_route53_resolver_endpointaws_route53_resolver_rule_associationaws_route53_zoneaws_route53_zone_associationaws_s3_bucketaws_security_groupaws_ses_active_receipt_rule_setaws_ses_configuration_setaws_ses_domain_dkimaws_ses_domain_identityaws_ses_domain_mail_fromaws_ses_identity_notification_topicaws_ses_receipt_filteraws_ses_receipt_ruleaws_ses_receipt_rule_setaws_ses_templateaws_sqs_queueaws_storagegateway_gatewayaws_subnetaws_volume_attachmentaws_vpcaws_vpc_peering_connectionaws_vpn_gateway"

func (i ResourceType) String() string {
	i -= 1
	if i < 0 || i >= ResourceType(len(_ResourceTypeIndex)-1) {
		return fmt.Sprintf("ResourceType(%d)", i+1)
	}
	return _ResourceTypeName[_ResourceTypeIndex[i]:_ResourceTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ResourceTypeNoOp() {
	var x [1]struct{}
	_ = x[Instance-(1)]
	_ = x[ALB-(2)]
	_ = x[ALBListener-(3)]
	_ = x[ALBListenerCertificate-(4)]
	_ = x[ALBListenerRule-(5)]
	_ = x[ALBTargetGroup-(6)]
	_ = x[ALBTargetGroupAttachment-(7)]
	_ = x[APIGatewayDeployment-(8)]
	_ = x[APIGatewayResource-(9)]
	_ = x[APIGatewayRestAPI-(10)]
	_ = x[APIGatewayStage-(11)]
	_ = x[AthenaWorkgroup-(12)]
	_ = x[AutoscalingGroup-(13)]
	_ = x[AutoscalingPolicy-(14)]
	_ = x[BatchJobDefinition-(15)]
	_ = x[CloudfrontDistribution-(16)]
	_ = x[CloudfrontOriginAccessIdentity-(17)]
	_ = x[CloudfrontPublicKey-(18)]
	_ = x[CloudwatchMetricAlarm-(19)]
	_ = x[DaxCluster-(20)]
	_ = x[DBInstance-(21)]
	_ = x[DBParameterGroup-(22)]
	_ = x[DBSubnetGroup-(23)]
	_ = x[DirectoryServiceDirectory-(24)]
	_ = x[DmsReplicationInstance-(25)]
	_ = x[DXGateway-(26)]
	_ = x[DynamodbGlobalTable-(27)]
	_ = x[DynamodbTable-(28)]
	_ = x[EBSVolume-(29)]
	_ = x[ECSCluster-(30)]
	_ = x[ECSService-(31)]
	_ = x[EFSFileSystem-(32)]
	_ = x[EIP-(33)]
	_ = x[EKSCluster-(34)]
	_ = x[ElasticacheCluster-(35)]
	_ = x[ElasticacheReplicationGroup-(36)]
	_ = x[ElasticBeanstalkApplication-(37)]
	_ = x[ElasticsearchDomain-(38)]
	_ = x[ElasticsearchDomainPolicy-(39)]
	_ = x[ELB-(40)]
	_ = x[EMRCluster-(41)]
	_ = x[FsxLustreFileSystem-(42)]
	_ = x[GlueCatalogDatabase-(43)]
	_ = x[GlueCatalogTable-(44)]
	_ = x[IAMAccessKey-(45)]
	_ = x[IAMAccountAlias-(46)]
	_ = x[IAMAccountPasswordPolicy-(47)]
	_ = x[IAMGroup-(48)]
	_ = x[IAMGroupMembership-(49)]
	_ = x[IAMGroupPolicy-(50)]
	_ = x[IAMGroupPolicyAttachment-(51)]
	_ = x[IAMInstanceProfile-(52)]
	_ = x[IAMOpenidConnectProvider-(53)]
	_ = x[IAMPolicy-(54)]
	_ = x[IAMRole-(55)]
	_ = x[IAMRolePolicy-(56)]
	_ = x[IAMRolePolicyAttachment-(57)]
	_ = x[IAMSAMLProvider-(58)]
	_ = x[IAMServerCertificate-(59)]
	_ = x[IAMUser-(60)]
	_ = x[IAMUserGroupMembership-(61)]
	_ = x[IAMUserPolicy-(62)]
	_ = x[IAMUserPolicyAttachment-(63)]
	_ = x[IAMUserSSHKey-(64)]
	_ = x[InternetGateway-(65)]
	_ = x[KeyPair-(66)]
	_ = x[KinesisStream-(67)]
	_ = x[LambdaFunction-(68)]
	_ = x[LaunchConfiguration-(69)]
	_ = x[LaunchTemplate-(70)]
	_ = x[LB-(71)]
	_ = x[LBCookieStickinessPolicy-(72)]
	_ = x[LBListener-(73)]
	_ = x[LBListenerCertificate-(74)]
	_ = x[LBListenerRule-(75)]
	_ = x[LBTargetGroup-(76)]
	_ = x[LBTargetGroupAttachment-(77)]
	_ = x[LightsailInstance-(78)]
	_ = x[MediaStoreContainer-(79)]
	_ = x[MQBroker-(80)]
	_ = x[NatGateway-(81)]
	_ = x[NeptuneCluster-(82)]
	_ = x[RDSCluster-(83)]
	_ = x[RDSGlobalCluster-(84)]
	_ = x[RedshiftCluster-(85)]
	_ = x[Route53DelegationSet-(86)]
	_ = x[Route53HealthCheck-(87)]
	_ = x[Route53QueryLog-(88)]
	_ = x[Route53Record-(89)]
	_ = x[Route53ResolverEndpoint-(90)]
	_ = x[Route53ResolverRuleAssociation-(91)]
	_ = x[Route53Zone-(92)]
	_ = x[Route53ZoneAssociation-(93)]
	_ = x[S3Bucket-(94)]
	_ = x[SecurityGroup-(95)]
	_ = x[SESActiveReceiptRuleSet-(96)]
	_ = x[SESConfigurationSet-(97)]
	_ = x[SESDomainDKIM-(98)]
	_ = x[SESDomainIdentity-(99)]
	_ = x[SESDomainMailFrom-(100)]
	_ = x[SESIdentityNotificationTopic-(101)]
	_ = x[SESReceiptFilter-(102)]
	_ = x[SESReceiptRule-(103)]
	_ = x[SESReceiptRuleSet-(104)]
	_ = x[SESTemplate-(105)]
	_ = x[SQSQueue-(106)]
	_ = x[StoragegatewayGateway-(107)]
	_ = x[Subnet-(108)]
	_ = x[VolumeAttachment-(109)]
	_ = x[VPC-(110)]
	_ = x[VPCPeeringConnection-(111)]
	_ = x[VPNGateway-(112)]
}

var _ResourceTypeValues = []ResourceType{Instance, ALB, ALBListener, ALBListenerCertificate, ALBListenerRule, ALBTargetGroup, ALBTargetGroupAttachment, APIGatewayDeployment, APIGatewayResource, APIGatewayRestAPI, APIGatewayStage, AthenaWorkgroup, AutoscalingGroup, AutoscalingPolicy, BatchJobDefinition, CloudfrontDistribution, CloudfrontOriginAccessIdentity, CloudfrontPublicKey, CloudwatchMetricAlarm, DaxCluster, DBInstance, DBParameterGroup, DBSubnetGroup, DirectoryServiceDirectory, DmsReplicationInstance, DXGateway, DynamodbGlobalTable, DynamodbTable, EBSVolume, ECSCluster, ECSService, EFSFileSystem, EIP, EKSCluster, ElasticacheCluster, ElasticacheReplicationGroup, ElasticBeanstalkApplication, ElasticsearchDomain, ElasticsearchDomainPolicy, ELB, EMRCluster, FsxLustreFileSystem, GlueCatalogDatabase, GlueCatalogTable, IAMAccessKey, IAMAccountAlias, IAMAccountPasswordPolicy, IAMGroup, IAMGroupMembership, IAMGroupPolicy, IAMGroupPolicyAttachment, IAMInstanceProfile, IAMOpenidConnectProvider, IAMPolicy, IAMRole, IAMRolePolicy, IAMRolePolicyAttachment, IAMSAMLProvider, IAMServerCertificate, IAMUser, IAMUserGroupMembership, IAMUserPolicy, IAMUserPolicyAttachment, IAMUserSSHKey, InternetGateway, KeyPair, KinesisStream, LambdaFunction, LaunchConfiguration, LaunchTemplate, LB, LBCookieStickinessPolicy, LBListener, LBListenerCertificate, LBListenerRule, LBTargetGroup, LBTargetGroupAttachment, LightsailInstance, MediaStoreContainer, MQBroker, NatGateway, NeptuneCluster, RDSCluster, RDSGlobalCluster, RedshiftCluster, Route53DelegationSet, Route53HealthCheck, Route53QueryLog, Route53Record, Route53ResolverEndpoint, Route53ResolverRuleAssociation, Route53Zone, Route53ZoneAssociation, S3Bucket, SecurityGroup, SESActiveReceiptRuleSet, SESConfigurationSet, SESDomainDKIM, SESDomainIdentity, SESDomainMailFrom, SESIdentityNotificationTopic, SESReceiptFilter, SESReceiptRule, SESReceiptRuleSet, SESTemplate, SQSQueue, StoragegatewayGateway, Subnet, VolumeAttachment, VPC, VPCPeeringConnection, VPNGateway}

var _ResourceTypeNameToValueMap = map[string]ResourceType{
	_ResourceTypeName[0:12]:           Instance,
	_ResourceTypeLowerName[0:12]:      Instance,
	_ResourceTypeName[12:19]:          ALB,
	_ResourceTypeLowerName[12:19]:     ALB,
	_ResourceTypeName[19:35]:          ALBListener,
	_ResourceTypeLowerName[19:35]:     ALBListener,
	_ResourceTypeName[35:63]:          ALBListenerCertificate,
	_ResourceTypeLowerName[35:63]:     ALBListenerCertificate,
	_ResourceTypeName[63:84]:          ALBListenerRule,
	_ResourceTypeLowerName[63:84]:     ALBListenerRule,
	_ResourceTypeName[84:104]:         ALBTargetGroup,
	_ResourceTypeLowerName[84:104]:    ALBTargetGroup,
	_ResourceTypeName[104:135]:        ALBTargetGroupAttachment,
	_ResourceTypeLowerName[104:135]:   ALBTargetGroupAttachment,
	_ResourceTypeName[135:161]:        APIGatewayDeployment,
	_ResourceTypeLowerName[135:161]:   APIGatewayDeployment,
	_ResourceTypeName[161:185]:        APIGatewayResource,
	_ResourceTypeLowerName[161:185]:   APIGatewayResource,
	_ResourceTypeName[185:209]:        APIGatewayRestAPI,
	_ResourceTypeLowerName[185:209]:   APIGatewayRestAPI,
	_ResourceTypeName[209:230]:        APIGatewayStage,
	_ResourceTypeLowerName[209:230]:   APIGatewayStage,
	_ResourceTypeName[230:250]:        AthenaWorkgroup,
	_ResourceTypeLowerName[230:250]:   AthenaWorkgroup,
	_ResourceTypeName[250:271]:        AutoscalingGroup,
	_ResourceTypeLowerName[250:271]:   AutoscalingGroup,
	_ResourceTypeName[271:293]:        AutoscalingPolicy,
	_ResourceTypeLowerName[271:293]:   AutoscalingPolicy,
	_ResourceTypeName[293:317]:        BatchJobDefinition,
	_ResourceTypeLowerName[293:317]:   BatchJobDefinition,
	_ResourceTypeName[317:344]:        CloudfrontDistribution,
	_ResourceTypeLowerName[317:344]:   CloudfrontDistribution,
	_ResourceTypeName[344:381]:        CloudfrontOriginAccessIdentity,
	_ResourceTypeLowerName[344:381]:   CloudfrontOriginAccessIdentity,
	_ResourceTypeName[381:406]:        CloudfrontPublicKey,
	_ResourceTypeLowerName[381:406]:   CloudfrontPublicKey,
	_ResourceTypeName[406:433]:        CloudwatchMetricAlarm,
	_ResourceTypeLowerName[406:433]:   CloudwatchMetricAlarm,
	_ResourceTypeName[433:448]:        DaxCluster,
	_ResourceTypeLowerName[433:448]:   DaxCluster,
	_ResourceTypeName[448:463]:        DBInstance,
	_ResourceTypeLowerName[448:463]:   DBInstance,
	_ResourceTypeName[463:485]:        DBParameterGroup,
	_ResourceTypeLowerName[463:485]:   DBParameterGroup,
	_ResourceTypeName[485:504]:        DBSubnetGroup,
	_ResourceTypeLowerName[485:504]:   DBSubnetGroup,
	_ResourceTypeName[504:535]:        DirectoryServiceDirectory,
	_ResourceTypeLowerName[504:535]:   DirectoryServiceDirectory,
	_ResourceTypeName[535:563]:        DmsReplicationInstance,
	_ResourceTypeLowerName[535:563]:   DmsReplicationInstance,
	_ResourceTypeName[563:577]:        DXGateway,
	_ResourceTypeLowerName[563:577]:   DXGateway,
	_ResourceTypeName[577:602]:        DynamodbGlobalTable,
	_ResourceTypeLowerName[577:602]:   DynamodbGlobalTable,
	_ResourceTypeName[602:620]:        DynamodbTable,
	_ResourceTypeLowerName[602:620]:   DynamodbTable,
	_ResourceTypeName[620:634]:        EBSVolume,
	_ResourceTypeLowerName[620:634]:   EBSVolume,
	_ResourceTypeName[634:649]:        ECSCluster,
	_ResourceTypeLowerName[634:649]:   ECSCluster,
	_ResourceTypeName[649:664]:        ECSService,
	_ResourceTypeLowerName[649:664]:   ECSService,
	_ResourceTypeName[664:683]:        EFSFileSystem,
	_ResourceTypeLowerName[664:683]:   EFSFileSystem,
	_ResourceTypeName[683:690]:        EIP,
	_ResourceTypeLowerName[683:690]:   EIP,
	_ResourceTypeName[690:705]:        EKSCluster,
	_ResourceTypeLowerName[690:705]:   EKSCluster,
	_ResourceTypeName[705:728]:        ElasticacheCluster,
	_ResourceTypeLowerName[705:728]:   ElasticacheCluster,
	_ResourceTypeName[728:761]:        ElasticacheReplicationGroup,
	_ResourceTypeLowerName[728:761]:   ElasticacheReplicationGroup,
	_ResourceTypeName[761:794]:        ElasticBeanstalkApplication,
	_ResourceTypeLowerName[761:794]:   ElasticBeanstalkApplication,
	_ResourceTypeName[794:818]:        ElasticsearchDomain,
	_ResourceTypeLowerName[794:818]:   ElasticsearchDomain,
	_ResourceTypeName[818:849]:        ElasticsearchDomainPolicy,
	_ResourceTypeLowerName[818:849]:   ElasticsearchDomainPolicy,
	_ResourceTypeName[849:856]:        ELB,
	_ResourceTypeLowerName[849:856]:   ELB,
	_ResourceTypeName[856:871]:        EMRCluster,
	_ResourceTypeLowerName[856:871]:   EMRCluster,
	_ResourceTypeName[871:897]:        FsxLustreFileSystem,
	_ResourceTypeLowerName[871:897]:   FsxLustreFileSystem,
	_ResourceTypeName[897:922]:        GlueCatalogDatabase,
	_ResourceTypeLowerName[897:922]:   GlueCatalogDatabase,
	_ResourceTypeName[922:944]:        GlueCatalogTable,
	_ResourceTypeLowerName[922:944]:   GlueCatalogTable,
	_ResourceTypeName[944:962]:        IAMAccessKey,
	_ResourceTypeLowerName[944:962]:   IAMAccessKey,
	_ResourceTypeName[962:983]:        IAMAccountAlias,
	_ResourceTypeLowerName[962:983]:   IAMAccountAlias,
	_ResourceTypeName[983:1014]:       IAMAccountPasswordPolicy,
	_ResourceTypeLowerName[983:1014]:  IAMAccountPasswordPolicy,
	_ResourceTypeName[1014:1027]:      IAMGroup,
	_ResourceTypeLowerName[1014:1027]: IAMGroup,
	_ResourceTypeName[1027:1051]:      IAMGroupMembership,
	_ResourceTypeLowerName[1027:1051]: IAMGroupMembership,
	_ResourceTypeName[1051:1071]:      IAMGroupPolicy,
	_ResourceTypeLowerName[1051:1071]: IAMGroupPolicy,
	_ResourceTypeName[1071:1102]:      IAMGroupPolicyAttachment,
	_ResourceTypeLowerName[1071:1102]: IAMGroupPolicyAttachment,
	_ResourceTypeName[1102:1126]:      IAMInstanceProfile,
	_ResourceTypeLowerName[1102:1126]: IAMInstanceProfile,
	_ResourceTypeName[1126:1157]:      IAMOpenidConnectProvider,
	_ResourceTypeLowerName[1126:1157]: IAMOpenidConnectProvider,
	_ResourceTypeName[1157:1171]:      IAMPolicy,
	_ResourceTypeLowerName[1157:1171]: IAMPolicy,
	_ResourceTypeName[1171:1183]:      IAMRole,
	_ResourceTypeLowerName[1171:1183]: IAMRole,
	_ResourceTypeName[1183:1202]:      IAMRolePolicy,
	_ResourceTypeLowerName[1183:1202]: IAMRolePolicy,
	_ResourceTypeName[1202:1232]:      IAMRolePolicyAttachment,
	_ResourceTypeLowerName[1202:1232]: IAMRolePolicyAttachment,
	_ResourceTypeName[1232:1253]:      IAMSAMLProvider,
	_ResourceTypeLowerName[1232:1253]: IAMSAMLProvider,
	_ResourceTypeName[1253:1279]:      IAMServerCertificate,
	_ResourceTypeLowerName[1253:1279]: IAMServerCertificate,
	_ResourceTypeName[1279:1291]:      IAMUser,
	_ResourceTypeLowerName[1279:1291]: IAMUser,
	_ResourceTypeName[1291:1320]:      IAMUserGroupMembership,
	_ResourceTypeLowerName[1291:1320]: IAMUserGroupMembership,
	_ResourceTypeName[1320:1339]:      IAMUserPolicy,
	_ResourceTypeLowerName[1320:1339]: IAMUserPolicy,
	_ResourceTypeName[1339:1369]:      IAMUserPolicyAttachment,
	_ResourceTypeLowerName[1339:1369]: IAMUserPolicyAttachment,
	_ResourceTypeName[1369:1389]:      IAMUserSSHKey,
	_ResourceTypeLowerName[1369:1389]: IAMUserSSHKey,
	_ResourceTypeName[1389:1409]:      InternetGateway,
	_ResourceTypeLowerName[1389:1409]: InternetGateway,
	_ResourceTypeName[1409:1421]:      KeyPair,
	_ResourceTypeLowerName[1409:1421]: KeyPair,
	_ResourceTypeName[1421:1439]:      KinesisStream,
	_ResourceTypeLowerName[1421:1439]: KinesisStream,
	_ResourceTypeName[1439:1458]:      LambdaFunction,
	_ResourceTypeLowerName[1439:1458]: LambdaFunction,
	_ResourceTypeName[1458:1482]:      LaunchConfiguration,
	_ResourceTypeLowerName[1458:1482]: LaunchConfiguration,
	_ResourceTypeName[1482:1501]:      LaunchTemplate,
	_ResourceTypeLowerName[1482:1501]: LaunchTemplate,
	_ResourceTypeName[1501:1507]:      LB,
	_ResourceTypeLowerName[1501:1507]: LB,
	_ResourceTypeName[1507:1538]:      LBCookieStickinessPolicy,
	_ResourceTypeLowerName[1507:1538]: LBCookieStickinessPolicy,
	_ResourceTypeName[1538:1553]:      LBListener,
	_ResourceTypeLowerName[1538:1553]: LBListener,
	_ResourceTypeName[1553:1580]:      LBListenerCertificate,
	_ResourceTypeLowerName[1553:1580]: LBListenerCertificate,
	_ResourceTypeName[1580:1600]:      LBListenerRule,
	_ResourceTypeLowerName[1580:1600]: LBListenerRule,
	_ResourceTypeName[1600:1619]:      LBTargetGroup,
	_ResourceTypeLowerName[1600:1619]: LBTargetGroup,
	_ResourceTypeName[1619:1649]:      LBTargetGroupAttachment,
	_ResourceTypeLowerName[1619:1649]: LBTargetGroupAttachment,
	_ResourceTypeName[1649:1671]:      LightsailInstance,
	_ResourceTypeLowerName[1649:1671]: LightsailInstance,
	_ResourceTypeName[1671:1696]:      MediaStoreContainer,
	_ResourceTypeLowerName[1671:1696]: MediaStoreContainer,
	_ResourceTypeName[1696:1709]:      MQBroker,
	_ResourceTypeLowerName[1696:1709]: MQBroker,
	_ResourceTypeName[1709:1724]:      NatGateway,
	_ResourceTypeLowerName[1709:1724]: NatGateway,
	_ResourceTypeName[1724:1743]:      NeptuneCluster,
	_ResourceTypeLowerName[1724:1743]: NeptuneCluster,
	_ResourceTypeName[1743:1758]:      RDSCluster,
	_ResourceTypeLowerName[1743:1758]: RDSCluster,
	_ResourceTypeName[1758:1780]:      RDSGlobalCluster,
	_ResourceTypeLowerName[1758:1780]: RDSGlobalCluster,
	_ResourceTypeName[1780:1800]:      RedshiftCluster,
	_ResourceTypeLowerName[1780:1800]: RedshiftCluster,
	_ResourceTypeName[1800:1826]:      Route53DelegationSet,
	_ResourceTypeLowerName[1800:1826]: Route53DelegationSet,
	_ResourceTypeName[1826:1850]:      Route53HealthCheck,
	_ResourceTypeLowerName[1826:1850]: Route53HealthCheck,
	_ResourceTypeName[1850:1871]:      Route53QueryLog,
	_ResourceTypeLowerName[1850:1871]: Route53QueryLog,
	_ResourceTypeName[1871:1889]:      Route53Record,
	_ResourceTypeLowerName[1871:1889]: Route53Record,
	_ResourceTypeName[1889:1918]:      Route53ResolverEndpoint,
	_ResourceTypeLowerName[1889:1918]: Route53ResolverEndpoint,
	_ResourceTypeName[1918:1955]:      Route53ResolverRuleAssociation,
	_ResourceTypeLowerName[1918:1955]: Route53ResolverRuleAssociation,
	_ResourceTypeName[1955:1971]:      Route53Zone,
	_ResourceTypeLowerName[1955:1971]: Route53Zone,
	_ResourceTypeName[1971:1999]:      Route53ZoneAssociation,
	_ResourceTypeLowerName[1971:1999]: Route53ZoneAssociation,
	_ResourceTypeName[1999:2012]:      S3Bucket,
	_ResourceTypeLowerName[1999:2012]: S3Bucket,
	_ResourceTypeName[2012:2030]:      SecurityGroup,
	_ResourceTypeLowerName[2012:2030]: SecurityGroup,
	_ResourceTypeName[2030:2061]:      SESActiveReceiptRuleSet,
	_ResourceTypeLowerName[2030:2061]: SESActiveReceiptRuleSet,
	_ResourceTypeName[2061:2086]:      SESConfigurationSet,
	_ResourceTypeLowerName[2061:2086]: SESConfigurationSet,
	_ResourceTypeName[2086:2105]:      SESDomainDKIM,
	_ResourceTypeLowerName[2086:2105]: SESDomainDKIM,
	_ResourceTypeName[2105:2128]:      SESDomainIdentity,
	_ResourceTypeLowerName[2105:2128]: SESDomainIdentity,
	_ResourceTypeName[2128:2152]:      SESDomainMailFrom,
	_ResourceTypeLowerName[2128:2152]: SESDomainMailFrom,
	_ResourceTypeName[2152:2187]:      SESIdentityNotificationTopic,
	_ResourceTypeLowerName[2152:2187]: SESIdentityNotificationTopic,
	_ResourceTypeName[2187:2209]:      SESReceiptFilter,
	_ResourceTypeLowerName[2187:2209]: SESReceiptFilter,
	_ResourceTypeName[2209:2229]:      SESReceiptRule,
	_ResourceTypeLowerName[2209:2229]: SESReceiptRule,
	_ResourceTypeName[2229:2253]:      SESReceiptRuleSet,
	_ResourceTypeLowerName[2229:2253]: SESReceiptRuleSet,
	_ResourceTypeName[2253:2269]:      SESTemplate,
	_ResourceTypeLowerName[2253:2269]: SESTemplate,
	_ResourceTypeName[2269:2282]:      SQSQueue,
	_ResourceTypeLowerName[2269:2282]: SQSQueue,
	_ResourceTypeName[2282:2308]:      StoragegatewayGateway,
	_ResourceTypeLowerName[2282:2308]: StoragegatewayGateway,
	_ResourceTypeName[2308:2318]:      Subnet,
	_ResourceTypeLowerName[2308:2318]: Subnet,
	_ResourceTypeName[2318:2339]:      VolumeAttachment,
	_ResourceTypeLowerName[2318:2339]: VolumeAttachment,
	_ResourceTypeName[2339:2346]:      VPC,
	_ResourceTypeLowerName[2339:2346]: VPC,
	_ResourceTypeName[2346:2372]:      VPCPeeringConnection,
	_ResourceTypeLowerName[2346:2372]: VPCPeeringConnection,
	_ResourceTypeName[2372:2387]:      VPNGateway,
	_ResourceTypeLowerName[2372:2387]: VPNGateway,
}

var _ResourceTypeNames = []string{
	_ResourceTypeName[0:12],
	_ResourceTypeName[12:19],
	_ResourceTypeName[19:35],
	_ResourceTypeName[35:63],
	_ResourceTypeName[63:84],
	_ResourceTypeName[84:104],
	_ResourceTypeName[104:135],
	_ResourceTypeName[135:161],
	_ResourceTypeName[161:185],
	_ResourceTypeName[185:209],
	_ResourceTypeName[209:230],
	_ResourceTypeName[230:250],
	_ResourceTypeName[250:271],
	_ResourceTypeName[271:293],
	_ResourceTypeName[293:317],
	_ResourceTypeName[317:344],
	_ResourceTypeName[344:381],
	_ResourceTypeName[381:406],
	_ResourceTypeName[406:433],
	_ResourceTypeName[433:448],
	_ResourceTypeName[448:463],
	_ResourceTypeName[463:485],
	_ResourceTypeName[485:504],
	_ResourceTypeName[504:535],
	_ResourceTypeName[535:563],
	_ResourceTypeName[563:577],
	_ResourceTypeName[577:602],
	_ResourceTypeName[602:620],
	_ResourceTypeName[620:634],
	_ResourceTypeName[634:649],
	_ResourceTypeName[649:664],
	_ResourceTypeName[664:683],
	_ResourceTypeName[683:690],
	_ResourceTypeName[690:705],
	_ResourceTypeName[705:728],
	_ResourceTypeName[728:761],
	_ResourceTypeName[761:794],
	_ResourceTypeName[794:818],
	_ResourceTypeName[818:849],
	_ResourceTypeName[849:856],
	_ResourceTypeName[856:871],
	_ResourceTypeName[871:897],
	_ResourceTypeName[897:922],
	_ResourceTypeName[922:944],
	_ResourceTypeName[944:962],
	_ResourceTypeName[962:983],
	_ResourceTypeName[983:1014],
	_ResourceTypeName[1014:1027],
	_ResourceTypeName[1027:1051],
	_ResourceTypeName[1051:1071],
	_ResourceTypeName[1071:1102],
	_ResourceTypeName[1102:1126],
	_ResourceTypeName[1126:1157],
	_ResourceTypeName[1157:1171],
	_ResourceTypeName[1171:1183],
	_ResourceTypeName[1183:1202],
	_ResourceTypeName[1202:1232],
	_ResourceTypeName[1232:1253],
	_ResourceTypeName[1253:1279],
	_ResourceTypeName[1279:1291],
	_ResourceTypeName[1291:1320],
	_ResourceTypeName[1320:1339],
	_ResourceTypeName[1339:1369],
	_ResourceTypeName[1369:1389],
	_ResourceTypeName[1389:1409],
	_ResourceTypeName[1409:1421],
	_ResourceTypeName[1421:1439],
	_ResourceTypeName[1439:1458],
	_ResourceTypeName[1458:1482],
	_ResourceTypeName[1482:1501],
	_ResourceTypeName[1501:1507],
	_ResourceTypeName[1507:1538],
	_ResourceTypeName[1538:1553],
	_ResourceTypeName[1553:1580],
	_ResourceTypeName[1580:1600],
	_ResourceTypeName[1600:1619],
	_ResourceTypeName[1619:1649],
	_ResourceTypeName[1649:1671],
	_ResourceTypeName[1671:1696],
	_ResourceTypeName[1696:1709],
	_ResourceTypeName[1709:1724],
	_ResourceTypeName[1724:1743],
	_ResourceTypeName[1743:1758],
	_ResourceTypeName[1758:1780],
	_ResourceTypeName[1780:1800],
	_ResourceTypeName[1800:1826],
	_ResourceTypeName[1826:1850],
	_ResourceTypeName[1850:1871],
	_ResourceTypeName[1871:1889],
	_ResourceTypeName[1889:1918],
	_ResourceTypeName[1918:1955],
	_ResourceTypeName[1955:1971],
	_ResourceTypeName[1971:1999],
	_ResourceTypeName[1999:2012],
	_ResourceTypeName[2012:2030],
	_ResourceTypeName[2030:2061],
	_ResourceTypeName[2061:2086],
	_ResourceTypeName[2086:2105],
	_ResourceTypeName[2105:2128],
	_ResourceTypeName[2128:2152],
	_ResourceTypeName[2152:2187],
	_ResourceTypeName[2187:2209],
	_ResourceTypeName[2209:2229],
	_ResourceTypeName[2229:2253],
	_ResourceTypeName[2253:2269],
	_ResourceTypeName[2269:2282],
	_ResourceTypeName[2282:2308],
	_ResourceTypeName[2308:2318],
	_ResourceTypeName[2318:2339],
	_ResourceTypeName[2339:2346],
	_ResourceTypeName[2346:2372],
	_ResourceTypeName[2372:2387],
}

// ResourceTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ResourceTypeString(s string) (ResourceType, error) {
	if val, ok := _ResourceTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ResourceTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ResourceType values", s)
}

// ResourceTypeValues returns all values of the enum
func ResourceTypeValues() []ResourceType {
	return _ResourceTypeValues
}

// ResourceTypeStrings returns a slice of all String values of the enum
func ResourceTypeStrings() []string {
	strs := make([]string, len(_ResourceTypeNames))
	copy(strs, _ResourceTypeNames)
	return strs
}

// IsAResourceType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ResourceType) IsAResourceType() bool {
	for _, v := range _ResourceTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
