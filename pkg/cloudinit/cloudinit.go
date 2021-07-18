// Spin up Linux VMs with QEMU and Apple virtualization framework
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

package cloudinit

import (
	"errors"
	"os"

	"github.com/diskfs/go-diskfs"
	"github.com/diskfs/go-diskfs/disk"
	"github.com/diskfs/go-diskfs/filesystem"
	"github.com/diskfs/go-diskfs/filesystem/iso9660"
)

type NewCloudInitISOOptions struct {
	MetaData      string
	Name          string
	NetworkConfig string
	UserData      string
}

func NewCloudInitISO(opts *NewCloudInitISOOptions) error {
	if err := os.Remove(opts.Name); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	dfs, err := diskfs.Create(opts.Name, 10*1024*1024, diskfs.Raw)
	if err != nil {
		return err
	}

	dfs.LogicalBlocksize = 2048

	fs, err := dfs.CreateFilesystem(disk.FilesystemSpec{
		Partition: 0,
		FSType:    filesystem.TypeISO9660,
	})
	if err != nil {
		return err
	}

	metaDataFile, err := fs.OpenFile("meta-data", os.O_CREATE|os.O_RDWR)
	if err != nil {
		return err
	}
	defer metaDataFile.Close()

	if _, err := metaDataFile.Write([]byte(opts.MetaData)); err != nil {
		return err
	}

	userDataFile, err := fs.OpenFile("user-data", os.O_CREATE|os.O_RDWR)
	if err != nil {
		return err
	}
	defer userDataFile.Close()

	if _, err := userDataFile.Write([]byte(opts.UserData)); err != nil {
		return err
	}

	networkConfigFile, err := fs.OpenFile("network-config", os.O_CREATE|os.O_RDWR)
	if err != nil {
		return err
	}
	defer networkConfigFile.Close()

	if _, err := networkConfigFile.Write([]byte(opts.NetworkConfig)); err != nil {
		return err
	}

	iso, ok := fs.(*iso9660.FileSystem)
	if !ok {
		return errors.New("invalid file system")
	}

	return iso.Finalize(iso9660.FinalizeOptions{
		RockRidge:        true,
		VolumeIdentifier: "cidata",
	})
}
