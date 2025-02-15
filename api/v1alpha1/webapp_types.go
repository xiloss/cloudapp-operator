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

// WebAppSpec defines the desired state of WebApp
type WebAppSpec struct {
	// Important: Run "make" to regenerate code after modifying this file
	// WebAppDeploymentReplicas is the number of replicas of a webserver deployment
	WebAppDeploymentReplicas int32 `json:"webAppReplicas,omitempty"`
	// WebAppDeploymentImage is the container image field of WebApp.
	WebAppDeploymentImage string `json:"webAppImage,omitempty"`
	// WebAppDeploymentExpose is the control to expose a WebApp deployment as a service
	WebAppDeploymentExpose bool `json:"webAppDeploymentExpose,omitempty"`
	// WebAppServicePorts is the list of maps of WebApp service ports to be exposed
	WebAppServicePorts []corev1.ServicePort `json:"webAppServicePorts,omitempty"`
	// WebAppServiceType is the type of service used to expose the deployment
	WebAppServiceType corev1.ServiceType `json:"webAppServiceType,omitempty"`
}

// WebAppStatus defines the observed state of WebApp
type WebAppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// WebApp is the Schema for the webapps API
type WebApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebAppSpec   `json:"spec,omitempty"`
	Status WebAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebAppList contains a list of WebApp
type WebAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WebApp{}, &WebAppList{})
}
