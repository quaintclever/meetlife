/*
Copyright 2022.

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

package v1

import (
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ReplicasDemoSpec defines the desired state of ReplicasDemo
type ReplicasDemoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// do quaint spec define
	// BatchSize 用来控制 replicas 启动时每次分批的数量
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=5
	BatchSize int32 `json:"batchSize,omitempty"`

	// DeploymentSpec  k8s 自带 deploymentSpec
	DeploymentSpec v1.DeploymentSpec `json:"deploymentSpec,omitempty"`
}

// ReplicasDemoStatus defines the observed state of ReplicasDemo
type ReplicasDemoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ReplicasDemo is the Schema for the replicasdemoes API
type ReplicasDemo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReplicasDemoSpec   `json:"spec,omitempty"`
	Status ReplicasDemoStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ReplicasDemoList contains a list of ReplicasDemo
type ReplicasDemoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicasDemo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ReplicasDemo{}, &ReplicasDemoList{})
}
