#ifndef network_device_h
#define network_device_h

#import <Virtualization/Virtualization.h>

void *VZVirtioNetworkDeviceConfiguration_init(void *attachment);
void VZVirtioNetworkDeviceConfiguration_setMACAddress(void *ptr,
                                                      void *macAddress);

void *VZNATNetworkDeviceAttachment_init();

#endif /* network_device_h */
