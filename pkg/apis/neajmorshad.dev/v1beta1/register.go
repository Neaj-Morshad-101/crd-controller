package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVerson = schema.GroupVersion{
	Group:   neajmorshad.dev,
	Version: v1beta1,
}

var (
	SchemeBuilder runtime.SchemeBuilder
)

func init() {
	SchemeBuilder.Register(addKnownTypes)
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVerson, &Kluster{}, &KlusterList{})
	metav1.AddToGroupVersion(scheme, SchemeGroupVerson)
	return nil
}
