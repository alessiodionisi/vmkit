package hypervisor

type Hypervisor interface {
	IsSupported() bool
}
