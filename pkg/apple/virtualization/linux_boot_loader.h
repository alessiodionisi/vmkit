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
