#import "serial_port.h"

void *VZVirtioConsoleDeviceSerialPortConfiguration_init(void *attachment) {
  VZVirtioConsoleDeviceSerialPortConfiguration *configuration =
      [[VZVirtioConsoleDeviceSerialPortConfiguration alloc] init];
  [configuration setAttachment:(VZSerialPortAttachment *)attachment];

  return configuration;
}

void *VZFileHandleSerialPortAttachment_init(int read, int write) {
  VZFileHandleSerialPortAttachment *attachment;

  @autoreleasepool {
    NSFileHandle *readFileHandle =
        [[NSFileHandle alloc] initWithFileDescriptor:read];
    NSFileHandle *writeFileHandle =
        [[NSFileHandle alloc] initWithFileDescriptor:write];

    attachment = [[VZFileHandleSerialPortAttachment alloc]
        initWithFileHandleForReading:readFileHandle
                fileHandleForWriting:writeFileHandle];
  }

  return attachment;
}
