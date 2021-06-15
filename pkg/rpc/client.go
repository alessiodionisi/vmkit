// Virtual Machine manager that supports QEMU and Apple virtualization framework on macOS
// Copyright (C) 2021 VMKit Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
