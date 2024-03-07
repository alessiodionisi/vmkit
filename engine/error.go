package engine

import "errors"

var (
	ErrCannotDetachRootVolume  = errors.New("cannot detach root volume")
	ErrImageNotFound           = errors.New("image not found")
	ErrInstanceAlreadyExist    = errors.New("instance already exists")
	ErrInstanceNotFound        = errors.New("instance not found")
	ErrInvalidFieldCombination = errors.New("invalid field combination")
	ErrRequiredFieldNotSet     = errors.New("required field not set")
	ErrVolumeAlreadyAttached   = errors.New("volume already attached")
	ErrVolumeAlreadyExist      = errors.New("volume already exists")
	ErrVolumeAttached          = errors.New("volume attached")
	ErrVolumeNotAttached       = errors.New("volume not attached")
	ErrVolumeNotFound          = errors.New("volume not found")
)
