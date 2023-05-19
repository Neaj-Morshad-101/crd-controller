package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KlusterSpec struct {
}

type Kluster struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec KlusterSpec
}

type KlusterList struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Items []Kluster
}
