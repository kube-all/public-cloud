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

// AliyunACK is the schema for the clusters API
type AliyunACK struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   AliyunACKSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status AliyunACKStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}
type AliyunACKSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=Kubernetes;ManagedKubernetes;Ask
	ClusterType       string            `json:"cluster_type" json:"cluster_type" protobuf:"bytes,1,opt,name=cluster_type,json=clusterType"`
	Ask               Ask               `json:"ask" yaml:"ask" protobuf:"bytes,2,opt,name=ask"`
	Kubernetes        Kubernetes        `json:"kubernetes" json:"kubernetes" protobuf:"bytes,3,opt,name=kubernetes"`
	ManagedKubernetes ManagedKubernetes `json:"managed_kubernetes" json:"managed_kubernetes" protobuf:"bytes,4,opt,name=managed_kubernetes,json=managedKubernetes"`
}
type Kubernetes struct {
}
type ManagedKubernetes struct {
}
type Ask struct {
	//集群类型。可选值为ManagedKubernetes，同时profile配置为Serverless，表示创建Serverless集群。
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=ManagedKubernetes;Ask
	ClusterType string `json:"cluster_type" json:"cluster_type" protobuf:"bytes,1,opt,name=cluster_type,json=clusterType"`
	//集群标识。参数cluster_type取值为ManagedKubernetes，同时该参数配置为Serverless，表示创建Serverless集群。
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:default=Serverless
	Profile string `json:"profile" json:"profile" protobuf:"bytes,2,opt,name=profile"`
	// +kubebuilder:validation:Enum=ack.pro.small;ack.standard
	// +kubebuilder:validation:default=ack.standard

	// +optional
	ClusterSpec string `json:"cluster_spec" json:"cluster_spec" protobuf:"bytes,3,opt,name=cluster_spec,json=clusterSpec"`
	// +optional
	// +kubebuilder:validation:Enum=1.24.3-aliyun.1;1.22.10-aliyun.1
	KubernetesVersion string `json:"kubernetes_version" json:"kubernetes_version" protobuf:"bytes,4,opt,name=kubernetes_version,json=kubernetesVersion"`
	// +optional
	PrivateZone bool `json:"private_zone" json:"private_zone" protobuf:"varint,5,opt,name=private_zone,json=privateZone"`
	//集群所在的地域ID
	// +kubebuilder:validation:Required
	RegionId string `json:"region_id" json:"region_id" protobuf:"bytes,6,opt,name=region_id,json=regionId"`
	// +optional
	EndpointPublicAccess bool `json:"endpoint_public_access" json:"endpoint_public_access" protobuf:"varint,7,opt,name=endpoint_public_access,json=endpointPublicAccess"`
	// +optional
	ServiceDiscoveryTypes []string `json:"service_discovery_types" json:"service_discovery_types" protobuf:"bytes,8,rep,name=service_discovery_types,json=serviceDiscoveryTypes"`
	// +optional
	Zoneid string `json:"zoneid" yaml:"zoneid" protobuf:"bytes,9,opt,name=zoneid"`
	// +optional
	LoggingType string `json:"logging_type" json:"logging_type" protobuf:"bytes,10,opt,name=logging_type,json=loggingType"`
	// +optional
	Tags []Tag `json:"tags" yaml:"tags" protobuf:"bytes,11,rep,name=tags"`
	// +optional
	DeletionProtection bool `json:"deletion_protection" json:"deletion_protection" protobuf:"varint,12,opt,name=deletion_protection,json=deletionProtection"`
	// +optional
	ServiceCIDR string `json:"service_cidr" json:"service_cidr" protobuf:"bytes,13,opt,name=service_cidr,json=serviceCidr"`
	// +optional
	Timezone string `json:"timezone" yaml:"timezone" protobuf:"bytes,14,opt,name=timezone"`
	// +optional
	Addons []Addon `json:"addons" yaml:"addons" protobuf:"bytes,15,rep,name=addons"`
	// +optional
	SnatEntry bool `json:"snat_entry" json:"snat_entry" protobuf:"varint,16,opt,name=snat_entry,json=snatEntry"`
	// +optional
	Vpcid string `json:"vpcid" yaml:"vpcid" protobuf:"bytes,17,opt,name=vpcid"`
	// +optional
	VswitchIds []string `json:"vswitch_ids" json:"vswitch_ids" protobuf:"bytes,18,rep,name=vswitch_ids,json=vswitchIds"`
	// +optional
	SecurityGroupId string `json:"security_group_id" json:"security_group_id" protobuf:"bytes,19,opt,name=security_group_id,json=securityGroupId"`
	// +optional
	ResourceGroupId string `json:"resourceGroupId" yaml:"resourceGroupId" protobuf:"bytes,20,opt,name=resourceGroupId"`
}
type Addon struct {
	Name     string `json:"name" json:"name" protobuf:"bytes,1,opt,name=name"`
	Config   string `json:"config" yaml:"config" protobuf:"bytes,2,opt,name=config"`
	Disabled bool   `json:"disabled" yaml:"disabled" protobuf:"varint,3,opt,name=disabled"`
}
type Tag struct {
	Key   string `json:"key" yaml:"key" protobuf:"bytes,1,opt,name=key"`
	Value string `json:"value" yaml:"value" protobuf:"bytes,2,opt,name=value"`
}
type AliyunACKStatus struct {
	// cluster id
	// +optional
	ClusterID string `json:"cluster_id" json:"cluster_id" protobuf:"bytes,1,opt,name=cluster_id,json=clusterId"`
	// cluster kubeconfig
	// +optional
	Config string `json:"config" json:"config" protobuf:"bytes,2,opt,name=config"`
	//kubeconfig expire time, format:RFC3339 UTC
	// +optional
	Expiration metav1.Time `json:"expiration" json:"expiration" protobuf:"bytes,3,opt,name=expiration"`
	// all cluster tasks
	// +optional
	Tasks []AliyunTask `json:"tasks" json:"tasks" protobuf:"bytes,4,rep,name=tasks"`
}
type AliyunTask struct {
	// request time
	// +kubebuilder:validation:Required
	Time metav1.Time `json:"time" yaml:"time" protobuf:"bytes,1,opt,name=time"`
	//aliyun response request_id
	// +kubebuilder:validation:Required
	RequestID string `json:"request_id" json:"request_id" protobuf:"bytes,2,opt,name=request_id,json=requestId"`
	//aliyun response task_id
	// +kubebuilder:validation:Required
	TaskID string `json:"task_id" json:"task_id" protobuf:"bytes,3,opt,name=task_id,json=taskId"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AliyunACKList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []AliyunACK `json:"items" protobuf:"bytes,2,rep,name=items"`
}
