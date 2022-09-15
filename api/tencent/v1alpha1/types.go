/*
Copyright 2022 The kubeall.com Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// TencentCluster is the schema for the clusters API
type TencentCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   TencentClusterSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status TencentClusterStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}
type TencentClusterSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=Kubernetes;ManagedKubernetes;Ask
	ClusterType string `json:"cluster_type" json:"cluster_type" protobuf:"bytes,1,opt,name=cluster_type,json=clusterType"`
}

type TencentClusterStatus struct {
	// cluster id
	// +optional
	ClusterID string `json:"cluster_id" json:"cluster_id" protobuf:"bytes,1,opt,name=cluster_id,json=clusterId"`
	// cluster kubeconfig
	// +optional
	Config string `json:"config" json:"config" protobuf:"bytes,2,opt,name=config"`
	//kubeconfig expire time, format:RFC3339 UTC
	// +optional
	Expiration metav1.Time `json:"expiration" json:"expiration" protobuf:"bytes,3,opt,name=expiration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TencentClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []TencentCluster `json:"items" protobuf:"bytes,2,rep,name=items"`
}
