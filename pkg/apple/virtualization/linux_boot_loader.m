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

#import "linux_boot_loader.h"

void *VZLinuxBootLoader_init(const char *kernelURL) {
  NSString *string = [NSString stringWithUTF8String:kernelURL];
  NSURL *url = [NSURL fileURLWithPath:string];

  return [[VZLinuxBootLoader alloc] initWithKernelURL:url];
}

void VZLinuxBootLoader_setCommandLine(void *ptr, const char *commandLine) {
  NSString *string = [NSString stringWithUTF8String:commandLine];

  [(VZLinuxBootLoader *)ptr setCommandLine:string];
}

const char *VZLinuxBootLoader_commandLine(void *ptr) {
  return [[(VZLinuxBootLoader *)ptr commandLine] UTF8String];
}

void VZLinuxBootLoader_setInitialRamdiskURL(void *ptr,
                                            const char *initialRamdiskURL) {
  NSString *string = [NSString stringWithUTF8String:initialRamdiskURL];
  NSURL *url = [NSURL fileURLWithPath:string];

  [(VZLinuxBootLoader *)ptr setInitialRamdiskURL:url];
}
