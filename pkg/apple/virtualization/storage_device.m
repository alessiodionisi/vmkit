#import "storage_device.h"

void *VZDiskImageStorageDeviceAttachment_init(const char *diskImageURL,
                                              bool readOnly, void **error) {
  NSString *string = [NSString stringWithUTF8String:diskImageURL];
  NSURL *url = [NSURL fileURLWithPath:string];

  return [[VZDiskImageStorageDeviceAttachment alloc]
      initWithURL:url
         readOnly:readOnly
            error:(NSError *_Nullable *)error];
}

void *VZVirtioBlockDeviceConfiguration_init(void *attachment) {
  return [[VZVirtioBlockDeviceConfiguration alloc]
      initWithAttachment:(VZStorageDeviceAttachment *)attachment];
}
