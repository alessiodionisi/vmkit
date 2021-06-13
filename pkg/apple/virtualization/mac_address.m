#import "mac_address.h"

void *VZMACAddress_init(const char *value) {
  NSString *string = [NSString stringWithUTF8String:value];

  return [[VZMACAddress alloc] initWithString:string];
}
