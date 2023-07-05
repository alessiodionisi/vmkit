# VMKit

Spin up Linux VMs with QEMU.

![Docker running on ARM64 Virtual Machine](/docs/docker.png)
_In the above image: Docker running on ARM64 Ubuntu Virtual Machine_

## Getting started

_TODO: docs coming soon_

## Requirements

- QEMU installed and available in the system (you can install it with homebrew or your package manager of choice)
  - VMKit uses `qemu-img` binary, `qemu-system-aarch64` binary on ARM64 and `qemu-system-x86_64` binary on AMD64

## Commands

_TODO: docs coming soon_

## Tested on

- macOS on ARM64 _(Apple Silicon)_ and AMD64 with the latest stable version of QEMU

## Building

### Requirements

- Go 1.20.x

### How to

Simply use `go run ./cmd/vmkit` to build and run VMKit directly or `make build` to build a binary for your system.
