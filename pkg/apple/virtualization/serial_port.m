#import "serial_port.h"

void *VZVirtioConsoleDeviceSerialPortConfiguration_init(void *attachment) {
  VZVirtioConsoleDeviceSerialPortConfiguration *configuration =
      [[VZVirtioConsoleDeviceSerialPortConfiguration alloc] init];
  [configuration setAttachment:(VZSerialPortAttachment *)attachment];

  return configuration;
}

void *VZFileHandleSerialPortAttachment_init(int read, int write) {
  NSFileHandle *readFileHandle =
      [[NSFileHandle alloc] initWithFileDescriptor:read];
  NSFileHandle *writeFileHandle =
      [[NSFileHandle alloc] initWithFileDescriptor:write];

  return [[VZFileHandleSerialPortAttachment alloc]
      initWithFileHandleForReading:readFileHandle
              fileHandleForWriting:writeFileHandle];
}
