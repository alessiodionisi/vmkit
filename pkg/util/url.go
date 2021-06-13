package util

import "net/url"

type NetworkAddress struct {
	Network string
	Address string
}

func NewNetworkAddress(address string) (*NetworkAddress, error) {
	parsedAddress, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	if parsedAddress.Scheme == "unix" {
		return &NetworkAddress{
			Network: parsedAddress.Scheme,
			Address: parsedAddress.Path,
		}, nil
	} else {
		return &NetworkAddress{
			Network: parsedAddress.Scheme,
			Address: parsedAddress.Host,
		}, nil
	}
}
