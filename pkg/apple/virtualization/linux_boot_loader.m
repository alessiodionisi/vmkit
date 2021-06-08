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
