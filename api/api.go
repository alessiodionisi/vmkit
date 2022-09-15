package api

import "strings"

type Metadata struct {
	Name string
}

const (
	KindDisk           string = "disk"
	KindVirtualMachine string = "virtualmachine"
)

type APIVersionAndKind struct {
	APIVersion string
	Kind       string
}

func NormalizeAPIVersion(apiVersion string) string {
	return strings.ToLower(apiVersion)
}

func NormalizeKind(kind string) string {
	return strings.ToLower(kind)
}
