#import "configuration.h"

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
