// NOTE: Boilerplate only.  Ignore this file.

// Package v1beta1 contains API Schema definitions for the app v1beta1 API group
// +k8s:deepcopy-gen=package,register
// +groupName=app.zcloud.cn
package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/zdnscloud/gok8s/scheme"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "app.zcloud.cn", Version: "v1beta1"}
)

func AddToScheme(s *runtime.Scheme) {
	builder := &scheme.Builder{GroupVersion: SchemeGroupVersion}
	builder.Register(&Application{}, &ApplicationList{})
	builder.AddToScheme(s)
}
