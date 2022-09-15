package v1alpha1

import "github.com/adnsio/vmkit/api"

type DiskSpec struct {
}

type Disk struct {
	api.APIVersionAndKind
	api.Metadata
	Spec *DiskSpec
}
