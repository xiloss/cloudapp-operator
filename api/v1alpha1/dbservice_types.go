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

// DbServiceSpec defines the desired state of DbService
type DbServiceSpec struct {
	// Important: Run "make" to regenerate code after modifying this file
	// DbServiceStsReplicas is the number of replicas of a db server stateful set
	DbServiceStsReplicas int32 `json:"dbServiceStsReplicas,omitempty"`
	// DbServiceStsImage is the container image field of DbService.
	DbServiceStsImage string `json:"dbServiceStsImage,omitempty"`
	// DbServiceStsExpose is the control to expose a DbService stateful set as a service
	DbServiceStsExpose bool `json:"dbServiceStsExpose,omitempty"`
	// DbServiceSvcPorts is the list of maps of DbService service ports to be exposed
	DbServiceSvcPorts []corev1.ServicePort `json:"dbServiceSvcPorts,omitempty"`
	// DbServiceSvcType is the type of service used to expose the stateful set
	DbServiceSvcType corev1.ServiceType `json:"dbServiceSvcType,omitempty"`
}

// DbServiceStatus defines the observed state of DbService
type DbServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DbService is the Schema for the dbservices API
type DbService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DbServiceSpec   `json:"spec,omitempty"`
	Status DbServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbServiceList contains a list of DbService
type DbServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DbService{}, &DbServiceList{})
}
