#import "boot_loader.h"

void *VZLinuxBootLoader_init(const char *kernelURL) {
	VZLinuxBootLoader *bootLoader;

	@autoreleasepool {
		NSString *string = [NSString stringWithUTF8String:kernelURL];
    NSURL *url = [NSURL fileURLWithPath:string];

		bootLoader = [[VZLinuxBootLoader alloc] initWithKernelURL:url];
	}

	return bootLoader;
}

void VZLinuxBootLoader_setCommandLine(void *ptr, const char *commandLine) {
	@autoreleasepool {
		NSString *string = [NSString stringWithUTF8String:commandLine];

		[(VZLinuxBootLoader *)ptr setCommandLine:string];
	}
}

const char *VZLinuxBootLoader_commandLine(void *ptr) {
	return [[(VZLinuxBootLoader *)ptr commandLine] UTF8String];
}

void VZLinuxBootLoader_setInitialRamdiskURL(void *ptr, const char *initialRamdiskURL) {
	@autoreleasepool {
		NSString *string = [NSString stringWithUTF8String:initialRamdiskURL];
		NSURL *url = [NSURL fileURLWithPath:string];

		[(VZLinuxBootLoader *)ptr setInitialRamdiskURL:url];
	}
}
