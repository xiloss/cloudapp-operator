/*
Copyright 2024.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DbAppSpec defines the desired state of DbApp
type DbAppSpec struct {
	// Important: Run "make" to regenerate code after modifying this file
	// DbAppStsReplicas is the number of replicas of a db server stateful set
	DbAppStsReplicas int32 `json:"dbAppStsReplicas,omitempty"`
	// DbAppDeploymentImage is the container image field of DbApp.
	DbAppDeploymentImage string `json:"dbServerImage,omitempty"`
	// DbAppDeploymentExpose is the control to expose a DbApp deployment as a service
	DbAppDeploymentExpose bool `json:"dbServerExpose,omitempty"`
	// DbAppServicePorts is the list of maps of DbApp service ports to be exposed
	DbAppServicePorts []corev1.ServicePort `json:"dbServerServicePorts,omitempty"`
	// DbAppServiceType is the type of service used to expose the stateful set
	DbAppServiceType corev1.ServiceType `json:"dbServerServiceType,omitempty"`
}

// DbAppStatus defines the observed state of DbApp
type DbAppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DbApp is the Schema for the dbapps API
type DbApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DbAppSpec   `json:"spec,omitempty"`
	Status DbAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbAppList contains a list of DbApp
type DbAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DbApp{}, &DbAppList{})
}
