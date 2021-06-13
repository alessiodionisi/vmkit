package main

import (
	"fmt"

	"github.com/adnsio/vmkit/pkg/rpc"
	"github.com/adnsio/vmkit/pkg/util"
)

func NewRPCClient(address string) (*rpc.Client, error) {
	parsedAddress, err := util.NewNetworkAddress(address)
	if err != nil {
		return nil, err
	}

	client := rpc.NewClient(&rpc.NewClientOptions{
		Network: parsedAddress.Network,
		Address: parsedAddress.Address,
	})

	if err := client.Dial(); err != nil {
		return nil, fmt.Errorf("%w. is the vmkit daemon running?", err)
	}

	return client, nil
}
