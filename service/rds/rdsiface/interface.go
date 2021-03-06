package rdsiface

import (
	"github.com/awslabs/aws-sdk-go/service/rds"
)

type RDSAPI interface {
	AddSourceIdentifierToSubscription(*rds.AddSourceIdentifierToSubscriptionInput) (*rds.AddSourceIdentifierToSubscriptionOutput, error)

	AddTagsToResource(*rds.AddTagsToResourceInput) (*rds.AddTagsToResourceOutput, error)

	ApplyPendingMaintenanceAction(*rds.ApplyPendingMaintenanceActionInput) (*rds.ApplyPendingMaintenanceActionOutput, error)

	AuthorizeDBSecurityGroupIngress(*rds.AuthorizeDBSecurityGroupIngressInput) (*rds.AuthorizeDBSecurityGroupIngressOutput, error)

	CopyDBParameterGroup(*rds.CopyDBParameterGroupInput) (*rds.CopyDBParameterGroupOutput, error)

	CopyDBSnapshot(*rds.CopyDBSnapshotInput) (*rds.CopyDBSnapshotOutput, error)

	CopyOptionGroup(*rds.CopyOptionGroupInput) (*rds.CopyOptionGroupOutput, error)

	CreateDBInstance(*rds.CreateDBInstanceInput) (*rds.CreateDBInstanceOutput, error)

	CreateDBInstanceReadReplica(*rds.CreateDBInstanceReadReplicaInput) (*rds.CreateDBInstanceReadReplicaOutput, error)

	CreateDBParameterGroup(*rds.CreateDBParameterGroupInput) (*rds.CreateDBParameterGroupOutput, error)

	CreateDBSecurityGroup(*rds.CreateDBSecurityGroupInput) (*rds.CreateDBSecurityGroupOutput, error)

	CreateDBSnapshot(*rds.CreateDBSnapshotInput) (*rds.CreateDBSnapshotOutput, error)

	CreateDBSubnetGroup(*rds.CreateDBSubnetGroupInput) (*rds.CreateDBSubnetGroupOutput, error)

	CreateEventSubscription(*rds.CreateEventSubscriptionInput) (*rds.CreateEventSubscriptionOutput, error)

	CreateOptionGroup(*rds.CreateOptionGroupInput) (*rds.CreateOptionGroupOutput, error)

	DeleteDBInstance(*rds.DeleteDBInstanceInput) (*rds.DeleteDBInstanceOutput, error)

	DeleteDBParameterGroup(*rds.DeleteDBParameterGroupInput) (*rds.DeleteDBParameterGroupOutput, error)

	DeleteDBSecurityGroup(*rds.DeleteDBSecurityGroupInput) (*rds.DeleteDBSecurityGroupOutput, error)

	DeleteDBSnapshot(*rds.DeleteDBSnapshotInput) (*rds.DeleteDBSnapshotOutput, error)

	DeleteDBSubnetGroup(*rds.DeleteDBSubnetGroupInput) (*rds.DeleteDBSubnetGroupOutput, error)

	DeleteEventSubscription(*rds.DeleteEventSubscriptionInput) (*rds.DeleteEventSubscriptionOutput, error)

	DeleteOptionGroup(*rds.DeleteOptionGroupInput) (*rds.DeleteOptionGroupOutput, error)

	DescribeDBEngineVersions(*rds.DescribeDBEngineVersionsInput) (*rds.DescribeDBEngineVersionsOutput, error)

	DescribeDBInstances(*rds.DescribeDBInstancesInput) (*rds.DescribeDBInstancesOutput, error)

	DescribeDBLogFiles(*rds.DescribeDBLogFilesInput) (*rds.DescribeDBLogFilesOutput, error)

	DescribeDBParameterGroups(*rds.DescribeDBParameterGroupsInput) (*rds.DescribeDBParameterGroupsOutput, error)

	DescribeDBParameters(*rds.DescribeDBParametersInput) (*rds.DescribeDBParametersOutput, error)

	DescribeDBSecurityGroups(*rds.DescribeDBSecurityGroupsInput) (*rds.DescribeDBSecurityGroupsOutput, error)

	DescribeDBSnapshots(*rds.DescribeDBSnapshotsInput) (*rds.DescribeDBSnapshotsOutput, error)

	DescribeDBSubnetGroups(*rds.DescribeDBSubnetGroupsInput) (*rds.DescribeDBSubnetGroupsOutput, error)

	DescribeEngineDefaultParameters(*rds.DescribeEngineDefaultParametersInput) (*rds.DescribeEngineDefaultParametersOutput, error)

	DescribeEventCategories(*rds.DescribeEventCategoriesInput) (*rds.DescribeEventCategoriesOutput, error)

	DescribeEventSubscriptions(*rds.DescribeEventSubscriptionsInput) (*rds.DescribeEventSubscriptionsOutput, error)

	DescribeEvents(*rds.DescribeEventsInput) (*rds.DescribeEventsOutput, error)

	DescribeOptionGroupOptions(*rds.DescribeOptionGroupOptionsInput) (*rds.DescribeOptionGroupOptionsOutput, error)

	DescribeOptionGroups(*rds.DescribeOptionGroupsInput) (*rds.DescribeOptionGroupsOutput, error)

	DescribeOrderableDBInstanceOptions(*rds.DescribeOrderableDBInstanceOptionsInput) (*rds.DescribeOrderableDBInstanceOptionsOutput, error)

	DescribePendingMaintenanceActions(*rds.DescribePendingMaintenanceActionsInput) (*rds.DescribePendingMaintenanceActionsOutput, error)

	DescribeReservedDBInstances(*rds.DescribeReservedDBInstancesInput) (*rds.DescribeReservedDBInstancesOutput, error)

	DescribeReservedDBInstancesOfferings(*rds.DescribeReservedDBInstancesOfferingsInput) (*rds.DescribeReservedDBInstancesOfferingsOutput, error)

	DownloadDBLogFilePortion(*rds.DownloadDBLogFilePortionInput) (*rds.DownloadDBLogFilePortionOutput, error)

	ListTagsForResource(*rds.ListTagsForResourceInput) (*rds.ListTagsForResourceOutput, error)

	ModifyDBInstance(*rds.ModifyDBInstanceInput) (*rds.ModifyDBInstanceOutput, error)

	ModifyDBParameterGroup(*rds.ModifyDBParameterGroupInput) (*rds.DBParameterGroupNameMessage, error)

	ModifyDBSubnetGroup(*rds.ModifyDBSubnetGroupInput) (*rds.ModifyDBSubnetGroupOutput, error)

	ModifyEventSubscription(*rds.ModifyEventSubscriptionInput) (*rds.ModifyEventSubscriptionOutput, error)

	ModifyOptionGroup(*rds.ModifyOptionGroupInput) (*rds.ModifyOptionGroupOutput, error)

	PromoteReadReplica(*rds.PromoteReadReplicaInput) (*rds.PromoteReadReplicaOutput, error)

	PurchaseReservedDBInstancesOffering(*rds.PurchaseReservedDBInstancesOfferingInput) (*rds.PurchaseReservedDBInstancesOfferingOutput, error)

	RebootDBInstance(*rds.RebootDBInstanceInput) (*rds.RebootDBInstanceOutput, error)

	RemoveSourceIdentifierFromSubscription(*rds.RemoveSourceIdentifierFromSubscriptionInput) (*rds.RemoveSourceIdentifierFromSubscriptionOutput, error)

	RemoveTagsFromResource(*rds.RemoveTagsFromResourceInput) (*rds.RemoveTagsFromResourceOutput, error)

	ResetDBParameterGroup(*rds.ResetDBParameterGroupInput) (*rds.DBParameterGroupNameMessage, error)

	RestoreDBInstanceFromDBSnapshot(*rds.RestoreDBInstanceFromDBSnapshotInput) (*rds.RestoreDBInstanceFromDBSnapshotOutput, error)

	RestoreDBInstanceToPointInTime(*rds.RestoreDBInstanceToPointInTimeInput) (*rds.RestoreDBInstanceToPointInTimeOutput, error)

	RevokeDBSecurityGroupIngress(*rds.RevokeDBSecurityGroupIngressInput) (*rds.RevokeDBSecurityGroupIngressOutput, error)
}