package hypervisor

import "os/exec"

type AVFVM struct {
}

func (a *AVFVM) IsSupported() bool {
	path, err := exec.LookPath("avfvm")
	if err != nil {
		return false
	}

	if path == "" {
		return false
	}

	return true
}

func NewAVFVM() (Hypervisor, error) {
	return &AVFVM{}, nil
}
