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

#ifndef virtual_machine_configuration_h
#define virtual_machine_configuration_h

#import <Foundation/Foundation.h>
#import <Virtualization/Virtualization.h>

void *VZVirtualMachineConfiguration_init();
void VZVirtualMachineConfiguration_setBootLoader(void *ptr, void *bootLoader);
void VZVirtualMachineConfiguration_setCPUCount(void *ptr,
                                               unsigned long cpuCount);
unsigned long VZVirtualMachineConfiguration_CPUCount(void *ptr);
void VZVirtualMachineConfiguration_setMemorySize(void *ptr,
                                                 unsigned long long memorySize);
unsigned long long VZVirtualMachineConfiguration_memorySize(void *ptr);
bool VZVirtualMachineConfiguration_validateWithError(void *ptr, void **error);
void VZVirtualMachineConfiguration_setSerialPorts(void *ptr, void *serialPorts);
void VZVirtualMachineConfiguration_setStorageDevices(void *ptr,
                                                     void *storageDevices);
void VZVirtualMachineConfiguration_setNetworkDevices(void *ptr,
                                                     void *networkDevices);
void VZVirtualMachineConfiguration_setEntropyDevices(void *ptr,
                                                     void *entropyDevices);

#endif /* virtual_machine_configuration_h */
