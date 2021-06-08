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
