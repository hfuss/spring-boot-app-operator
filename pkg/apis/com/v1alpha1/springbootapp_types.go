package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SpringBootAppSpec defines the desired state of SpringBootApp
// +k8s:openapi-gen=true
type SpringBootAppSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Image string `json:"image"`
	ImageTag string `json:"imageTag,omitempty"`
	JavaToolOptions []string `json:"javaToolOptions,omitempty"`
	ActiveProfile string `json:"activeProfile,omitempty"`
	RouteEnabled bool `json:"routeEnabled,omitempty"`
	LivenessEndpoint string `json:"livenessEndpoint,omitempty"`
	ReadinessEndpoint string `json:"readinessEndpoint,omitempty"`
	Resources v1.ResourceRequirements `json:"resources,omitempty" protobuf:"bytes,8,opt,name=resources"`
	EnvFrom v1.EnvFromSource `json:"envFrom,omitempty" protobuf:"bytes,19,rep,name=envFrom"`
}

// SpringBootAppStatus defines the observed state of SpringBootApp
// +k8s:openapi-gen=true
type SpringBootAppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SpringBootApp is the Schema for the springbootapps API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=springbootapps,scope=Namespaced
type SpringBootApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpringBootAppSpec   `json:"spec,omitempty"`
	Status SpringBootAppStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SpringBootAppList contains a list of SpringBootApp
type SpringBootAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringBootApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpringBootApp{}, &SpringBootAppList{})
}
