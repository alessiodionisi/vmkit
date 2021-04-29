#import "storage_device.h"

void *VZDiskImageStorageDeviceAttachment_init(const char *diskImageURL,
                                              bool readOnly, void **error) {
  VZDiskImageStorageDeviceAttachment *attachment;

  @autoreleasepool {
    NSString *string = [NSString stringWithUTF8String:diskImageURL];
    NSURL *url = [NSURL fileURLWithPath:string];

    attachment = [[VZDiskImageStorageDeviceAttachment alloc]
        initWithURL:url
           readOnly:readOnly
              error:(NSError *_Nullable *_Nullable)error];
  }

  return attachment;
}

void *VZVirtioBlockDeviceConfiguration_init(void *attachment) {
  return [[VZVirtioBlockDeviceConfiguration alloc]
      initWithAttachment:(VZStorageDeviceAttachment *)attachment];
}
