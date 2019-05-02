// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerConfiguration defines the configuration for the AWS provider.
type ControllerConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	// MachineImages is the list of machine images that are understood by the controller. It maps
	// logical names and versions to AWS-specific identifiers, i.e. AMIs.
	MachineImages []MachineImage `json:"machineImages,omitempty"`
}

// MachineImage is a mapping from logical names and versions to AWS-specific identifiers, i.e. AMIs.
type MachineImage struct {
	// Name is the logical name of the machine image.
	Name string `json:"name"`
	// Version is the logical version of the machine image.
	Version string `json:"version"`
	// Regions is a mapping to the correct AMI for the machine image in the supported regions.
	Regions []RegionAMIMapping `json:"regions"`
}

// RegionAMIMapping is a mapping to the correct AMI for the machine image in the given region.
type RegionAMIMapping struct {
	// Name is the name of the region.
	Name string `json:"name"`
	// AMI is the AMI for the machine image.
	AMI string `json:"ami"`
}
