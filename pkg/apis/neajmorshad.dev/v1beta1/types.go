package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ContainerSpec struct {
	Image string `json:"image,omitempty"`
	Port  int32  `json:"port,omitempty"`
}

// KlusterSpec defines the desired state of KlusterCRD
type KlusterSpec struct {
	Name      string        `json:"name,omitempty"`
	Replicas  *int32        `json:"replicas"`
	Container ContainerSpec `json:"container,container"`
}

type KlusterStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

type Kluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KlusterSpec   `json:"spec"`
	Status KlusterStatus `json:"status,omitempty"`
}

type KlusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kluster `json:"items"`
}
