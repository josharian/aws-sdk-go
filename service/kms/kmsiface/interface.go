package kmsiface

import (
	"github.com/awslabs/aws-sdk-go/service/kms"
)

type KMSAPI interface {
	CreateAlias(*kms.CreateAliasInput) (*kms.CreateAliasOutput, error)

	CreateGrant(*kms.CreateGrantInput) (*kms.CreateGrantOutput, error)

	CreateKey(*kms.CreateKeyInput) (*kms.CreateKeyOutput, error)

	Decrypt(*kms.DecryptInput) (*kms.DecryptOutput, error)

	DeleteAlias(*kms.DeleteAliasInput) (*kms.DeleteAliasOutput, error)

	DescribeKey(*kms.DescribeKeyInput) (*kms.DescribeKeyOutput, error)

	DisableKey(*kms.DisableKeyInput) (*kms.DisableKeyOutput, error)

	DisableKeyRotation(*kms.DisableKeyRotationInput) (*kms.DisableKeyRotationOutput, error)

	EnableKey(*kms.EnableKeyInput) (*kms.EnableKeyOutput, error)

	EnableKeyRotation(*kms.EnableKeyRotationInput) (*kms.EnableKeyRotationOutput, error)

	Encrypt(*kms.EncryptInput) (*kms.EncryptOutput, error)

	GenerateDataKey(*kms.GenerateDataKeyInput) (*kms.GenerateDataKeyOutput, error)

	GenerateDataKeyWithoutPlaintext(*kms.GenerateDataKeyWithoutPlaintextInput) (*kms.GenerateDataKeyWithoutPlaintextOutput, error)

	GenerateRandom(*kms.GenerateRandomInput) (*kms.GenerateRandomOutput, error)

	GetKeyPolicy(*kms.GetKeyPolicyInput) (*kms.GetKeyPolicyOutput, error)

	GetKeyRotationStatus(*kms.GetKeyRotationStatusInput) (*kms.GetKeyRotationStatusOutput, error)

	ListAliases(*kms.ListAliasesInput) (*kms.ListAliasesOutput, error)

	ListGrants(*kms.ListGrantsInput) (*kms.ListGrantsOutput, error)

	ListKeyPolicies(*kms.ListKeyPoliciesInput) (*kms.ListKeyPoliciesOutput, error)

	ListKeys(*kms.ListKeysInput) (*kms.ListKeysOutput, error)

	PutKeyPolicy(*kms.PutKeyPolicyInput) (*kms.PutKeyPolicyOutput, error)

	ReEncrypt(*kms.ReEncryptInput) (*kms.ReEncryptOutput, error)

	RetireGrant(*kms.RetireGrantInput) (*kms.RetireGrantOutput, error)

	RevokeGrant(*kms.RevokeGrantInput) (*kms.RevokeGrantOutput, error)

	UpdateKeyDescription(*kms.UpdateKeyDescriptionInput) (*kms.UpdateKeyDescriptionOutput, error)
}