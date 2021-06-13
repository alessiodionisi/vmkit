package rpc

import (
	"net/rpc"
)

type NewClientOptions struct {
	Network string
	Address string
}

type Client struct {
	network string
	address string
	client  *rpc.Client
}

func (c *Client) Dial() error {
	cli, err := rpc.Dial(c.network, c.address)
	if err != nil {
		return err
	}

	c.client = cli

	return nil
}

func (c *Client) ListVirtualMachines() ([]*VirtualMachine, error) {
	var res []*VirtualMachine
	if err := c.client.Call("VirtualMachine.List", new(interface{}), &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) StartVirtualMachine(name string) error {
	if err := c.client.Call("VirtualMachine.Start", &StartVirtualMachineOptions{
		Name: name,
	}, new(interface{})); err != nil {
		return err
	}

	return nil
}

func (c *Client) VirtualMachineLogs(name string) (string, error) {
	var res string
	if err := c.client.Call("VirtualMachine.Logs", &VirtualMachineLogsOptions{
		Name: name,
	}, &res); err != nil {
		return "", err
	}

	return res, nil
}

func NewClient(opts *NewClientOptions) *Client {
	return &Client{
		network: opts.Network,
		address: opts.Address,
	}
}
