# VMKit

Spin up Linux VMs with QEMU.

![Docker running on ARM64 Virtual Machine](/docs/docker.png)
_In the above image: Docker running on ARM64 Ubuntu Virtual Machine._

## Requirements

- Linux, macOS or Windows (needs testing) host OS.
- [QEMU](https://www.qemu.org) installed and available in the image, you can install it with homebrew or your package manager of choice. VMKit uses `qemu-img` binary, `qemu-system-aarch64` binary on ARM64 and `qemu-system-x86_64` binary on AMD64.

## Getting started

### Install VMKit

The only way for now is to have a working Go environment and install VMKit by running this command:

```shell
go install github.com/alessiodionisi/vmkit@latest
```

### Create your first vitual machine

Create a Debian 12 (Bookworm) virtual machine with 4 CPUs, 4096 mebibytes of ram and 20 GB of disk by running this command:

```shell
vmkit run debian12 -i debian:bookworm -c 4 -m 4096 -d 20
```

## Commands

### Create a virtual machines (`vmkit run`)

With `vmkit run` you can create and start a new virtual machine.

VMKit will automatically create a pair of SSH keys and configure the chosen system via `cloud-init`. If not specified, a forward to guest port 22 will be created using a free port. This will be used to access the virtual machine via SSH.

Example:

```shell
vmkit run debian12 -i debian:bookworm -c 4 -m 4096 -d 20
```

Available options:

- `-c`, `--cpu` number of cpu(s) (example: `-c 4`)
- `-d`, `--disk-size` disk size in gigabytes (GB) (example: `-d 20`)
- `-i`, `--image` image to use (example: `-i debian:bookworm`)
- `-m`, `--memory` ram in mebibytes (MiB) (example: `-m 4096`)
- `-p`, `--port-forward` forward host port to the virtual machine (example: `-p 8080-80`, `-p [host]-[guest]`)

### Remove a virtual machine (`vmkit remove|rm`)

With `vmkit remove` or `vmkit rm` you can remove a virtual machine.

Example:

```shell
vmkit remove debian12
```

### Start a virtual machine (`vmkit start`)

With `vmkit start` you can start a virtual machine.

Example:

```shell
vmkit start debian12
```

### Stop a virtual machine (`vmkit stop`)

With `vmkit stop` you can stop a virtual machine.

Example:

```shell
vmkit stop debian12
```

## Contributing

### Requirements

- [Go](https://golang.org) installed and available in the system.
- [Task](https://taskfile.dev) installed and available in the system.
