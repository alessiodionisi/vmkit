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

	dfs, err := diskfs.Create(opts.Name, 10*1024*1024, diskfs.Raw, diskfs.SectorSizeDefault)
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
