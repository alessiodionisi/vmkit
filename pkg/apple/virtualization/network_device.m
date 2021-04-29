#import "network_device.h"

void *VZVirtioNetworkDeviceConfiguration_init(void *attachment) {
  VZVirtioNetworkDeviceConfiguration *configuration =
      [[VZVirtioNetworkDeviceConfiguration alloc] init];
  [configuration setAttachment:(VZNetworkDeviceAttachment *)attachment];

  return configuration;
}

void *VZNATNetworkDeviceAttachment_init() {
  return [[VZNATNetworkDeviceAttachment alloc] init];
}
