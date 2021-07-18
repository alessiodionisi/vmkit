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

#import "virtual_machine_configuration.h"

void *VZVirtualMachineConfiguration_init() {
  return [[VZVirtualMachineConfiguration alloc] init];
}

void VZVirtualMachineConfiguration_setBootLoader(void *ptr, void *bootLoader) {
  [(VZVirtualMachineConfiguration *)ptr
      setBootLoader:(VZLinuxBootLoader *)bootLoader];
}

void VZVirtualMachineConfiguration_setCPUCount(void *ptr,
                                               unsigned long cpuCount) {
  [(VZVirtualMachineConfiguration *)ptr setCPUCount:cpuCount];
}

unsigned long VZVirtualMachineConfiguration_CPUCount(void *ptr) {
  return [(VZVirtualMachineConfiguration *)ptr CPUCount];
}

void VZVirtualMachineConfiguration_setMemorySize(
    void *ptr, unsigned long long memorySize) {
  [(VZVirtualMachineConfiguration *)ptr setMemorySize:memorySize];
}

unsigned long long VZVirtualMachineConfiguration_memorySize(void *ptr) {
  return [(VZVirtualMachineConfiguration *)ptr memorySize];
}

bool VZVirtualMachineConfiguration_validateWithError(void *ptr, void **error) {
  return [(VZVirtualMachineConfiguration *)ptr
      validateWithError:(NSError *_Nullable *)error];
}

void VZVirtualMachineConfiguration_setSerialPorts(void *ptr,
                                                  void *serialPorts) {
  [(VZVirtualMachineConfiguration *)ptr
      setSerialPorts:[(NSMutableArray *)serialPorts copy]];
}

void VZVirtualMachineConfiguration_setStorageDevices(void *ptr,
                                                     void *storageDevices) {
  [(VZVirtualMachineConfiguration *)ptr
      setStorageDevices:[(NSMutableArray *)storageDevices copy]];
}

void VZVirtualMachineConfiguration_setNetworkDevices(void *ptr,
                                                     void *networkDevices) {
  [(VZVirtualMachineConfiguration *)ptr
      setNetworkDevices:[(NSMutableArray *)networkDevices copy]];
}

void VZVirtualMachineConfiguration_setEntropyDevices(void *ptr,
                                                     void *entropyDevices) {
  [(VZVirtualMachineConfiguration *)ptr
      setEntropyDevices:[(NSMutableArray *)entropyDevices copy]];
}
