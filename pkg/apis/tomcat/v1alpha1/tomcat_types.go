package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TomcatSpec defines the desired state of Tomcat
// +k8s:openapi-gen=true
type TomcatSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Number of desired tomcat pods. This is a pointer to distinguish between
	// explicit zero and not specified. Defaults to 1.
	// +optional
	Replicas *int32 `json:"replicas,omitempty"`

	// Tomcat runtime image to use. Defaults to tomcat:latest.
	// +optional
	Image string `json:"image,omitempty"`

	// ImagePullPolicy overrides TomcatRuntime spec.imagePullPolicy
	// +kubebuilder:validation:Enum=Always,IfNotPresent,Never
	// +optional
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty"`

	// Port represents the open port for tomcat service. Defaults to 80.
	// +optional
	ServicePort *int32 `json:"servicePort,omitempty"`

	// WebArchiveImage is the init image that provides the source war package.
	// Default to ananwaresystems/webarchive:1.0
	// +optional
	WebArchiveImage string `json:"webArchiveImage,omitempty"`

	// DeployDirectory must match the tomcat setup directory of your Tomcat Image.
	// Default to /usr/local/tomcat/webapps
	// +optional
	DeployDirectory string `json:"deployDirectory,omitempty"`
}

// TomcatStatus defines the observed state of Tomcat
// +k8s:openapi-gen=true
type TomcatStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Total number of non-terminated pods targeted by tomcat deployment
	// This is copied over from the deployment object
	// +optional
	Replicas int32 `json:"replicas,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Tomcat is the Schema for the tomcats API
// +k8s:openapi-gen=true
type Tomcat struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TomcatSpec   `json:"spec,omitempty"`
	Status TomcatStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TomcatList contains a list of Tomcat
type TomcatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tomcat `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Tomcat{}, &TomcatList{})
}
