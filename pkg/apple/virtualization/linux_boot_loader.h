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

#ifndef linux_boot_loader_h
#define linux_boot_loader_h

#import <Foundation/Foundation.h>
#import <Virtualization/Virtualization.h>

void *VZLinuxBootLoader_init(const char *kernelURL);
void VZLinuxBootLoader_setCommandLine(void *ptr, const char *commandLine);
const char *VZLinuxBootLoader_commandLine(void *ptr);
void VZLinuxBootLoader_setInitialRamdiskURL(void *ptr,
                                            const char *initialRamdiskURL);

#endif /* linux_boot_loader_h */
