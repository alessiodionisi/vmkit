#ifndef storage_device_h
#define storage_device_h

#import <Foundation/Foundation.h>
#import <Virtualization/Virtualization.h>

void *VZDiskImageStorageDeviceAttachment_init(const char *diskImageURL,
                                              bool readOnly, void **error);
void *VZVirtioBlockDeviceConfiguration_init(void *attachment);

#endif /* storage_device_h */
