/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha3

// ANCHOR: CommonConditions

// Common ConditionTypes used by Cluster API objects.
const (
	// ReadyCondition defines the Ready condition type that summarizes the operational state of a Cluster API object.
	ReadyCondition ConditionType = "Ready"
)

// ANCHOR_END: CommonConditions

// Conditions and condition Reasons for the Machine object

// MachineSummaryConditionsCount defines the total number of conditions on the Machine object.
const MachineSummaryConditionsCount = 2

const (
	// BootstrapReadyCondition reports a summary of current status of the bootstrap object defined for this machine.
	// This condition is mirrored from the Ready condition in the bootstrap ref object, and
	// the absence of this condition might signal problems in the reconcile external loops or the fact that
	// the bootstrap provider does not not implements the Ready condition yet.
	BootstrapReadyCondition ConditionType = "BootstrapReady"

	// WaitingForDataSecretFallbackReason (Severity=Info) documents a machine waiting for the bootstrap data secret
	// to be available.
	// NOTE: This reason is used only as a fallback when the bootstrap object is not reporting its own ready condition.
	WaitingForDataSecretFallbackReason = "WaitingForDataSecret"
)

const (
	// InfrastructureReadyCondition reports a summary of current status of the infrastructure object defined for this machine.
	// This condition is mirrored from the Ready condition in the infrastructure ref object, and
	// the absence of this condition might signal problems in the reconcile external loops or the fact that
	// the infrastructure provider does not not implements the Ready condition yet.
	InfrastructureReadyCondition ConditionType = "InfrastructureReady"

	// WaitingForInfrastructureFallbackReason (Severity=Info) documents a machine waiting for the machine infrastructure
	// to be available.
	// NOTE: This reason is used only as a fallback when the infrastructure object is not reporting its own ready condition.
	WaitingForInfrastructureFallbackReason = "WaitingForInfrastructure"
)
